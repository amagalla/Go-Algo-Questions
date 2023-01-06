package ctci

import (
	"strconv"
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

func OneWay(str1 string, str2 string) bool {
	// check if str2 length is larger than str1 length + 1
		// return false
	if len(str2) > len(str1) + 1 {
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
			if (charMap[string(char)] <= 0) {
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

func StringCompression(str string) string {
	// create a new empty str
	var newStr string

	// create a counter
	counter := 1

	// create prev var
	var prev string

	// iterate through str
		// inrement counter by 1
		// check if char is not equal to char + 1 value
			// assign char to newStr + counter
			// set counter to 0
	for i, val := range str {
		char := string(val)

		// check if prev is empty
			// assign prev with value of char
		if prev == "" {
			prev = char
		}

		// check if prev is equal to char and i is not equal to 0 (so first iteration isn't counted)
		if prev == char && i != 0 {
			counter++
		}

		// check if prev is not equal to char
			// add value of prev with counter to newStr
			// reassign prev to char
			// reset counter to 1
		if prev != char {
			newStr += prev + strconv.Itoa(counter)
			prev = char
			counter = 1
		}
		
		// check if prev is qual to char and i is equal to str length (to get last char index value)
			// add char with counter to newStr
		if prev == char && i + 1 == len(str) {
			newStr += prev + strconv.Itoa(counter)
		}
	}

	// check if original string is less than or equal to newStr
		// return original str
	if len(str) <= len(newStr) {
		return str
	}

	// return newStr
	return newStr
}