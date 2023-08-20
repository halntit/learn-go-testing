package main

import (
	"testing"
	"os"
	"io"
	"strings"
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



func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}



func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Is it prime?") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}
