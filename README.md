# COSC 581 -- HW 7 
Solutions for hw 7

## Folder Structute

## Compile & Run

## Problem 1: RSA
You must create an RSA encryption algorithm that uses 2048-bit primes p and q such that: p = 2q + 1

It should run in the following manner:
```
P1 = prime 1
P2 = prime 2
M = message (string)
&gt;&gt; ./exe p1 p2 m
```

Output will read as follows:
```
&gt;&gt; Your E value
&gt;&gt; Your D value
&gt;&gt; message
&gt;&gt; encoded message
&gt;&gt; decoded message
```
You may not use any crypto packages or built in gcd() There are great heuristics / efficient methods for many of the steps required to complete this. DO NOT BRUTE FORCE THESE CALCS. Your code should run ‘generally’ quickly. If you have a concern about run time, meet with me. Prime generation is something you do separately and not included in the runtime of your program, however, for your own experience you should try writing a script to create and test the validity of large primes so that they fit the required scheme. If you don’t… it’ll be much harder.

TODO: 
- [ ] look into the algorithm itself

## Problem 2: ILP


### TODO: 


### RSA Notes
- For the primes: 
    - Can use one of the popular algorithms or primality testing for this: 
        - [Sieve of Atkin](https://en.wikipedia.org/wiki/Sieve_of_Atkin)
        - [Miller Rabin Primality Test](https://en.wikipedia.org/wiki/Miller–Rabin_primality_test)

- rest of the algo is pretty easy.
    - Just make helpers for e, d, etc. 
    - Also manually do GCD, don't use any built in things.
