package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
)

func generatePrimes(bits int) (*big.Int, *big.Int, error) {
	random, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), uint(bits)))
	if err != nil {
		return nil, nil, err
	}

	random.SetBit(random, 0, 1)
	p := new(big.Int)
	q := new(big.Int)

	q.Set(random)

	// p = 2q + 1
	p.Mul(q, big.NewInt(2))
	p.Add(p, big.NewInt(1))

	if !q.ProbablyPrime(20) || !p.ProbablyPrime(20) {
		q, p, err = generatePrimes(bits)
		if err != nil {
			return nil, nil, err
		}
	}

	return p, q, nil
}

func writePrimes(p, q *big.Int) error {
	fileP, err := os.Create("rsa/p.txt")
	fileQ, err := os.Create("rsa/q.txt")
	if err != nil {
		return err
	}
	defer fileP.Close()
	defer fileQ.Close()

	fileP.WriteString(p.String())
	fileQ.WriteString(q.String())

	return nil
}

func main() {
	fmt.Println("Generating large primes...")

	p, q, err := generatePrimes(2048)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Primes Found!")
	log.Println("Generated p: ", p)
	log.Println("Generated q: ", q)

	err = writePrimes(p, q)
	if err != nil {
		log.Fatal(err)
	}
}
