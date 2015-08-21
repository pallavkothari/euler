package main

import "fmt"

func main() {
	max := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			if isPalindrome(i*j) && i*j > max {
				max = i * j
			}
		}
	}
	fmt.Println(max)
}

func isPalindrome(num int) bool {
	n := num
	rev := 0
	for n > 0 {
		dig := n % 10
		rev = rev*10 + dig
		n = n / 10
	}
	return num == rev
}
