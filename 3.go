package main

import "fmt"

func main() {
	n := 600851475143
	i := 2
	for i*i < n { // up to the sqrt of n
		if n%i == 0 && isPrime(i) {
			fmt.Println(i)
		}
		i++
	}
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
