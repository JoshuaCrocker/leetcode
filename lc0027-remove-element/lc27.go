package lc0027_remove_element

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the
// elements may be changed. Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the
// following things:
//
// - Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
//	 The remaining elements of nums are not important as well as the size of nums.
// - Return k.

func removeElement(nums []int, val int) int {
	front, back := 0, len(nums)-1

	k := 0
	for front <= back {
		if nums[front] == val {
			nums[front], nums[back] = nums[back], nums[front]
			back--
		} else {
			front++
			k++
		}
	}

	return k
}
