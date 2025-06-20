package main 

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int // declared a pointer
	count2 := new(int) // created a variable
	countTemp := 5
	count3 := &countTemp // created a pointer from an existing var
	t := &time.Time{}
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)
}