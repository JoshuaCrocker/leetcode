package lc7reverseinteger

import (
	"strconv"
	"strings"
)

// Reverse takes a 32-bit integer x and returns x with its digits reversed.
//
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1],
// then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
// Constraints:
//     -231 <= x <= 231 - 1
func Reverse(x int) int {
	str := strconv.Itoa(x)

	parts := strings.Split(str, "")

	negative := false

	if parts[0] == "-" {
		negative = true
		_, parts = parts[0], parts[1:]
	}

	reversed := make([]string, len(parts))

	for index, value := range parts {
		reversed[len(parts)-index-1] = value
	}

	joined := strings.Join(reversed, "")

	output, _ := strconv.Atoi(joined)

	if negative {
		output = output * -1
	}

	if output < -2147483648 || output > 2147483647 {
		return 0
	}

	return output
}
