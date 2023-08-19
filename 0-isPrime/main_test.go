package main

import (
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"not prime", -1, false, "-1 is negative and is not prime"},
		{"not prime", 0, false, "0 is not prime"},
		{"not prime", 1, false, "1 is not prime"},
		{"prime", 2, true, "2 is prime"},
		{"prime", 3, true, "3 is prime"},
		{"not prime", 4, false, "4 is not prime number because it is devisible by 2"},
		{"prime", 5, true, "5 is prime"},
		{"not prime", 6, false, "6 is not prime number because it is devisible by 2"},
		{"prime", 7, true, "7 is prime"},
		{"not prime", 8, false, "8 is not prime number because it is devisible by 2"},
		{"not prime", 9, false, "9 is not prime number because it is devisible by 3"},
		{"not prime", 10, false, "10 is not prime number because it is devisible by 2"},
		{"prime", 11, true, "11 is prime"},
		{"not prime", 12, false, "12 is not prime number because it is devisible by 2"},
		{"prime", 13, true, "13 is prime"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%d %s expected true but got false", e.testNum, e.name)
		}
		if !e.expected && result {
			t.Errorf("%d %s expected false but got true", e.testNum, e.name)
		}

		if msg != e.msg {
			t.Errorf("%d %s expected %s, but got %s", e.testNum, e.name, msg, e.msg)
		}
	}
}
