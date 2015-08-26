package main

import (
	"fmt"
	"time"
)

func main() {
	sundays := 0
	for y := 1901; y <= 2000; y++ {
		for m := time.January; m <= time.December; m++ {
			t := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == time.Sunday {
				sundays++
			}
		}
	}
	fmt.Println(sundays)
}

