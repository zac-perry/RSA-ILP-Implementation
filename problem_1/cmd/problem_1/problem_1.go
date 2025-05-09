/*
 * Zachary Perry
 * COSC 581 Assignment 7: Part 1
 * 4/10/25
 */
package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

/*
RSA encryption algorithm that uses 2048-bit primes p and q such that: p = 2q + 1
Follows the general RSA encryption algorithm defined in class.
*/
func rsa(p *big.Int, q *big.Int, message string) (*big.Int, *big.Int, *big.Int, string) {
	// n = p * q
	// totient := (p-1)*(q-1)
	n := new(big.Int).Mul(p, q)
	totient := new(big.Int).Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))

	// e : ensure gcd(e, totient) == 1
	e := big.NewInt(65537)
	if gcd(e, totient).Cmp(big.NewInt(1)) != 0 {
		log.Fatal("error (rsa): issue finding e")
	}

	// d : d*e == 1 mod totient
	d := findD(e, totient)

	// encryption / decryption
	cipherText := encrypt(message, e, n)
	plainText := decrypt(cipherText, d, n)

	return e, d, cipherText, plainText
}

// Euclidean Algorithm to find GCD
// source: https://en.wikipedia.org/wiki/Euclidean_algorithm
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

	return gcd
}

// Extended GCD to assist in finding the value for D.
// NOTE: this is used for the extended euclidean algorithm.
func extendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a, big.NewInt(1), big.NewInt(0)
	}
	temp1 := new(big.Int).Mod(a, b)
	g, y, x := extendedGCD(b, temp1)

	// y = y - (a/b)* x
	temp := new(big.Int).Div(a, b)
	temp = new(big.Int).Mul(temp, x)
	temp = new(big.Int).Sub(y, temp)

	return g, x, temp
}

// Extended Euclidean Algorithm to find D
// source: https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
func findD(e, totient *big.Int) *big.Int {
	// mod inverse to get D.
	gcd, x, _ := extendedGCD(e, totient)

	if gcd.Cmp(big.NewInt(1)) != 0 || gcd == nil {
		log.Fatal("error (findD): modInverse returned nil")
	}

	return new(big.Int).Mod(x, totient)
}

// Encrypt the given message
func encrypt(message string, e, n *big.Int) *big.Int {
	// NOTE: need to do exp with all three args here, as doing the opeartions seperate will cause a massive loop
	cipherBytes := new(big.Int).SetBytes([]byte(message))
	cipherText := new(big.Int).Exp(cipherBytes, e, n)

	return cipherText
}

// Decrypt the given message
func decrypt(cipherText, d, n *big.Int) string {
	// NOTE: need to do exp with all three args here, as doing the opeartions seperate will cause a massive loop
	plainText := new(big.Int).Exp(cipherText, d, n)

	return string(plainText.Bytes())
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("usage: ./bin/problem_1 prime1 prime2 message")
	}

	p, _ := new(big.Int).SetString(os.Args[1], 10)
	q, _ := new(big.Int).SetString(os.Args[2], 10)
	message := os.Args[3]

	if p == nil || q == nil {
		log.Fatal("error: issue decoding input numbers from string to BigInt")
	}

	e, d, cipherText, plainText := rsa(p, q, message)

	fmt.Printf(">>> E: %5v\n", e)
	fmt.Printf(">>> D: %5v\n", d)
	fmt.Printf(">>> Message: %5s\n", message)
	fmt.Printf(">>> CipherText: %5v\n", cipherText)
	fmt.Printf(">>> PlainText: %5s\n", plainText)
}
