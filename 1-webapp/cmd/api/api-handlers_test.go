package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email":"admin@example.com", "password":"secret"}`, http.StatusOK},
		{"not json", `I am not JSON`, http.StatusUnauthorized},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email":""}`, http.StatusUnauthorized},
		{"empty password", `{"email":"admin@example.com","password":""}`, http.StatusUnauthorized},
		{"invalid user", `{"email":"random@example.com","password":"wordpass"}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {
		var reader io.Reader
		reader = strings.NewReader(e.requestBody)
		req, _ := http.NewRequest("POST", "/authenticate", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)

		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s returned wrong status code: got %d, want %d", e.name, rr.Code, e.expectedStatusCode)
		}
	}
}
