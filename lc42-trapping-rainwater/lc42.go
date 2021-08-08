package lc42trappingrainwater

import (
	"log"
	"reflect"
)

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

// Example 2:

// Constraints:

//     n == height.length
//     0 <= n <= 2 * 104
//     0 <= height[i] <= 105

func Trap(height []int) int {
	// If we have less than three columns, there's no way water could be held.
	mapLength := len(height)
	// log.Printf("mapLength=%v", mapLength)
	if mapLength < 3 {
		return 0
	}

	// Build world array
	maxHeight := maxInArray(height)

	world := make([][]string, len(height))

	for i, h := range height {
		world[i] = make([]string, maxHeight)

		for h > 0 {
			world[i][h-1] = "x"
			h--
		}
	}

	log.Printf("height=%v", height)
	log.Printf("world=%v", world)

	// Simulate rain until the world stops chainging
	// 1. Fill all empty squares with water
	for columnIndex, column := range world {
		for rowIndex, row := range column {
			if row != "x" {
				world[columnIndex][rowIndex] = "w"
			}
		}
	}

	log.Printf("world=%v", world)

	// 2. Start simulation
	// 2.1. Make a copy of the world array
	equalityCounter := 0
	threshold := 5

	// 2.2. Flow water
	// We want to remove water if it's on the edge of the world, or if there isn't anything next to it.
	for equalityCounter < threshold {
		oldWorld := world

		// TODO flip this to work row-by-row

		for columnIndex, column := range world {
			for rowIndex, _ := range column {
				// Check if we're working with water
				if world[columnIndex][rowIndex] == "w" {
					// Is first column
					if columnIndex == 0 {
						world[columnIndex][rowIndex] = ""
					}

					// Is last column
					if columnIndex == len(height)-1 {
						world[columnIndex][rowIndex] = ""
					}

					// Empty cell left
					if columnIndex >= 1 && world[columnIndex-1][rowIndex] == "" {
						world[columnIndex][rowIndex] = ""
					}

					// Empty cell right
					if columnIndex < len(height)-1 && world[columnIndex+1][rowIndex] == "" {
						world[columnIndex][rowIndex] = ""
					}
				}
			}
		}

		log.Printf("world=%v", world)
		log.Printf("equal=%v", reflect.DeepEqual(world, oldWorld))

		if reflect.DeepEqual(world, oldWorld) {
			equalityCounter++
		} else {
			equalityCounter = threshold
		}
	}

	// 3. Count up water tiles
	total := 0

	for _, column := range world {
		for _, row := range column {
			if row == "w" {
				total++
			}
		}
	}

	return total
}

func maxInArray(arr []int) int {
	max := 0

	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}
