package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webapp/pkg/data"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "127.0.0.1", "", false},
		{"", "", "hello:word", false},
	}

	// create a dummy handler that we'll use to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		// make sure we got a string back
		_, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
	})

	for _, e := range tests {
		// create a request
		handlerToTest := app.addIPToContext(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Set(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	// get a context
	ctx := context.Background()

	// put something in the context
	ctx = context.WithValue(ctx, contextUserKey, "whatever")

	// call the function
	ip := app.ipFromContext(ctx)

	// perform the test
	if !strings.EqualFold("whatever", ip) {
		t.Error("wrong value returned")
	}
}

func Test_app_auth(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
	})

	var tests = []struct {
		name string
		isAuth bool
	}{
		{"logged in", false},
		{"not logged in", true},
	}

	for _, e := range tests {
		// create a request
		handlerToTest := app.auth(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if e.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if e.isAuth && rr.Code != http.StatusOK {
			t.Errorf("%s returned wrong status code: got %v want 200", e.name, rr.Code)
		}

		if !e.isAuth && rr.Code != http.StatusTemporaryRedirect {
			t.Errorf("%s returned wrong status code: got %v want 307", e.name, rr.Code)
		}
	}
}
