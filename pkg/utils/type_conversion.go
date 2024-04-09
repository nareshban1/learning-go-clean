package utils

import "strconv"

// Converts a string to an unsigned integer.
// If the input string is empty, it returns 0 without an error.
func StringToUInt(str string) (uint, error) {
	if str == "" {
		return 0, nil
	}

	parsedInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return uint(parsedInt), nil
}
