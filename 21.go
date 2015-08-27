package main

import "fmt"

var amicable = make(map[int]bool)

func main() {
	sum := 0
	for a := 1; a < 10000; a++ {
		if amicable, present := amicable[a]; amicable && present {
			continue
		}
		if b := d(a); a != b && a == d(b) {
			sum += a + b
			amicable[b] = true
		}
	}
	fmt.Println(sum)
}

func d(n int) int {
	divisors := divisors(n)
	sum := 0
	for _, v := range divisors {
		sum += v
	}
	return sum
}

func divisors(n int) []int {
	d := make([]int, 0)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			d = append(d, i)
		}
	}
	return d
}
