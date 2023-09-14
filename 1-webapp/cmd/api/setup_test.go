package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2OTQzNDk1MjIsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.MLgzoUwN3noJvMp2vmPxoZN8cNVcmvg2tkvV8ToiR6g"

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "s3c4et"
	os.Exit(m.Run())
}
