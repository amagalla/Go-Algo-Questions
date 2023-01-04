package main

import (
	"fmt"

	"github.com/amagalla/algos/ctci"
)

func main() {

	
	// Implement an algorithm to determine if a string has all unique characters.

	fmt.Println(1, ctci.IsUnique("abcd")) // true
	fmt.Println(2, ctci.IsUnique("abcdabcd")) // false
	fmt.Println(3, ctci.IsUnique("Hanzi")) // true

	fmt.Println("______________________")
	fmt.Println("")

	// Check Permutation: Given two strings, write a method to decide if one is
	// as permutation of the other.

	fmt.Println(1, ctci.CheckPermutation("bad", "dab")) // true
	fmt.Println(2, ctci.CheckPermutation("like", "ilek")) // true
	fmt.Println(3, ctci.CheckPermutation("abcd", "badc")) // true
	fmt.Println(4, ctci.CheckPermutation("eeaabb", "aabbcc")) // false
	fmt.Println(5, ctci.CheckPermutation("abcde", "abc")) // false
	fmt.Println(6, ctci.CheckPermutation("abcd", "abcdef")) // false
	fmt.Println(7, ctci.CheckPermutation("abcd", "abde")) // false

	fmt.Println("______________________")
	fmt.Println("")
}