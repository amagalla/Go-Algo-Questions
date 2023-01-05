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
	fmt.Println(4, ctci.IsUnique("")) // false

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

	/*

	Given a string, write a function to check if it is a permutation
	of a palindrome. A palindrome is a word or phrase that is the
	same forward and backwards. A permutation is a rearrangement
	of letters. The palindrome does not need to be limited to just
	dictionary words.
	You can ignore casing and non-letter characters.

	EXAMPLE

	Input: Tact Coa

	Output: True (permutation: "taco cat", "atco cta", etc).

	*/

	fmt.Println(1, ctci.PalindromePermutation("Tact Coa")) // true
	fmt.Println(2, ctci.PalindromePermutation("T@ac3t   C2o]a")) // true
	fmt.Println(3, ctci.PalindromePermutation("aabbcddee")) // true
	fmt.Println(4, ctci.PalindromePermutation("abcdeabfde")) // false
	fmt.Println(5, ctci.PalindromePermutation("aaaaaa")) // false
	fmt.Println(6, ctci.PalindromePermutation("aaaaaaa")) // true
	fmt.Println(7, ctci.PalindromePermutation("")) // false

	fmt.Println("______________________")
	fmt.Println("")
}