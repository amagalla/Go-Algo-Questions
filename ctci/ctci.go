package ctci

func IsUnique(str string) bool {
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