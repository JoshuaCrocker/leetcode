package lc20validparenthesis

import "strings"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
// * Open brackets must be closed by the same type of brackets.
// * Open brackets must be closed in the correct order.
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		// If we have an odd-length string, there has to be an unclosed bracket.
		return false
	}

	opening := []string{"(", "{", "["}
	closing := []string{")", "}", "]"}

	parts := strings.Split(s, "")
	stack := []string{}

	for _, part := range parts {
		if index := getIndex(part, opening); index != -1 {
			stack = append(stack, part)
		} else if index := getIndex(part, closing); index != -1 {
			if stack[len(stack)-1] == opening[index] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func getIndex(needle string, haystack []string) int {
	for i, val := range haystack {
		if needle == val {
			return i
		}
	}

	return -1
}
