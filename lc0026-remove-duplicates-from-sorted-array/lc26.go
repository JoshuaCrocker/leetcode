package lc0026_remove_duplicates_from_sorted_array

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique
// element appears only once. The relative order of the elements should be kept the same. Then return the number of
// unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
//
//    - Change the array nums such that the first k elements of nums contain the unique elements in the order they were
//      present in nums initially. The remaining elements of nums are not important as well as the size of nums.
//    - Return k.

// f  b
// 0, 0, 1, 1, 1, 2, 2, 3, 3, 4

func removeDuplicates(nums []int) int {
	insert, read := 1, 1

	for read < len(nums) {
		if nums[read-1] != nums[read] {
			nums[insert] = nums[read]
			insert++
		}
		read++
	}

	return insert
}
