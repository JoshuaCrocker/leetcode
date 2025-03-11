package lc8stringtointeger

import (
	"errors"
	"math"
	"strings"
)

const Minimum32bitInteger int32 = -2147483648
const Maximum32bitIntger int32 = 2147483647

// MyAtoi converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).
//
// The algorithm for myAtoi(string s) is as follows:
//
// Read in and ignore any leading whitespace.
//     1. Check if the next character (if not already at the end of the string) is '-' or '+'.
//        Read this character in if it is either. This determines if the final result is negative
//        or positive respectively. Assume the result is positive if neither is present.
//     2. Read in next the characters until the next non-digit charcter or the end of the input is
//        reached. The rest of the string is ignored.
//     3. Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were
//        read, then the integer is 0. Change the sign as necessary (from step 2).
//     4. If the integer is out of the 32-bit signed integer range [-2^31, 2^31 - 1], then clamp the
//        integer so that it remains in the range. Specifically, integers less than -2^31 should be
//        clamped to -2^31, and integers greater than 2^31 - 1 should be clamped to 2^31 - 1.
//     5. Return the integer as the final result.
//
// Note:
//     Only the space character ' ' is considered a whitespace character.
//     Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, "")

	if len(parts) == 0 {
		return 0
	}

	// 1. Check if the next charcter is '-' or '+'.
	multiplier := 1

	if parts[0] == "-" {
		multiplier = -1
		parts = parts[1:]
	} else if parts[0] == "+" {
		multiplier = 1
		parts = parts[1:]
	}

	// 2. Read in next the characters until the next non-digit charcter or the end of the input is
	//    reached. The rest of the string is ignored.
	digits := []int{}

	for _, char := range parts {
		converted, err := convertToDigit(char)

		if err != nil {
			break
		} else {
			// We got a digit
			digits = append(digits, converted)
		}
	}

	// 2.1. Remove leading zeroes
	leadingZeroes := 0

	for _, digit := range digits {
		if digit == 0 {
			leadingZeroes++
		} else {
			break
		}
	}

	digits = digits[leadingZeroes:]

	// 3. Convert the digits to an integer
	output := 0

	if len(digits) > 10 {
		// Over the size of the max/min
		output = int(Maximum32bitIntger) + 1
	} else {
		for index, digit := range digits {
			factor := math.Pow10(len(digits) - index - 1)
			output = output + (digit * int(factor))
		}
	}

	output = output * multiplier

	// 4. Clamp to 32-bit range
	if output > int(Maximum32bitIntger) {
		output = int(Maximum32bitIntger)
	}

	if output < int(Minimum32bitInteger) {
		output = int(Minimum32bitInteger)
	}

	return output
}

func convertToDigit(s string) (int, error) {
	switch s {
	case "0":
		return 0, nil
	case "1":
		return 1, nil
	case "2":
		return 2, nil
	case "3":
		return 3, nil
	case "4":
		return 4, nil
	case "5":
		return 5, nil
	case "6":
		return 6, nil
	case "7":
		return 7, nil
	case "8":
		return 8, nil
	case "9":
		return 9, nil
	}

	return 0, errors.New("undefined digit")
}
