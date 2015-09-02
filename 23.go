package main

import "fmt"

var abundants = make(map[int]bool)
var abundantMultiples = make(map[int]bool)
var limit = 28123

func main() {
	for i := 12; i <= limit; i++ {
		if isAbundant(i) {
			markMultiples(i)
		}
	}

	sum := 0
	for i := 1; i <= limit; i++ {
		if !abundantMultiples[i] {
			sum += i
		}
	}
	fmt.Println(sum)
}

func markMultiples(n int) {
	abundants[n] = true
	for j := 12; j <= n; j++ {
		if abundants[n] && abundants[j] {
			abundantMultiples[n+j] = true
		}
	}
	for i := 2 * n; i <= limit; i += n {
		abundantMultiples[i] = true
	}
}

func isAbundant(i int) bool {
	return sumOfProperDivisors(i) > i
}

func sumOfProperDivisors(n int) int {
	sum := 0
	for i := 1; i <= (n / 2); i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
