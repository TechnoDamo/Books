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
    Go is compliant with memory safety and channel-based concurrency.
    Go is compiled into a self-contained executable, which can run everywhere.
    Go has a statically typed and type-safe memory with a garbage collector that automates memory managment.

## What does Go look like?
    All go files must start with package declarations.
    '''go
    import (
      "pckgname"
      ...
    )
    '''
    All files in the same dir are considered part of the same package.
    The text or strings in Go support multi-type UTG-8 encoding, making them safe for any language.
    Go functions may recieve and return 0 to many variables.
    main() functions is an entry point of your Go code.
    '''go
    // Seed random number generator using the current time
    rand.Seed(time.Now().UnixNano())
    // Generate a random number in the range of our list
    index := rand.Intn(len(helloList))
    '''

## Declaring variables

## Changing the value of a variable 

## Operators

## Value versus pointer

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

