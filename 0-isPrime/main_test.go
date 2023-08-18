package main

import (
	"testing"
)

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not prime" {
		t.Error("wrong message returned: ", msg)
	}


	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", 0)
	}

	if msg != "7 is prime" {
		t.Error("wrong message returned: ", msg)
	}
}