package main

import (
	"fmt"

	"github.com/amagalla/algos/ctci"
)

func main() {

	
	// Implement an algorithm to determine if a string has all unique characters.

	fmt.Println(ctci.IsUnique("abcd")) // true
	fmt.Println(ctci.IsUnique("abcdabcd")) // false
}