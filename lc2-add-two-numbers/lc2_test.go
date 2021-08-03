package lc2addtwonumbers_test

import (
	"reflect"
	"testing"

	lc2 "github.com/JoshuaCrocker/leetcode/lc2-add-two-numbers"
)

// TestCreateLinkedList
func TestCreateLinkedList(t *testing.T) {
	data := []int{2, 4, 3}

	result := lc2.CreateLinkedList(data)

	if result.Val != 2 {
		t.Errorf("first node is not 2, got %d", result.Val)
	}

	if result.Next.Val != 4 {
		t.Errorf("second node is not 4, got %d", result.Next.Val)
	}

	if result.Next.Next.Val != 3 {
		t.Errorf("third node is not 3, got %d", result.Next.Next.Val)
	}

	if result.Next.Next.Next != nil {
		t.Errorf("third node not terminated correctly")
	}
}

// TestAddTwoNumbersExample1
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
func TestAddTwoNumbersExample1(t *testing.T) {
	l1data := []int{2, 4, 3}
	l2data := []int{5, 6, 4}
	expectedData := []int{7, 0, 8}

	l1 := lc2.CreateLinkedList(l1data)
	l2 := lc2.CreateLinkedList(l2data)
	expected := lc2.CreateLinkedList(expectedData)

	result := lc2.AddTwoNumbers(l1, l2)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestAddTwoNumbersExample2
//
// Input: l1 = [0], l2 = [0]
// Output: [0]
func TestAddTwoNumbersExample2(t *testing.T) {
	l1data := []int{0}
	l2data := []int{0}
	expectedData := []int{0}

	l1 := lc2.CreateLinkedList(l1data)
	l2 := lc2.CreateLinkedList(l2data)
	expected := lc2.CreateLinkedList(expectedData)

	result := lc2.AddTwoNumbers(l1, l2)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestAddTwoNumbersExample3
//
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
func TestAddTwoNumbersExample3(t *testing.T) {
	l1data := []int{9, 9, 9, 9, 9, 9, 9}
	l2data := []int{9, 9, 9, 9}
	expectedData := []int{8, 9, 9, 9, 0, 0, 0, 1}

	l1 := lc2.CreateLinkedList(l1data)
	l2 := lc2.CreateLinkedList(l2data)
	expected := lc2.CreateLinkedList(expectedData)

	result := lc2.AddTwoNumbers(l1, l2)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
