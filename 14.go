package main

import "fmt"

// cache chain lengths
var collatz map[int]int

func main() {
	collatz = make(map[int]int)
	n, max := 0, 0
	for i := 1; i <= 1000000; i++ {
		if c := chainLen(i); c > max {
			n = i
			max = c
		}
	}
	fmt.Println(n, max)
}

func chainLen(n int) int {
	if n == 1 {
		return 1
	}

	if c, ok := collatz[n]; ok {
		return c
	}

	if n%2 == 0 {
		c := 1 + chainLen(n/2)
		collatz[n] = c
		return c
	} else {
		c := 1 + chainLen(3*n+1)
		collatz[n] = c
		return c
	}
}
