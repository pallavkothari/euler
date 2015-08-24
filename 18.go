package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var triangle string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func main() {
	rows := make([][]int, 15)
	scanner := bufio.NewScanner(strings.NewReader(triangle))
	row := 0
	for scanner.Scan() {
		numScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		numScanner.Split(bufio.ScanWords)
		rows[row] = make([]int, 15)
		col := 0
		for numScanner.Scan() {
			if d, err := strconv.Atoi(numScanner.Text()); err == nil {
				rows[row][col] = d
			}
			col++
		}
		row++
	}
	path := findMaxPath(0, 0, rows)
	sum := 0
	for i := range path {
		sum += path[i]
	}
	fmt.Println(path, "(in reverse)")
	fmt.Println(sum)
}

func findMaxPath(i, j int, tree [][]int) []int {
	if i+1 == len(tree) {
		return []int{tree[i][j]}
	}
	max := max(findMaxPath(i+1, j, tree), findMaxPath(i+1, j+1, tree))
	max = append(max, tree[i][j])
	return max
}

func max(left, right []int) []int {
	leftSum, rightSum := 0, 0
	for i := 0; i < len(left); i++ {
		leftSum += left[i]
		rightSum += right[i]
	}
	if leftSum > rightSum {
		return left
	}
	return right
}
