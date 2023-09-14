package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"webapp/pkg/data"
)

func Test_app_enableCORS(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	var tests = []struct {
		name string
		method string
		expectedHeader bool
	}{
		{"preflight", "OPTIONS", true},
		{"Get", "GET", false},
	}

	for _, e := range tests {
		handlerToTest := app.enableCORS(nextHandler)

		req, _ := http.NewRequest(e.method, "http://testing", nil)
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if e.expectedHeader && rr.Header().Get("Access-Control-Allow-Credentials") == "" {
			t.Errorf("%s: expected header", e.name)
		}

		if !e.expectedHeader && rr.Header().Get("Access-Control-Allow-Credentials") != "" {
			t.Errorf("%s: unexpected header", e.name)
		}
	}
}

func Test_app_authRequired(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
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
		expectedAuthorized bool
		setHeader bool
	} {
		{"valid", fmt.Sprintf("Bearer %s", tokens.Token), true, true},
		{"no token", "", false, false},
		{"invalid", fmt.Sprintf("Bearer %s", expiredToken), false, true},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		if e.setHeader {
			req.Header.Set("Authorization", e.token)
		}
		rr := httptest.NewRecorder()
		handlerToTest := app.authRequired(nextHandler)
		handlerToTest.ServeHTTP(rr, req)

		if e.expectedAuthorized && rr.Code == http.StatusUnauthorized {
			t.Errorf("%s: expected authorized but got unauthorized", e.name)
		}
			
		if !e.expectedAuthorized && rr.Code != http.StatusUnauthorized {
			t.Errorf("%s: expected unauthorized but got authorized", e.name)
		}
	}
}
