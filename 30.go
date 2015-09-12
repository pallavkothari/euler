package main

import "fmt"

func main() {
	ans := 0
	for i := 2; i <= 6*pow(9, 5); i++ {
		if narcisstic(i) {
			ans += i
		}
	}
	fmt.Println(ans)
}

// check sum of 5th powers of digits = n
func narcisstic(n int) bool {
	m := n
	sum := 0
	for m >= 1 {
		sum += pow(m%10, 5)
		m /= 10
	}
	return n == sum
}

func pow(b, e int) int {
	ret := 1
	for i := 1; i <= e; i++ {
		ret *= b
	}
	return ret
}
