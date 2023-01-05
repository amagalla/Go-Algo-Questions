package ctci

import (
	"strings"
	"unicode/utf8"
)

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
			return false;
			// else
			// add char to charMap
		} else {
			charMap[string(char)] = true
		}
	}
	// return true
	return true
}

func CheckPermutation(str1 string, str2 string) bool {

	// check if length of str1 is not equal to lengh of str2
		// return false
	if (len(str1) != len(str2)) {
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
			if (charMap[string(char)] == 0) {
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

		if (ascii >= a && ascii <= z) {
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
		if (element % 2 != 0) {
			if isOdd {
				return false
			}
			isOdd = true
		}
	}
	// return isOdd
	return isOdd
}