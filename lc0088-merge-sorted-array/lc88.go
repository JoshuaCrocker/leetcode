package lc0088_merge_sorted_array

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
// representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To
// accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func merge(nums1 []int, m int, nums2 []int, n int) {
	im := m - 1
	in := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if in < 0 {
			break
		}

		if im >= 0 && nums1[im] > nums2[in] {
			nums1[i] = nums1[im]
			im--
		} else {
			nums1[i] = nums2[in]
			in--
		}
	}
}
