package lc1twosum_test

import (
	"reflect"
	"testing"

	lc1 "github.com/JoshuaCrocker/leetcode/lc0001-two-sum"
)

// TestTwoSumExample1
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
func TestTwoSumExample1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}

	result := lc1.TwoSum(nums, target)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestTwoSumExample2
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
func TestTwoSumExample2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}

	result := lc1.TwoSum(nums, target)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestTwoSumExample3
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
func TestTwoSumExample3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	expected := []int{0, 1}

	result := lc1.TwoSum(nums, target)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
