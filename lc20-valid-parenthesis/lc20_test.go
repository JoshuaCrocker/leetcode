package lc20validparenthesis_test

import (
	"testing"

	lc20validparenthesis "github.com/JoshuaCrocker/leetcode/lc20-valid-parenthesis"
)

// Example 1:
//
// Input: s = "()"
// Output: true
func TestIsValidExample1(t *testing.T) {
	input := "()"
	expected := true
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 2:
//
// Input: s = "()[]{}"
// Output: true
func TestIsValidExample2(t *testing.T) {
	input := "()[]{}"
	expected := true
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 3:
//
// Input: s = "(]"
// Output: false
func TestIsValidExample3(t *testing.T) {
	input := "(]"
	expected := false
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 4:
//
// Input: s = "([)]"
// Output: false
func TestIsValidExample4(t *testing.T) {
	input := "([)]"
	expected := false
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

// Example 5:
//
// Input: s = "{[]}"
// Output: true
func TestIsValidExample5(t *testing.T) {
	input := "{[]}"
	expected := true
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestIsValidOddLengthString(t *testing.T) {
	input := "{[}"
	expected := false
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestIsValidDoubleOpening(t *testing.T) {
	input := "{{"
	expected := false
	result := lc20validparenthesis.IsValid(input)

	if result != expected {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
