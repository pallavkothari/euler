package main

import (
	"fmt"
	"math"
)

// Find the index of the first fibonacci number to contain 1000 digits
//
// F(n) = phi^n/sqrt(5) for large n
// so we are looking for n s.t.
// phi^n/sqrt(5) > 10^999
//
// taking logs,
// n*log(phi) - log(5)/2 > 999*log(10)
//
// or, n > (999*log(10) + log(5)/2) / log(phi)
// where phi ~= 1.6180

func main() {
	fmt.Println(math.Floor((999*math.Log(10) + (math.Log(5) / 2)) / math.Log(1.6180)))
}
