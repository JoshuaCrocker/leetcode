package lc4medianoftwosortedarrays

import (
	"log"
	"math"
)

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	log.Println("")
	log.Println("--- START ---")

	if len(nums1) == 1 && len(nums2) == 0 {
		return float64(nums1[0])
	}

	if len(nums2) == 1 && len(nums1) == 0 {
		return float64(nums2[0])
	}

	combined := CombineArrays(nums1, nums2)
	log.Printf("combined %v", combined)
	return FindMedian(combined)
}

func CombineArrays(nums1 []int, nums2 []int) []int {
	output := []int{}
	output = append(output, nums1...)
	output = append(output, nums2...)

	changed := true
	for changed {
		index := 0
		changed = false

		for index < len(output)-1 {
			log.Printf("index=%d", index)
			current := output[index]
			next := output[index+1]
			log.Printf("next=%d, current=%d", next, current)
			if next < current {
				log.Println("--- swapping")
				output[index] = next
				output[index+1] = current
				changed = true
			}
			index++
		}
	}

	return output
}

func FindMedian(nums []int) float64 {
	length := len(nums)
	if length > 0 {
		if length%2 == 0 {
			min := (length / 2) - 1
			max := (length / 2)

			return (float64(nums[min]) + float64(nums[max])) / 2
		} else {
			index := int(math.Ceil(float64(length)/2.0)) - 1
			return float64(nums[index])
		}
	}

	return 0.0
}
