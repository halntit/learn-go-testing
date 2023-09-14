package main

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"webapp/pkg/data"
)

func Test_app_getTokenFromHeaderAndVerify(t *testing.T) {
	testUser := data.User {
		ID: 1,
		FirstName: "Admin",
		LastName: "User",
		Email: "admin@example.com",
		Password: "password",
		IsAdmin: 1,
	}

	tokens, _ := app.generateTokenPair(&testUser)

	var tests = []struct {
		name string
		token string
		errorExpected bool
		setHeader bool
		issuer string
	} {
		{"valid", fmt.Sprintf("Bearer %s", tokens.Token), false, true, app.Domain},
		{"expired", fmt.Sprintf("Bearer %s", expiredToken), true, true, app.Domain},
		{"no header", "", true, false, app.Domain},
		{"invalid", fmt.Sprintf("Bearer %s1", tokens.Token), true, true, app.Domain},
		{"no Bearer", fmt.Sprintf("Bear %s1", tokens.Token), true, true, app.Domain},
		{"3 header parts", fmt.Sprintf("Bear %s 1", tokens.Token), true, true, app.Domain},
		{"wrong issuer", fmt.Sprintf("Bearer %s", tokens.Token), true, true, "anotherdomain.com"},
	}

	for _, e := range tests {
		if e.issuer != "" {
			app.Domain = e.issuer
			tokens, _ = app.generateTokenPair(&testUser)
		}
		req, _ := http.NewRequest("GET", "/", nil)
		if e.setHeader {
			req.Header.Set("Authorization", e.token)
		}
		rr := httptest.NewRecorder()
		_, _, err := app.getTokenFromHeaderAndVerify(rr, req)

		if err != nil && !e.errorExpected {
			t.Errorf("%s: did not expect error but got %s", e.name, err.Error())
		}

		if err == nil && e.errorExpected {
			t.Errorf("%s: expected error but got none", e.name)
		}

		app.Domain = "example.com"
	}
}
