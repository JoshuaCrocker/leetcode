package lc4medianoftwosortedarrays_test

import (
	"testing"

	lc4 "github.com/JoshuaCrocker/leetcode/lc0004-median-of-two-sorted-arrays"
)

// Example 1:
//
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
func TestFindMedianSortedArraysExample1(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{2}
	expected := 2.0

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}

// Example 2:
//
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
func TestFindMedianSortedArraysExample2(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	expected := 2.5

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}

// Example 3:
//
// Input: nums1 = [0,0], nums2 = [0,0]
// Output: 0.00000
func TestFindMedianSortedArraysExample3(t *testing.T) {
	input1 := []int{0, 0}
	input2 := []int{0, 0}
	expected := 0.0

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}

// Example 4:
//
// Input: nums1 = [], nums2 = [1]
// Output: 1.00000
func TestFindMedianSortedArraysExample4(t *testing.T) {
	input1 := []int{}
	input2 := []int{1}
	expected := 1.0

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}

// Example 5:
//
// Input: nums1 = [2], nums2 = []
// Output: 2.00000
func TestFindMedianSortedArraysExample5(t *testing.T) {
	input1 := []int{2}
	input2 := []int{}
	expected := 2.0

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}

// Example 6:
//
// Input: nums1 = [1, 3], nums2 = [2, 7]
// Output: 2.50000
func TestFindMedianSortedArraysExample6(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{2, 7}
	expected := 2.5

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}

// Example 7:
//
// Input: nums1 = [1, 2], nums2 = [-1, 3]
// Output: 2.50000
func TestFindMedianSortedArraysExample7(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{-1, 3}
	expected := 1.5

	if lc4.FindMedianSortedArrays(input1, input2) != expected {
		t.Errorf("got %f, expected %f", lc4.FindMedianSortedArrays(input1, input2), expected)
	}
}
