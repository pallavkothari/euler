package main

import "fmt"

func main() {
	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 1000; b++ {
			if sq, c := isSq(a*a + b*b); sq && a+b+c == 1000 {
				fmt.Printf("%d %d %d = %d\n", a, b, c, a*b*c)
			}
		}
	}
}

func isSq(n int) (bool, int) {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true, i
		}
	}
	return false, 0
}
