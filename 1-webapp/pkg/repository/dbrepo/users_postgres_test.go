package dbrepo

import (
	"database/sql"
	"fmt"
	"time"
	"log"
	"os"
	"testing"
	"webapp/pkg/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB
var testRepo *PostgresDBRepo

func TestMain(m *testing.M) {
	// connect to docker; fail if docker not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker; is it running? %s", err)
	}

	pool = p

	// set up our docker options, specifying the image and so forth
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
		_ = pool.Purge(resource)
	}

	// start the image and wait until it's ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to database: %s", err)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("Error creating tables: %s", err)
	}

	// set up test repo, must be before running tests
	testRepo = &PostgresDBRepo{DB: testDB}

	// run tests
	code := m.Run()

	// clean up (remove container and image)
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("count not purge resource: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestPostgresDBRepoInsertUser(t *testing.T) {
	testUser := data.User{
		Email: "x@x.com",
		Password: "password",
		FirstName: "Admin",
		LastName: "User",
		IsAdmin: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("Insert user return error: %s", err)
	}

	if id != 1 {
		t.Errorf("Insert user return wrong id: %d", id)
	}
}

func TestPostgresDBRepoAllUsers(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Errorf("All users return error: %s", err)
	}

	if len(users) != 1 {
		t.Errorf("All users return wrong number of users: %d", len(users))
	}

	testUser := data.User{
		Email: "x2@x.com",
		Password: "password",
		FirstName: "Admin",
		LastName: "Super",
		IsAdmin: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, _ = testRepo.InsertUser(testUser)
	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("All users return error: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("All users return wrong number of users: %d", len(users))
	}
}

func TestPostgresDBRepoGetUser(t *testing.T) {
	user, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("Get user by id return error: %s", err)
	}

	if user.Email != "x@x.com" {
		t.Errorf("Get user by id return wrong email: %s", user.Email)
	}

	user, err = testRepo.GetUserByEmail("x@x.com")
	if err != nil {
		t.Errorf("Get user by id return error: %s", err)
	}

	if user.ID != 1 {
		t.Errorf("Get user by email return wrong id: %d", user.ID)
	}
}

func TestPostgresDBRepoUpdateUser(t *testing.T) {
	testUser := data.User{
		Email: "x@x.com",
		Password: "password",
		FirstName: "Admin",
		LastName: "User",
		IsAdmin: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("Insert user return error: %s", err)
	}

	testUser.Email = "x2@x.com"
	testUser.ID = id
	err = testRepo.UpdateUser(testUser)
	if err != nil {
		t.Errorf("Update user return error: %s", err)
	}

	user, err := testRepo.GetUser(id)
	if user.Email != "x2@x.com" {
		t.Errorf("Get user by id return wrong email: %s", user.Email)
	}
}

func TestPostgresDBRepoDeleteUser(t *testing.T) {
	err := testRepo.DeleteUser(1)
	if err != nil {
		t.Errorf("Delete user return error: %s", err)
	}

	user, err := testRepo.GetUser(1)
	if user != nil {
		t.Errorf("Get user by id return wrong user: %v", user)
	}
}
