package main

import (
	"fmt"
	"math/big"
)

// this is just a pascal's triangle problem in disguise
// since we are working with a square, not a triangle,
// we want the middle coefficient at the 40th row of
// pascal's triangle (for a 20x20 grid's answer)
// and that binomial coefficient is given by 40c20
//
func main() {
	f40 := factorial(40)
	f20 := factorial(20)
	d := big.NewInt(1)
	d.Mul(f20, f20)

	fmt.Println(f40.Div(f40, d))
}

func factorial(n int64) *big.Int {
	if n <= 0 {
		return big.NewInt(1)
	}
	f := big.NewInt(n)
	return f.Mul(f, factorial(n-1))
}
