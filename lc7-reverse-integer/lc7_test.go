package lc7reverseinteger_test

import (
	"testing"

	lc7 "github.com/JoshuaCrocker/leetcode/lc7-reverse-integer"
)

// Input: x = 123
// Output: 321
func TestReverseExample1(t *testing.T) {
	input := 123
	expected := 321

	if lc7.Reverse(input) != expected {
		t.Errorf("got %d, expected %d", lc7.Reverse(input), expected)
	}
}

// Input: x = -123
// Output: -321
func TestReverseExample2(t *testing.T) {
	input := -123
	expected := -321

	if lc7.Reverse(input) != expected {
		t.Errorf("got %d, expected %d", lc7.Reverse(input), expected)
	}
}

// Input: x = 120
// Output: 21
func TestReverseExample3(t *testing.T) {
	input := 120
	expected := 21

	if lc7.Reverse(input) != expected {
		t.Errorf("got %d, expected %d", lc7.Reverse(input), expected)
	}
}

// Input: x = 0
// Output: 0
func TestReverseExample4(t *testing.T) {
	input := 0
	expected := 0

	if lc7.Reverse(input) != expected {
		t.Errorf("got %d, expected %d", lc7.Reverse(input), expected)
	}
}

// Test outside of 32-bit range
func TestReverseExample5(t *testing.T) {
	input := 2147483647
	expected := 0

	if lc7.Reverse(input) != expected {
		t.Errorf("got %d, expected %d", lc7.Reverse(input), expected)
	}
}
