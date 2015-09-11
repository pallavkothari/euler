package main

import "fmt"

func main() {
	max := 0
	maxI := 0
	for i := 1000; i > 0; i-- {
		remainders := make(map[int]bool)
		n := 1
		cycle := 0
		for !remainders[n%i] {
			cycle++
			remainders[n%i] = true
			n = 10 * (n % i)
		}
		if cycle > max {
			max = cycle
			maxI = i
		}
	}
	fmt.Println(maxI, max)
}
