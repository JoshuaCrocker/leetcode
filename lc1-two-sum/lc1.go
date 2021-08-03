package lc1twosum

// TwoSum finds the indexes of the two integers within the given 'nums' array, which add up to the 'target' value.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// Constraints:
//     2 <= nums.length <= 104
//     -109 <= nums[i] <= 109
//     -109 <= target <= 109
//     Only one valid answer exists.
func TwoSum(nums []int, target int) []int {
	for xi, x := range nums {
		for yi, y := range nums {
			if xi != yi {
				if (x + y) == target {
					return []int{xi, yi}
				}
			}
		}
	}

	return []int{}
}
