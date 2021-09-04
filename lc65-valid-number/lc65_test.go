package lc65validnumber_test

import (
	"testing"

	lc65 "github.com/JoshuaCrocker/leetcode/lc65-valid-number"
)

// Example 1:
//
// Input: s = "0"
// Output: true
func TestIsNumberExampleZero(t *testing.T) {
	input := "0"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 2:
//
// Input: s = "e"
// Output: false
func TestIsNumberExampleJustE(t *testing.T) {
	input := "e"
	expected := false

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 3:
//
// Input: s = "."
// Output: false
func TestIsNumberExampleJustDecimalPoint(t *testing.T) {
	input := "."
	expected := false

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 4:
//
// Input: s = ".1"
// Output: true
func TestIsNumberExampleDecimalNoLeadingZero(t *testing.T) {
	input := ".1"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 5:
//
// Input: s = "0.1"
// Output: true
func TestIsNumberExampleDecimalWithLeadingZero(t *testing.T) {
	input := "0.1"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 6:
//
// Input: s = "3.14159"
// Output: true
func TestIsNumberExampleDecimalPi(t *testing.T) {
	input := "3.14159"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 6:
//
// Input: s = "12345"
// Output: true
func TestIsNumberExampleInteger(t *testing.T) {
	input := "12345"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 7:
//
// Input: s = "-12345"
// Output: true
func TestIsNumberExampleNegativeInteger(t *testing.T) {
	input := "-12345"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 8:
//
// Input: s = "-3.14159"
// Output: true
func TestIsNumberExampleNegativeDecimal(t *testing.T) {
	input := "-3.14159"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 9:
//
// Input: s = "+12345"
// Output: true
func TestIsNumberExamplePositiveInteger(t *testing.T) {
	input := "+12345"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 10:
//
// Input: s = "+0.1"
// Output: true
func TestIsNumberExamplePositiveDecimal(t *testing.T) {
	input := "+0.1"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 11:
//
// Input: s = "+.1"
// Output: true
func TestIsNumberExamplePositiveDecimalNoLeadingZero(t *testing.T) {
	input := "+.1"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 12:
//
// Input: s = "+.1"
// Output: true
func TestIsNumberExampleNegativeDecimalNoLeadingZero(t *testing.T) {
	input := "-.1"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 13:
//
// Input: s = "2e2"
// Output: true
func TestIsNumberExampleExponential(t *testing.T) {
	input := "2e2"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 13:
//
// Input: s = "+2e2"
// Output: true
func TestIsNumberExampleExponential2(t *testing.T) {
	input := "+2e2"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 14:
//
// Input: s = "-2e2"
// Output: true
func TestIsNumberExampleExponential3(t *testing.T) {
	input := "-2e2"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 15:
//
// Input: s = "2e+2"
// Output: true
func TestIsNumberExampleExponential4(t *testing.T) {
	input := "2e+2"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 16:
//
// Input: s = "2e-2"
// Output: true
func TestIsNumberExampleExponential5(t *testing.T) {
	input := "2e-2"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 17:
//
// Input: s = "3.14e5"
// Output: true
func TestIsNumberExampleExponential6(t *testing.T) {
	input := "3.14e5"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

// Example 18:
//
// Input: s = "2E2"
// Output: true
func TestIsNumberExampleExponential7(t *testing.T) {
	input := "2E2"
	expected := true

	result := lc65.IsNumber(input)

	if result != expected {
		t.Errorf("%s: got %v, expected %v", input, result, expected)
	}
}

func TestIsNumberAllValidExamples(t *testing.T) {
	testData := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}

	for _, data := range testData {
		t.Run(data, func(t *testing.T) {
			expected := true

			result := lc65.IsNumber(data)

			if result != expected {
				t.Errorf("got %v, expected %v", result, expected)
			}
		})
	}
}

func TestIsNumberAllInvalidExamples(t *testing.T) {
	testData := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}

	for _, data := range testData {
		t.Run(data, func(t *testing.T) {
			expected := false

			result := lc65.IsNumber(data)

			if result != expected {
				t.Errorf("got %v, expected %v", result, expected)
			}
		})
	}
}
