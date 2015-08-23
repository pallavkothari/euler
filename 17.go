package main

import "fmt"

var words = [100]string{"", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}

func init() {
	words[30] = "thirty"
	words[40] = "forty"
	words[50] = "fifty"
	words[60] = "sixty"
	words[70] = "seventy"
	words[80] = "eighty"
	words[90] = "ninety"
}

var count int

func main() {
	for i := 1; i < 1001; i++ {
		buf := split(i)
		if len(buf) == 4 {
			// fmt.Println("one thousand")
			count += 11
		}
		if len(buf) == 3 {
			// fmt.Print(words[buf[2]], " hundred")
			count += len(words[buf[2]]) + 7
			if buf[0] != 0 || buf[1] != 0 {
				// fmt.Print(" and ")
				count += 3
			}
		}
		printUpto100(i%100, buf[0:len(buf)])
	}
	fmt.Println(count)
}

func printUpto100(i int, buf []int) {
	if i <= 20 {
		// fmt.Println(words[i])
		count += len(words[i])
		return
	}
	// fmt.Print(words[buf[1]*10], " ")
	count += len(words[buf[1]*10])
	// fmt.Print(words[buf[0]])
	count += len(words[buf[0]])
	// fmt.Println()
}

func split(n int) []int {
	var buf []int
	num := n
	for num > 0 {
		buf = append(buf, num%10)
		num /= 10
	}
	return buf
}
