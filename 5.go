package main

import "fmt"

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	a, b := 1, 2
	for b <= 20 {
		a = lcm(a, b)
		b++
	}
	fmt.Println(a)
}
