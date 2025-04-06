package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// TODO: Function to find primes
// 2048 bit primes
// restriction: p = 2q + 1
func primes() {}

// TODO: RSA
func rsa(p int, q int, message string) {
	n := p * q
	totient := (p - 1) * (q - 1)

	// e : gcd(e, totient)

	// d : d*e == 1 mod totient

	fmt.Println(n, totient)
}

// TODO: Manual GCD
// Euclidean Algorithm
// NOTE: make sure that you pass a decent 'e' value here
func gcd(e int, totient int) int {
	remainder := 1
	gcdVal := 0

	for {
		remainder = totient % e
		if remainder == 0 {
			gcdVal = e
			break
		}
		totient = e
		e = remainder
	}

	return gcdVal
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("usage: ./bin/problem_1 prime1 prime2 message")
	}

	// may need to fix ATOI depending on the values passed?
	p, _ := strconv.Atoi(os.Args[1])
	q, _ := strconv.Atoi(os.Args[2])
	fmt.Printf("Finding the GCD for: %d %d\n", p, q)
	fmt.Printf("GCD: %d\n", gcd(p, q))
}
