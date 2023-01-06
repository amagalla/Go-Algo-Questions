package main

import (
	"fmt"

	"github.com/amagalla/algos/ctci"
)

func main() {

	
	// Implement an algorithm to determine if a string has all unique characters.

	fmt.Println("IsUnique")
	fmt.Println("")

	fmt.Println(1, ctci.IsUnique("abcd")) // true
	fmt.Println(2, ctci.IsUnique("abcdabcd")) // false
	fmt.Println(3, ctci.IsUnique("Hanzi")) // true
	fmt.Println(4, ctci.IsUnique("")) // false

	fmt.Println("______________________")
	fmt.Println("")

	// Check Permutation: Given two strings, write a method to decide if one is
	// as permutation of the other.

	fmt.Println("CheckPermutation")
	fmt.Println("")

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

	fmt.Println("PalindromePermutation")
	fmt.Println("")

	fmt.Println(1, ctci.PalindromePermutation("Tact Coa")) // true
	fmt.Println(2, ctci.PalindromePermutation("T@ac3t   C2o]a")) // true
	fmt.Println(3, ctci.PalindromePermutation("aabbcddee")) // true
	fmt.Println(4, ctci.PalindromePermutation("abcdeabfde")) // false
	fmt.Println(5, ctci.PalindromePermutation("aaaaaa")) // false
	fmt.Println(6, ctci.PalindromePermutation("aaaaaaa")) // true
	fmt.Println(7, ctci.PalindromePermutation("")) // false

	fmt.Println("______________________")
	fmt.Println("")

	/*
	There are three types of edits that can be performed on strings: insert a character,
	remove a character, or replace a character. Given two strings, write a function to
	check if they are one edit (or zero edits) away.

	EXAMPLE

	pale, ple -> true
	pales, pale -> true
	pale, bale -> true
	pale, bake -> false

	*/

	fmt.Println("OneWay")
	fmt.Println("")

	fmt.Println(1, ctci.OneWay("pale", "pale")) // true
	fmt.Println(2, ctci.OneWay("pale", "ple")) // true
	fmt.Println(3, ctci.OneWay("pales", "pale")) // true
	fmt.Println(4, ctci.OneWay("pale", "bale")) // true
	fmt.Println(5, ctci.OneWay("e", "a")) // true
	fmt.Println(6, ctci.OneWay("", "a")) // true
	fmt.Println(7, ctci.OneWay("a", "")) // true
	fmt.Println(8, ctci.OneWay("pale", "bake")) // false
	fmt.Println(9, ctci.OneWay("ke", "bake")) // false
	fmt.Println(10, ctci.OneWay("baaaaaake", "bake")) // false
	fmt.Println(11, ctci.OneWay("bake", "baaaaaaaaaake")) // false

	fmt.Println("______________________")
	fmt.Println("")
	
	
	/*

	Implement a method to peform basic string compression using the counts of
	repeated characters. For example, the string aabcccccaaa would become
	a2b1c5a3. If the "compressed" string would not become smaller than the
	original string, your method should return the original string. You can assume
	the string has only uppercase and lowercase letters (a-z).

	*/

	fmt.Println(1, ctci.StringCompression("aabcccccaaa")) // a2b1c5a3
	fmt.Println(1, ctci.StringCompression("aabccccc")) // a2b1c5
	fmt.Println(1, ctci.StringCompression("a")) // a
	fmt.Println(1, ctci.StringCompression("aa")) // aa

	fmt.Println("______________________")
	fmt.Println("")
}