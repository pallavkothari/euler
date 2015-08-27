package main

import (
	"bufio"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	a := factorial(100)
	scanner := bufio.NewScanner(strings.NewReader(a.String()))
	scanner.Split(bufio.ScanRunes)
	sum := 0
	for scanner.Scan() {
		if d, e := strconv.Atoi(scanner.Text()); e == nil {
			sum += d
		}
	}
	fmt.Println(sum)
}

func factorial(n int64) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}
	f := big.NewInt(n)
	return f.Mul(f, factorial(n-1))
}
