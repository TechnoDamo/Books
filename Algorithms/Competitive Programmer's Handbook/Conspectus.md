# Table of Contents

## Basic Techniques
1. [Introduction](#1-introduction)  
2. [Time complexity](#2-time-complexity)  
3. [Sorting](#3-sorting)  
4. [Data structures](#4-data-structures)  
5. [Complete search](#5-complete-search)  
6. [Greedy algorithms](#6-greedy-algorithms)  
7. [Dynamic programming](#7-dynamic-programming)  
8. [Amortized analysis](#8-amortized-analysis)  
9. [Range queries](#9-range-queries)  
10. [Bit manipulation](#10-bit-manipulation)

## Graph Algorithms
11. [Basics of graphs](#11-basics-of-graphs)  
12. [Graph traversal](#12-graph-traversal)  
13. [Shortest paths](#13-shortest-paths)  
14. [Tree algorithms](#14-tree-algorithms)  
15. [Spanning trees](#15-spanning-trees)  
16. [Directed graphs](#16-directed-graphs)  
17. [Strong connectivity](#17-strong-connectivity)  
18. [Tree queries](#18-tree-queries)  
19. [Paths and circuits](#19-paths-and-circuits)  
20. [Flows and cuts](#20-flows-and-cuts)

## Advanced Topics
21. [Number theory](#21-number-theory)  
22. [Combinatorics](#22-combinatorics)  
23. [Matrices](#23-matrices)  
24. [Probability](#24-probability)  
25. [Game theory](#25-game-theory)  
26. [String algorithms](#26-string-algorithms)  
27. [Square root algorithms](#27-square-root-algorithms)  
28. [Segment trees revisited](#28-segment-trees-revisited)  
29. [Geometry](#29-geometry)  
30. [Sweep line algorithms](#30-sweep-line-algorithms)

# Basic Technniques
## 1. Inroduction
### Input and output
```cpp
#include <bits/stdc++.h> // includes all standard library
using namespace std;
int main() {
    // solution
}
```
We can compile code with `g++ -std=c++11 -02 -Wall test.cpp -o test`. This command produces a binary file test from the source code test.cpp. 
The following lines at the beginning of the code make input and output more efficient:
```cpp
ios::sync_with_stdio(0)
cin.tie(0)
```
* `\n` works faster than `endl`.
* functions *scanf* and *printf* are usually a bit faster than C++ standard streams. 
```cpp
int a, b;
scanf("%d %d", &a, &b);
printf("%d %d\n", a, b);
```
* To read the whole line from the input (which possibly contains spaces), we can use *getline* function.
```cpp
string s;
getline(cin, s);
```
* If the amount of data is unknown, we can use such a loop.
```cpp
while (cin >> x) {
    // code
}
```
* To read/write from/to files as we are using standard input/output we can use *freopen*.
```cpp
freopen("input.txt", "r", stdin);
freopen("output.txt", "w", stdout);
```
### Working with nubmers
#### Integers
#### Modular arithmetics
Important property of the remainder.
$$
(a + b) \mod m = (a \mod m + b \mod m) \mod m \\
(a − b) \mod m = (a \mod m − b \mod m) \mod m \\
(a \cdot b) \mod m = (a \mod m \cdot b \mod m) \mod m
$$
For example, the following calculates *n!, modulo m*:
```cpp
long long x = 1;
for (int i = 2; i <= n; i++) {
    x = (x*i)%m;
}
cout << x%m << "\n";
```
#### Floating point numbers
Good way to compare floating point numbers is to evaluate their difference.
```cpp
if (abs(a-b) < 1e-9) {
    // a and b are equal
}
```
### Shortening code
#### Type names
Using the command *typedef* it is possible to give a shorter name to a datatype. 
```cpp
typedef long long ll;
long long a = 123456789;
ll b = 123456789;
```
It can also be used with more complex datatypes.
```cpp
typedef vector<int> vi;
typedef pair<int, int> pi;
```
#### Macros
A macro means that certain string in the code will be changed before compilation.
### Mathematics
#### Sum formulas
#### Set theory
#### Logic
#### Functions
#### Logarithms

## 2. Time complexity
## 3. Sorting
## 4. Data structures
## 5. Complete search
## 6. Greedy algorithms
## 7. Dynamic programming
## 8. Amortized analysis
## 9. Range queries
## 10. Bit manipulation

# Graph algorithms
## 11. Basics of graphs
## 12. Graph traversal
## 13. Shortest paths
## 14. Tree algorithms
## 15. Spanning trees
## 16. Directed graphs
## 17. Strong connectivity
## 18. Tree queries
## 19. Paths and circuits
## 20. Flows and cuts

# Advanced topics
## 21. Number theory
## 22. Combinatorics
## 23. Matrices
## 24. Probability
## 25. Game theory
## 26. String algorithms
## 27. Square root algorithms
## 28. Segment trees revisited
## 29. Geometry
## 30. Sweep line algorithms
