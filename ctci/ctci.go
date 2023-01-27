package ctci

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// Implement an algorithm to determine if a string has all unique characters.

func IsUnique(str string) bool {
	// check is str is empty
	if len(str) <= 0 {
		return false
	}

	// create a map
	charMap := make(map[string]bool)

	for _, char := range str {
		// check if char is inside charMap
		// return false
		if charMap[string(char)] {
			return false
			// else
			// add char to charMap
		} else {
			charMap[string(char)] = true
		}
	}
	// return true
	return true
}

// Check Permutation: Given two strings, write a method to decide if one is
// as permutation of the other.

func CheckPermutation(str1 string, str2 string) bool {

	// check if length of str1 is not equal to lengh of str2
	// return false
	if len(str1) != len(str2) {
		return false
	}

	// create an empty map
	charMap := make(map[string]int)

	// iterate through str1
	// check if obj has char
	// increment obj char value by 1
	// else
	// add char to obj with value of 1
	for _, char := range str1 {
		if charMap[string(char)] >= 1 {
			charMap[string(char)]++
		} else {
			charMap[string(char)] = 1
		}
	}

	// iterate through str2
	// check if char is in obj
	// decrement obj char value by 1
	// check if obj char value is equal to 0
	// delete map char
	for _, char := range str2 {
		if charMap[string(char)] >= 1 {
			charMap[string(char)]--
			if charMap[string(char)] == 0 {
				delete(charMap, string(char))
			}
		}
	}
	// check if map is not empty
	// return false
	// else
	// return true
	return len(charMap) <= 0
}

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

func PalindromePermutation(str string) bool {
	// check if string is empty
	// return false
	if len(str) <= 0 {
		return false
	}

	// create an empty char map
	charMap := make(map[string]int)

	// create isOdd with value to false
	isOdd := false

	// iterate through str
	// create lower var with value of converted char to lowercase
	// create ascii var with lowercase ascii value

	// check if ascii value is greater than or equal to lowercase a AND less then or eqaul to lowercase Z ascii values
	// check if map has lower as key
	// increment map lower value by 1
	// else
	// assign map lower value to 1
	for _, char := range str {
		lower := strings.ToLower(string(char))
		ascii, _ := utf8.DecodeRuneInString(lower)
		a, _ := utf8.DecodeRuneInString("a")
		z, _ := utf8.DecodeRuneInString("z")

		if ascii >= a && ascii <= z {
			_, ok := charMap[lower]

			if ok {
				charMap[lower]++
			} else {
				charMap[lower] = 1
			}
		}

	}

	// iterate through char map
	// check if char map value is odd
	// check if isOdd is true
	// return false
	// set isOdd with value of true
	for _, element := range charMap {
		if element%2 != 0 {
			if isOdd {
				return false
			}
			isOdd = true
		}
	}
	// return isOdd
	return isOdd
}

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

func OneWay(str1 string, str2 string) bool {
	// check if str2 length is larger than str1 length + 1
	// return false
	if len(str2) > len(str1)+1 {
		return false
	}

	// create an empty map
	charMap := make(map[string]int)

	// create a counter to keep track of number of unused char
	counter := 0

	// iterate through str1
	// check if map has char
	// increment map char value by 1
	// else
	// assign map char val to 1
	for _, char := range str1 {
		_, ok := charMap[string(char)]

		if ok {
			charMap[string(char)]++
		} else {
			charMap[string(char)] = 1
		}
	}

	// iterate through str2
	// check if map has char
	// decrement map char value by 1
	// check if map char value is equal to 0
	// delete char map
	for _, char := range str2 {
		_, ok := charMap[string(char)]

		if ok {
			charMap[string(char)]--
			if charMap[string(char)] <= 0 {
				delete(charMap, string(char))
			}
		} else {
			counter++
		}
	}
	// test if counter value is greater than 1 (means str1 has inserted more than one char)
	if counter > 1 {
		return false
	}

	// iterate through map
	// check if value is greater than 1
	// return false
	for _, value := range charMap {
		if value > 1 {
			return false
		}
	}

	return true
}

/*

	Implement a method to peform basic string compression using the counts of
	repeated characters. For example, the string aabcccccaaa would become
	a2b1c5a3. If the "compressed" string would not become smaller than the
	original string, your method should return the original string. You can assume
	the string has only uppercase and lowercase letters (a-z).

*/

func StringCompression(str string) string {
	// check if string is empty
	// return empty str
	if len(str) == 0 {
		return str
	}

	// create new string var
	// create prev string var
	// create counter var
	var newStr string
	var prev string
	counter := 1

	// iterate through str
	// check if index is 0
	// prev = char
	// couninue

	// check if prev == char
	// increment counter by 1

	// check if prev != char
	// assign newStr with prev + counter
	// assign prev with char
	// assign counter to 0

	// check if index == length of string - 1 && prev == char
	// assign newStr with prev + counter
	for i, val := range str {
		char := string(val)

		if i == 0 {
			prev = char
			continue
		}

		if prev == char {
			counter++
		}

		if prev != char {
			newStr += prev + strconv.Itoa(counter)
			prev = char
			counter = 1
		}

		if i == len(str)-1 && prev == char {
			newStr += prev + strconv.Itoa(counter)
		}
	}

	if len(str) <= len(newStr) || len(str) == 1 {
		return str
	}

	// return newStr
	return newStr
}
