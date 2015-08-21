package main

import "fmt"

var memo map[int64]int64

func init() {
	memo = make(map[int64]int64)
	memo[1] = 1
	memo[2] = 1
}

func fib(i int64) int64 {
	if f, found := memo[i]; found {
		return f
	}
	memo[i] = fib(i-1) + fib(i-2)
	return memo[i]
}

func main() {
	var i int64 = 2
	var sum int64 = 0
	for ; fib(i) <= 4000000; i = i + 1 {
		if fib(i)%2 == 0 {
			sum += fib(i)
			// fmt.Printf("fib(%v) = %v\n", i, fib(i))
		}
	}
	fmt.Println(sum)
}
