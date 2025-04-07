package main

import (
	"log"
	"math/big"
	"os"
)

func rsa(p *big.Int, q *big.Int, message string) {
  n := new(big.Int)
  totient := new(big.Int)

  // n = p * q
  n.Mul(p, q)

  // totient := (p-1)*(q-1)
  totient.Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))

	// e : ensure gcd(e, totient) == 1
  // TODO: Find a way to generate random number and test gcd until i find
  // Do this efficiently
  e := big.NewInt(65537)
  if gcd(e, totient).Cmp(big.NewInt(1)) == 0 {
    log.Println("GCD found and E has been found -- ", e)
  }

	// d : d*e == 1 mod totient
  // consider the extended euclidean algorithm for this?
  d := big.NewInt(2)
  for d.Cmp(totient) == -1 {
    mul := new(big.Int) 
    mul.Mul(e, d)
    if mul.Mod(mul, totient).Cmp(big.NewInt(1)) == 0 {
      break
    }

    d.Add(d, big.NewInt(1))
  }

  log.Println("D FOUND -- ", d)
}

// Euclidean Algorithm
// NOTE: make sure that you pass a decent 'e' value here
func gcd(e *big.Int, totient *big.Int) *big.Int {
	gcd := new(big.Int)

	for {
    remainder := new(big.Int).Mod(totient, e)
		if remainder.Sign() == 0 {
			gcd = e
			break
		}
		totient = e
		e = remainder
	}

  log.Println("GCD FOUND: ", gcd)
	return gcd
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("usage: ./bin/problem_1 prime1 prime2 message")
	}

	p, _ := new(big.Int).SetString(os.Args[1], 10)
	q, _ := new(big.Int).SetString(os.Args[2], 10)

	if p == nil || q == nil {
    log.Fatal("error: issue decoding input numbers from string to BigInt")
	}

  rsa(p, q, os.Args[3])
}
