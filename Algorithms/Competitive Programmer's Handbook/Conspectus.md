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
```cpp
#define F first
#define S seconds
#define PB push_back
#define MP make_pair
// After this, the code 
v.push_back(make_pair(y1, x1));
v.push_back(make_pair(y2, x2));
int d = v[i].first+v[i].second;
// can be shortened as follows:
v.PB(MP(y1, x1));
v.PB(MP(y2, x2));
int d = v[i].F + v[i].S;
```
A macro can also have parameters which makes it possible to shorten loops and other structures.
```cpp
#define REP(i, a, b) for (int i = a; i <= b; i++)
// Afther this, the code
for (int i = 1; i <= n; i++) {
    search(i);
}
can be shortened as follows:
REP(i, 1, b) {
    search(i);
}
```
### Mathematics
#### Sum formulas
$$
\sum_{x=1}^n x^k = 1^k + 2^k + 3^k + \dots + n^k,
$$

where \( k \) is a positive integer, has a closed-form formula that is a **polynomial of degree \( k + 1 \)**.

For example:

- When \( k = 1 \):

$$
\sum_{x=1}^n x = 1 + 2 + 3 + \dots + n = \frac{n(n + 1)}{2}
$$

- When \( k = 2 \):

$$
\sum_{x=1}^n x^2 = 1^2 + 2^2 + 3^2 + \dots + n^2 = \frac{n(n + 1)(2n + 1)}{6}
$$
* *There is a general formula for arbitrary k, but it is too complex to be mentioned here*.

An **arithmetic progression** is a sequence of numbers where the difference between any two consecutive numbers is constant.
- **\( i-th \) term:**

$$
a_i = a_1 + (i - 1)d
$$

- **Sum of first \( n \) terms:**

$$
S_n = \frac{n}{2} \big(2a_1 + (n - 1)d \big) = \frac{n}{2} (a_1 + a_n)
$$

A **geometric progression** is a sequence of numbers where the ratio between any two consecutive numbers is constant. 

**i-th** Term
$$
a_i = a_1 \cdot r^{i - 1}
$$

Sum of First **n** Terms

If r < 1:

$$
S_n = a_1 \cdot \frac{1 - r^n}{1 - r}
$$

Or (equivalently, if \( r > 1 \)):

$$
S_n = a_1 \cdot \frac{r^n - 1}{r - 1}
$$

Alternative form using first and last term:

$$
S_n = \frac{br - a}{r - 1}
$$

A **harmonic sum** is the sum of the reciprocals of the first n natural numbers. 
$$
H_n = \sum_{k=1}^n \frac{1}{k} = 1 + \frac{1}{2} + \frac{1}{3} + \cdots + \frac{1}{n}
$$

#### Set theory
#### Logic
#### Functions
#### Logarithms
$$
\log_k (ab) = \log_k (a) + \log_k (b),
$$

$$
\log_k (x^n) = n \cdot \log_k (x).
$$

$$
\log_k \left(\frac{a}{b}\right) = \log_k (a) - \log_k (b).
$$

$$
\log_u (x) = \frac{\log_k (x)}{\log_k (u)}.
$$


## 2. Time complexity
The **time complexity** of an algorithm estimates how much time the algorithms will use for some input. The idea is to represent the efficiency as a function whose parameter is the size of the input, and the output is the amount of steps we need to take to accomplish the computational task. 
### Calculation rules
The time complexity is denoted as $O(func(n))$, where ***n*** denotes the input size.
#### Loops
#### Order of magnitude
Time complexity doesn't give the exact number of loop iterations - it only describes how the number of iterations *grows with input size*.
For example, all these are O(n):
```cpp
for (int i = 1; i <= 3*n; i++) {
// code
}
for (int i = 1; i <= n+5; i++) {
// code
}
for (int i = 1; i <= n; i += 2) {
// code
}
```
#### Phases
If the algorithm consists of consecutive phases, the total time complexity is the largest time complexity of a single phase.
#### Several variables
Sometimes the time complexity depends on several factors. In this case, the time
complexity formula contains several variables.
For example, the time complexity of the following code is O(nm):
```cpp
for (int i = 1; i <= n; i++) {
    for (int j = 1; j <= m; j++) {
    // code 
    }
}
```
#### Recursion
The time complexity of a recursive function depends on:
- number of times it called
- time complexity of a single call
### Complexity classes
### Complexity classes
- $O(1)$ – **constant-time**, the algorithm does not depend on the input size.
- $O(\log n)$ - **logarithmic**, iften halves the input size at each step.
- $O(\sqrt{n})$ - **square root algorithm**
- $O(n)$ - **linear**, goes throught input a constant amount of times.
- $O(n \log n)$ - often indicates that the algorithm sorts the input.
- $O(n^2)$ - **quadratic**, often contains two nested loops.
- $O(n^3)$ - **cubic**, often contains three nested loops.
- $O(2^n)$ - often indicates that the algorithm iterates through all subsets of the input elements.
- $O(n!)$ - often indicates that the algorithm iterates through all permutations of the input element. 

An algorythm is **polynomial** if its time complexity is at most $O(n^k)$, where $k$ is a *constant*.

