package lc12integertoroman_test

import (
	"testing"

	lc12integertoroman "github.com/JoshuaCrocker/leetcode/lc12-integer-to-roman"
)

// Example 1:
// Input: num = 3
// Output: "III"
func TestIntToRomanExample1(t *testing.T) {
	input := 3
	expected := "III"
	result := lc12integertoroman.IntToRoman(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 2:
// Input: num = 4
// Output: "IV"
func TestIntToRomanExample2(t *testing.T) {
	input := 4
	expected := "IV"
	result := lc12integertoroman.IntToRoman(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 3:
// Input: num = 9
// Output: "IX"
func TestIntToRomanExample3(t *testing.T) {
	input := 9
	expected := "IX"
	result := lc12integertoroman.IntToRoman(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 4:
// Input: num = 58
// Output: "LVIII"
// Explanation: L = 50, V = 5, III = 3.
func TestIntToRomanExample4(t *testing.T) {
	input := 58
	expected := "LVIII"
	result := lc12integertoroman.IntToRoman(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 5:
// Input: num = 1994
// Output: "MCMXCIV"
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func TestIntToRomanExample5(t *testing.T) {
	input := 1994
	expected := "MCMXCIV"
	result := lc12integertoroman.IntToRoman(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
