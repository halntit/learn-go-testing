package main

import "fmt"

func main() {
	n := 7

	_, msg := isPrime(n)
	fmt.Println(msg)
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
		if n % i == 0 {
			return false, fmt.Sprintf("%d is not prime number because it is devisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime", n)
}