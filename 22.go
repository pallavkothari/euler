package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	names := string(dat)
	a := strings.Split(names, ",")
	sort.Strings(a)
	var sum int64 = 0
	for i, name := range a {
		sum += int64((i + 1) * value(strings.Trim(name, "\"")))
	}
	fmt.Println(sum)
}

func value(name string) int {
	sum := 0
	for _, r := range name {
		sum += int(r - 'A' + 1)
	}
	return sum
}
