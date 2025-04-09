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
./exe p1 p2 m
```

Output will read as follows:
```
>>> Your E value
>>> Your D value
>>> message
>>> encoded message
>>> decoded message
```
You may not use any crypto packages or built in gcd() There are great heuristics / efficient methods for many of the steps required to complete this. DO NOT BRUTE FORCE THESE CALCS. Your code should run ‘generally’ quickly. If you have a concern about run time, meet with me. Prime generation is something you do separately and not included in the runtime of your program, however, for your own experience you should try writing a script to create and test the validity of large primes so that they fit the required scheme. If you don’t… it’ll be much harder.

### Implementation Details & How to Run
TODO
### How to Run: 
```
./bin/problem_1 $(cat input-rsa/p.txt) $(cat input-rsa/q.txt) message 
```

## Problem 2: ILP
The power of ILP (highly highly recommend using python here. If you choose not to you are on your own soldier!)

 a. Bron-Kerbosch: Given a graph in modified DIMACS format. Use the Bron- Kerbosch algorithm to return the maximum clique. (Pseudocode for this algorithm can be found on Wiki) It should run in the following manner:
```
>>> python script.py graph.txt
Max clique: [0, 1, 2, 3]
```

b. Cast the maximum clique problem as an ILP problem. Solve using Gurobi. It should run in the following manner:

```
>>> python ilp.py graph.txt
{there will be a LOT of Gurobi output here}
Max clique: [0, 1, 2, 3]
```

### TODO (RSA): 
- [x] Extended euclidean algorithm for getting private key
- [x] Actual encrypt + decrypt
- [ ] Clean up
- [x] Output formatting
- [ ] Documentation


### RSA Notes
- For the primes: 
        - [Miller Rabin Primality Test](https://en.wikipedia.org/wiki/Miller–Rabin_primality_test)

- For finding d:
        - [Extended Euclidean Algorithm](https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm)
