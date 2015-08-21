package main

import "fmt"

func main() {
	n := 100
	sum := (n + 1) * n / 2
	sqOfSum := sum * sum
	sumOfSquares := n * (n + 1) * (2*n + 1) / 6
	fmt.Println(sqOfSum - sumOfSquares)
}
