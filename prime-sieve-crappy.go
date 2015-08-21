package main

import "fmt"

// send the seq 2,3,4.. to channel ch
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // send i to ch
	}
}

// copy from in to out, filtering out those divisible by prime
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// daisy-chain filter processes
func main() {
	ch := make(chan int)
	go Generate(ch)
	lastPrime := 2
	for i := 0; i < 10001; i++ {
		prime := <-ch
		lastPrime = prime
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println(lastPrime)
}
