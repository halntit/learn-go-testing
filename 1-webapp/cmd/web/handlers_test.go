package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/random", http.StatusNotFound},
	}

	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close() // as soon as test finish, close server

	// range through the tests
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestAppHome(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder() // response recorder

	handler := http.HandlerFunc(app.Home)
	handler.ServeHTTP(rr, req)

	// check the status code is what we expect
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want 200", rr.Code)
	}

	// check the response body is what we expect
	body, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(body), "From Session") {
		t.Error("home page should not have session info")
	}
}

func TestAppHome_new(t *testing.T) {
	var tests = []struct {
		name          string
		putInSessoion string
		expectedHtml  string
	}{
		{"first visit", "", "From Session"},
		{"second visit", "hello", "From Session hello"},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		req = addContextAndSessionToRequest(req, app)
		_ = app.Session.Destroy(req.Context()) // make sure nothing in session

		if e.putInSessoion != "" {
			app.Session.Put(req.Context(), "test", e.putInSessoion)
		}

		rr := httptest.NewRecorder() // response recorder
		handler := http.HandlerFunc(app.Home)
		handler.ServeHTTP(rr, req)

		// check the status code is what we expect
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want 200", rr.Code)
		}

		// check the response body is what we expect
		body, _ := io.ReadAll(rr.Body)
		if !strings.Contains(string(body), e.expectedHtml) {
			t.Errorf("%s did not find %s in body", e.name, e.expectedHtml)
		}
	}
}

func TestApp_renderWithBadTemplate(t *testing.T) {
	// set template path to a location with a bad template
	pathToTemplates = "./testdata/"

	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder()
	err := app.render(rr, req, "bad.page.gohtml", &TemplateData{})
	if err == nil {
		t.Error("expected error from bad template but not get one")
	}

	pathToTemplates = "./../../templates/" // reset to what it was
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))
	return req.WithContext(ctx)
}
