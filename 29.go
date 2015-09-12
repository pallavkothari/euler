package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b float64
	m := make(map[string]bool)
	for a = 2; a <= 100; a++ {
		for b = 2; b <= 100; b++ {
			// wow, too weird. 7-14 places of precision does the trick
			// too few or too many and you're SOL. Also a different
			// format than 'G' and you're SOL
			s := strconv.FormatFloat(math.Pow(a, b), 'G', 14, 64)
			m[s] = true
		}
	}
	fmt.Println(len(m))
}
