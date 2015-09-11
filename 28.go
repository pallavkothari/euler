package main

import "fmt"

func main() {
	e, sum := 1, 1
	for spiral := 1; spiral <= 500; spiral++ {
		for corner := 1; corner <= 4; corner++ {
			e += spiral * 2
			sum += e
		}
	}
	fmt.Println(sum)
}
