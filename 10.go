package main

import "fmt"

func main() {
	sum := 0

	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Printf("sum = %d", sum)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
