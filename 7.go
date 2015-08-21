package main

import "fmt"

func main() {
	numPrimes := 0
	lastPrime := 2

	for i := 2; numPrimes < 10001; i++ {
		if isPrime(i) {
			numPrimes++
			lastPrime = i
		}
	}
	fmt.Printf("numPrimes = %d lastPrime = %d", numPrimes, lastPrime)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