**NP-hard** problems - problems, for which no polynomial algorithm is known.

### Estimating efficiency
It is important to remember that a time complexity is only an estimate of efficiency, because it hides *constant factors*.
An Algorithm that runs $O(n)$ time may perofrm *n/2* or *5n* operations. 

### Maximum subarray sum
Let's discuss a classic problem that has $O(n^3)$, $O(n^2)$, $O(n)$ solutions.

Problem: Given an array of **n** numbers, our task is to caculate the maximum ***subarray sum***, i.e., the largest possible sum of a sequence of consecutive values in the array. 
#### Algorithm 1 (Brute Force - Triple Loop)
```cpp
int best = 0;
for (int a = 0; a < n; a++) {
    for (int b = a; b < n; b++) {
        int sum = 0;
        for (int k = a; k <= b; k++) {
            sum += array[k];
        }
        best = max(best,sum);
    }
}
cout << best << "\n";
```
#### Algorithm 2 (Prefix Sum Optimixation - Double Loop)
```cpp
int best = 0;
for (int a = 0; a < n; a++) {
    int sum = 0;
    for (int b = a; b < n; b++) {
        sum += array[b];
        best = max(best,sum);
    }
}
cout << best << "\n";
```
#### Algorithm 3 (Kabane's algorithm - Single Loop)
```cpp
int best = 0, sum = 0;
for (int k = 0; k < n; k++) {
    sum = max(array[k],sum+array[k]);
    best = max(best,sum);
}
cout << best << "\n";
```

## 3. Sorting
**Sorting** is a fundamental algorithm design problem. 
Many efficient algorithms use sorting as a subroutine. 
THe efficient general sorting algorithms work in $O(nlogn)$ time, as many algorithms that use sorting as a subroutine.
### Sorting theory
#### $O(n^2)$ algorithms
Such algorithms are usually simple and consist of two nested loops.
**Bubble sort** is a famous one. 
The algorithm consisnts of *n* rounds.
On each round, algorithm iterates through the elements, and swaps all consecutive elements, which are in the wron order.
After *k* rounds the *k* largest elements will be in their correct positions. 
```cpp
for (int i = 0; i < n; i++) {
    for (int j = 0; j < n-1; j++) {
        if (array[j] > array[j+1]) {
            swap(array[j], array[j+1]);
        }
    }
}
```
#### Inversions
**Inversion** - a pair of consecutive elements, positioned in the wrong order.
Swapping a pair of consecutive elements that are in the wrong order removes exactly one inversion. 
#### $O(nlogn)$ algorithms
We can implement such algorithms if we are not limited to swapping consecutive elements.
One such algorithm is **merge sort**, which is *based on recursion*.
How it works:
1. If $a = b$, the subarray is already sorted, do nothing.
2. Calculate the position of the middle element: $k = (a+b)/2$.
3. Recursively sort the subarray $array[a...k]$.
4. Rcursively sort the subarray $array[k+1...b]$.
5. *Merge* the sorted subarrays $array[a...k]$ and $array[k+1...b]$ into a sorted subarray $array[a...b]$.
Algorithm is efficient because it halves the size of the subarray at each step.
The recursion consists of $O(logn)$ levels and processing each level takes $O(n)$ time, merging subarrays is linear time.
#### Sorting lower bound
In short, no comparison-based algorithm can do better that $n \log n$.
#### Counting sort
**Counting sort** is an efficient non-comparison-based sorting algorithms.

Its complexity is $O(n)$.

The main idea is that the algorithm iterates through all array elements, and whenever it finds an element **X**, it increments the value of the *bookkeeping* array with index of **X** (barray[X]++). In the end we can restore the sorted sequence.  
The only *limitation* is that the distance between the smallest and the biggest element should not be too large. 
### Sorting in C++
It is almost always better to use a standard library function for sorting.

Sorting a vector:
```cpp
vector<int> v = {4, 2, 5, 3, 5, 8, 3}
sort(v.begin(), v.end());
sort(v.rbegin(), v.rend()) //reverse order
```
We can also sort an ordinary array:
```cpp
int n = 7
int a[] = {4, 2, 5, 3, 8, 3};
sort(a, a+n);
```
And also a string:
```cpp
string s = "monkey";
sort(s.begin(), s.end());
```
#### Comparison operators
The function *sort* requires that a **comparison operator** is defined for the data type of the elements to be sorted. 
We can also define custom comparison operators.
#### User-defined structs
User-defined structs do not have a comparison operator automatically, so we have to define it.
```cpp
struct P {
    int x, y;
    bool operator<(const P &p) const {
        if (x != p.x) return x < p.x;
        else return y < p.y;
    }
};
```
#### Comparison functions
We can also define an external **comparison function** o the sort function as a callback function. 
```cpp
bool comp(string a, string b) {
if (a.size() != b.size()) return a.size() < b.size();
return a < b;
}
// now we can use
sort(v.begin(), v.end(), comp);
```

### Binary search
#### Method 1
#### Method 2
#### C++ functions
#### Finding the smallest solutions 
#### Finding the maximum value 


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
