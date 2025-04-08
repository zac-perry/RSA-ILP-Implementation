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
	e := big.NewInt(65537)
	if gcd(e, totient).Cmp(big.NewInt(1)) == 0 {
		log.Println("GCD found and E has been found -- ", e)
	}

	// d : d*e == 1 mod totient
	// consider the extended euclidean algorithm for this
	d := findD(e, totient)
	log.Println("D FOUND -- ", d)

	// encryption / decryption
	cipherText := encrypt(message, e, n)
	log.Println("CIPHER TEXT -- ", cipherText)

	plainText := decrypt(cipherText, d, n)
	log.Println("plaintext -- ", plainText)
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

// Extended Euclidean Algorithm
func findD(e, totient *big.Int) *big.Int {
	// extended gcd.
	// TODO: refactor so it's less disgusting looking
	var egcd func(a, b *big.Int) (*big.Int, *big.Int, *big.Int)
	egcd = func(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
		if b.Cmp(big.NewInt(0)) == 0 {
			return a, big.NewInt(1), big.NewInt(0)
		}
		temp1 := new(big.Int).Mod(a, b)
		g, y, x := egcd(b, temp1)

		// y = y - (a/b)* x
		temp := new(big.Int).Div(a, b)
		temp = new(big.Int).Mul(temp, x)
		temp = new(big.Int).Sub(y, temp)

		return g, x, temp
	}

	// mod inverse to get D.
	gcd, x, _ := egcd(e, totient)

	if gcd.Cmp(big.NewInt(1)) != 0 || gcd == nil {
		log.Fatal("error (findD): modInverse returned nil")
	}

	return new(big.Int).Mod(x, totient)
}

func encrypt(message string, e, n *big.Int) *big.Int {
	// NOTE: need to do exp with all three args here, as doingt the opeartions seperate will cause
	// Massive loop
	cipherBytes := new(big.Int).SetBytes([]byte(message))
	cipherText := new(big.Int).Exp(cipherBytes, e, n)

	return cipherText
}

func decrypt(cipherText, d, n *big.Int) string {
	// NOTE: need to do exp with all three args here, as doingt the opeartions seperate will cause
	// Massive loop
	plainText := new(big.Int).Exp(cipherText, d, n)

	return string(plainText.Bytes())
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
