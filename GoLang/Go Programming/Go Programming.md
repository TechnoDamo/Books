# Table of Contents

1. [Variables and Operators](#1-variables-and-operators)
2. [Command and Control](#2-command-and-control)
3. [Core Types](#3-core-types)
4. [Complex Types](#4-complex-types)
5. [Functions - Reduce, Reuse, and Recycle](#5-functions---reduce-reuse-and-recycle)
6. [Don't Panic! Handle Your Errors](#6-dont-panic-handle-your-errors)
7. [Interfaces](#7-interfaces)
8. [Generic Algorithm Superpowers](#8-generic-algorithm-superpowers)
9. [Using Go Modules to define a project](#9-using-go-modules-to-define-a-project)
10. [Packages Keep Projects Manageable](#10-packages-keep-projects-manageable)
11. [Bug-Busting Debugging Skills](#11-bug-busting-debugging-skills)
12. [About Time](#12-about-time)
13. [Programming from the Command Line](#13-programming-from-the-command-line)
14. [File and Systems](#14-file-and-systems)
15. [SQL and Databases](#15-sql-and-databases)
16. [Web Servers](#16-web-servers)
17. [Using the Go HTTP client](#17-using-the-go-http-client)
18. [Concurrent Work](#18-concurrent-work)
19. [Testing](#19-testing)
20. [Using Go Tools](#20-using-go-tools)
21. [Go in the Cloud](#21-go-in-the-cloud)

# 1. Variables and Operators
  - Go is compliant with memory safety and channel-based concurrency.
  - Go is compiled into a self-contained executable, which can run everywhere.
  - Go has a statically typed and type-safe memory with a garbage collector that automates memory managment.

## What does Go look like?
  All go files must start with package declarations.
    
  ```go
  import (
    "pckgname"
    ...
  )
  ```

  - All files in the same dir are considered part of the same package.
  - The text or strings in Go support multi-type UTG-8 encoding,making them safe for any language.
  - Go functions may recieve and return from 0 to many variables.
  - main() functions is an entry point of your Go code.
    
  ```go
  // Seed random number generator using the current time
  rand.Seed(time.Now().UnixNano())
  // Generate a random number in the range of our list
  index := rand.Intn(len(helloList))
  ```

## Declaring variables
### Declaring a variable using var
```go
var foo string = "bar
```
### Declaring multiple variables at once with var
```go
var (
  name1 string = "5"
  name2 int = 5
  name3 bool = true
)
```
### Skipping the type or value when declaring variables
In real-world code it's not common to use the full var notation.
### Type inference gone wrong
There are times when we need to use full var declaration because Go isn't able to correctly guess the right type.
### Short variable declaration
```go
variable := 123
```
### Declaring multiple variables with a short variable declaration
```go
Debug, LogLevel, startUpTime := false, "info", time.Now()
```
### Using var to declare multiple variables in one line 
```go
func getConfig() (bool, string, time.Time) {
  return false, "info", time.Now()
}
// Type only
var start, middle, end float32
// Mixed type (initial value provided)
var name, left, right, top, bottom = "one", 1, 1.4, 2, 2.5
// Also works with functions
var Debug, LogLevel, startUpTime = getConfig()
```
## Changing the value of a variable 
```go
variable := 5
variable = 10
v1, v2, v3 := 1, "2", 3.0
v1, v2, v3 = 3, "4", 5.0
```
## Operators
- **Arithmetic operators**  
- **Comparison operators**  
- **Logical operators**  
- **Address operators**  
- **Receive operators**  
### Shorthand operations
--, ++, +=, -=
### Comparing values
All the trevial ones.
### Zero values
bool - false
Numbers (ints and floats) - 0
String - "
Other base types - nil

## Value versus pointer
When we pas primitive data types, arrays and structs to a function, Go makes copies of their values and passes them instead. This approach may potentially uses much memory.
We can also pass ***pointers*** instead, but this approach potentially uses CPU a lot.
**Pointer** is not the value itself, but an address of that value.
Go has two types of memory, ***heap-based*** and ***stack-based***.
*Stack* is used for copying values, *heap* for pointers. 
Working out whether a value needs to be put on the heap is called **escape analysis**
### Getting a pointer

### Getting a value from a pointer
### Function design with pointers

## Enums 


# 2. Command and Control

# 3. Core Types

# 4. Complex Types

# 5. Functions - Reduce, Reuse, and Recycle

# 6. Don't Panic! Handle Your Errors

# 7. Interfaces

# 8. Generic Algorithm Superpowers

# 9. Using Go Modules to define a project

# 10. Packages Keep Projects Manageable

# 11. Bug-Busting Debugging Skills

# 12 About Time

# 13. Programming form the COmmand Line

# 14. File and Systems

# 15. SQL and Databases

# 16. Web Servers

# 17. Using the Go HTTP client

# 18. Concurrent Work

# 19. Testing 

# 20. Using Go Tools

# 21. Go in the Cloud

