package main

import "fmt"

func main() {
	maxN, ab := 0, 0
	for a := -1000; a < 1000; a++ {
		for b := -1000; b < 1000; b++ {
			for n := 0; isPrime(abs(n*n + a*n + b)); n++ {
				if n > maxN {
					ab = a * b
					maxN = n
				}
			}
		}
	}
	fmt.Println(ab, maxN)
}
func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
