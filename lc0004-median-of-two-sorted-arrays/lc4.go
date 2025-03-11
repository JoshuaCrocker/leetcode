package lc4medianoftwosortedarrays

import (
	"log"
	"math"
)

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	input := []int{}
	input = append(input, nums1...)
	input = append(input, nums2...)

	log.Println("")
	log.Printf("input=%v", input)

	sorted := MergeSort(input)
	log.Printf("sorted=%v", sorted)

	length := len(sorted)

	if length%2 == 0 {
		min := (length / 2) - 1
		max := length / 2
		log.Printf("min=%v", min)
		log.Printf("max=%v", max)

		return (float64(sorted[min]) + float64(sorted[max])) / 2
	}

	return float64(sorted[int(math.Floor(float64(length)/2))])
}

// MergeSort algorithm
//
// (1) If it is only one element in the list it is already sorted, return.
// (2) Divide the list recusively into two halves until it cannot be divided anymore.
// (3) Merge the smaller lists into new list in sorted order.
//
// Source: https://www.tutorialspoint.com/data_structures_algorithms/merge_sort_algorithm.htm
func MergeSort(input []int) []int {
	length := len(input)
	if length == 1 {
		return input
	}

	midpoint := length / 2

	left := input[:midpoint]
	right := input[midpoint:length]

	left = MergeSort(left)
	right = MergeSort(right)

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	output := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			output = append(output, right[0])
			right = right[1:]
		} else {
			output = append(output, left[0])
			left = left[1:]
		}
	}

	for len(left) > 0 {
		output = append(output, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		output = append(output, right[0])
		right = right[1:]
	}

	return output
}
