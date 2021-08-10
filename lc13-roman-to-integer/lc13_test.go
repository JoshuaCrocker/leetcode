package lc13romantointeger_test

import (
	"testing"

	lc13romantointeger "github.com/JoshuaCrocker/leetcode/lc13-roman-to-integer"
)

// Example 1:
// Input: s = "III"
// Output: 3
func TestRomanToIntExample1(t *testing.T) {
	input := "III"
	expected := 3
	result := lc13romantointeger.RomanToInt(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 2:
// Input: s = "IV"
// Output: 4
func TestRomanToIntExample2(t *testing.T) {
	input := "IV"
	expected := 4
	result := lc13romantointeger.RomanToInt(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 3:
// Input: s = "IX"
// Output: 9
func TestRomanToIntExample3(t *testing.T) {
	input := "IX"
	expected := 9
	result := lc13romantointeger.RomanToInt(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 4:
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
func TestRomanToIntExample4(t *testing.T) {
	input := "LVIII"
	expected := 58
	result := lc13romantointeger.RomanToInt(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 5:

// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func TestRomanToIntExample5(t *testing.T) {
	input := "MCMXCIV"
	expected := 1994
	result := lc13romantointeger.RomanToInt(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
