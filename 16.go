package main

import (
	"bufio"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	n := new(big.Int)
	n.Exp(big.NewInt(int64(2)), big.NewInt(int64(1000)), nil)
	scanner := bufio.NewScanner(strings.NewReader(n.String()))
	scanner.Split(bufio.ScanRunes)
	sum := 0
	for scanner.Scan() {
		if d, err := strconv.Atoi(scanner.Text()); err == nil {
			sum += d
		}
	}
	fmt.Println(sum)
}
