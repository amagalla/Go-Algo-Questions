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