package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var mem = make(map[string]int)

func main() {
	solve()
}

func solve() {
	dat, err := ioutil.ReadFile("p067_triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	triangle := string(dat)
	rows := make([][]int, 100)
	scanner := bufio.NewScanner(strings.NewReader(triangle))
	row := 0
	for scanner.Scan() {
		numScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		numScanner.Split(bufio.ScanWords)
		rows[row] = make([]int, 100)
		col := 0
		for numScanner.Scan() {
			if d, err := strconv.Atoi(numScanner.Text()); err == nil {
				rows[row][col] = d
			}
			col++
		}
		row++
	}
	sum := findMaxPath(0, 0, rows)
	fmt.Println(sum)
}

func findMaxPath(i, j int, tree [][]int) int {
	if i+1 == len(tree) {
		return tree[i][j]
	}
	if s, ok := mem[string(i)+":"+string(j)]; ok {
		return s
	}
	max := max(findMaxPath(i+1, j, tree), findMaxPath(i+1, j+1, tree))
	max += tree[i][j]
	mem[string(i)+":"+string(j)] = max
	return max
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
