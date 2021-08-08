package lc8stringtointeger_test

import (
	"testing"

	lc8 "github.com/JoshuaCrocker/leetcode/lc8-string-to-integer"
)

// Example 1:
// Input: s = "42"
// Output: 42
// Explanation: The underlined characters are what is read in, the caret is the current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "42" ("42" is read in)
//            ^
// The parsed integer is 42.
// Since 42 is in the range [-2^31, 2^31 - 1], the final result is 42.
func TestMyAtoiExample1(t *testing.T) {
	input := "42"
	expected := 42

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

// Example 2:
// Input: s = "   -42"
// Output: -42
// Explanation:
// Step 1: "   -42" (leading whitespace is read and ignored)
//             ^
// Step 2: "   -42" ('-' is read, so the result should be negative)
//              ^
// Step 3: "   -42" ("42" is read in)
//                ^
// The parsed integer is -42.
// Since -42 is in the range [-2^31, 2^31 - 1], the final result is -42.
func TestMyAtoiExample2(t *testing.T) {
	input := "-42"
	expected := -42

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

// Example 3:
// Input: s = "4193 with words"
// Output: 4193
// Explanation:
// Step 1: "4193 with words" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
//              ^
// The parsed integer is 4193.
// Since 4193 is in the range [-2^31, 2^31 - 1], the final result is 4193.
func TestMyAtoiExample3(t *testing.T) {
	input := "4193 with words"
	expected := 4193

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

// Example 4:
// Input: s = "words and 987"
// Output: 0
// Explanation:
// Step 1: "words and 987" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "words and 987" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "words and 987" (reading stops immediately because there is a non-digit 'w')
//          ^
// The parsed integer is 0 because no digits were read.
// Since 0 is in the range [-2^31, 2^31 - 1], the final result is 0.
func TestMyAtoiExample4(t *testing.T) {
	input := "words and 987"
	expected := 0

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

// Example 5:
// Input: s = "-91283472332"
// Output: -2147483648
// Explanation:
// Step 1: "-91283472332" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "-91283472332" ('-' is read, so the result should be negative)
//           ^
// Step 3: "-91283472332" ("91283472332" is read in)
//                      ^
// The parsed integer is -91283472332.
// Since -91283472332 is less than the lower bound of the range [-2^31, 2^31 - 1], the final result is clamped to -2^31 = -2147483648.
func TestMyAtoiExample5(t *testing.T) {
	input := "-91283472332"
	expected := -2147483648

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample6(t *testing.T) {
	input := "3.14159"
	expected := 3

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample7(t *testing.T) {
	input := " "
	expected := 0

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample8(t *testing.T) {
	input := ""
	expected := 0

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample9(t *testing.T) {
	input := "+1"
	expected := 1

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample10(t *testing.T) {
	input := "20000000000000000000"
	expected := 2147483647

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample11(t *testing.T) {
	input := "-20000000000000000000"
	expected := -2147483648

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestMyAtoiExample12(t *testing.T) {
	input := "  0000000000012345678"
	expected := 12345678

	result := lc8.MyAtoi(input)

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
