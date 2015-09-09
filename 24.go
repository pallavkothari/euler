package main

import "fmt"

func main() {
	i := 1
	for s := range generate("0123456789") {
		if i == 1000000 {
			fmt.Println("the millionth permutation is", s)
			return
		}
		i++
	}
}

func generate(s string) <-chan string {
	c := make(chan string)
	// create a separate goroutine to feed into the channel
	go func(c chan string) {
		defer close(c) // this goroutine needs to close the channel
		permute(c, "", s)
	}(c)
	return c
}

func permute(c chan<- string, prefix string, s string) {
	if len(s) <= 1 {
		c <- (prefix + s)
		return
	}

	for _, ch := range s {
		theRest := filter(s, ch)
		permute(c, prefix+string(ch), theRest)
	}
}

func filter(s string, r rune) string {
	var ret string
	for _, ch := range s {
		if ch != r {
			ret += string(ch)
		}
	}
	return ret
}
