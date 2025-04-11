# COSC 581 -- HW 7 
Zachary Perry

## Folder Structute
```
    problem_1/: contains all code related to RSA
    problem_2/: contains all code realted to ILP
```

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

### Requirements
Go version 1.22.9 (this is what's on the hydra + tesla)

### How to Run: 
```
NOTE: Compile using: make
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

You will need to write a short README on the run times for part a and b. Test your
codes on graphs of 1000 vertices and give me the runtimes for part a and b. You will
need to be connected to UTK pulse secure vpn to get a Gurobi license and install if you
are off campus. Gurobipy is very easy to use and well documented. I recommend using
networkx to generate large pseudo-random graphs to test on.


### How to run
```
python3 src/problem_2a.py input/graph.txt
python3 src/problem_2b.py input/graph.txt
```

### Results: 
From problem 2, I found that Gurobi was slightly slower than the Bron-Kerbosch algorithm for a graph with 1000 vertices. My runtimes for both were as follows: 
```
Bron-Kerbosch: 0.00111 seconds

Gurobi: 0.88 seconds
```

