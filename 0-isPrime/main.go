package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print welcome message
	intro()

	// create a channel to indicate when user want to quit
	doneChan := make(chan bool)

	// start a go routine to check if the number is prime
	go readUserInput(os.Stdin, doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNum(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNum(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	// check if want to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("-------------")
	fmt.Println("Enter a number:")

	// get user input
	prompt()
}

func prompt() {
	fmt.Print("> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime", n)
	}
	if n < 0 {
		return false, fmt.Sprintf("%d is negative and is not prime", n)
	}

	// use the modulus operator repeatedly to see if it is prime
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime number because it is devisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime", n)
}
