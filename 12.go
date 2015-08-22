package main

import "fmt"

var isPrime map[int]bool
var primeFactors map[int]map[int]int

func main() {
	isPrime = make(map[int]bool)
	primeFactors = make(map[int]map[int]int)
	for i := 1; ; i++ {
		tr := (i + 1) * i / 2 // triangle numbers are just the sum of the first N
		trPrimeFactors := getFactors(tr)
		numDivisors := 1
		for _, v := range trPrimeFactors {
			numDivisors *= (v + 1)
		}
		if numDivisors > 500 {
			fmt.Println(tr, ":", trPrimeFactors, ":", numDivisors)
			return
		}
	}
}

func getFactors(n int) map[int]int {
	myMap := make(map[int]int)

	num := n
outer:
	for num > 1 {
		for p := 2; p <= num; p++ {
			if isPrimeNumber(p) {
				for num%p == 0 {
					num /= p
					myMap[p]++
					if v, found := primeFactors[num]; num > 1 && found {
						// found the factors of this smaller number!
						for k, v := range v {
							myMap[k] = myMap[k] + v
						}
						break outer
					}
				}
			}
		}
	}
	primeFactors[n] = myMap
	return myMap
}

func isPrimeNumber(n int) bool {
	if _, found := isPrime[n]; !found {
		isPrime[n] = checkPrime(n)
	}
	return isPrime[n]
}

func checkPrime(n int) bool {
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			return false
		}
	}
	return true
}
