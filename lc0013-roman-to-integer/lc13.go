// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
// Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same
//  principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
//     I can be placed before V (5) and X (10) to make 4 and 9.
//     X can be placed before L (50) and C (100) to make 40 and 90.
//     C can be placed before D (500) and M (1000) to make 400 and 900.
//
// Given a roman numeral, convert it to an integer.
package lc13romantointeger

import (
	"log"
	"strings"
)

func RomanToInt(s string) int {
	log.Println("---")

	values := make(map[string]int)
	values["I"] = 1
	values["V"] = 5
	values["X"] = 10
	values["L"] = 50
	values["C"] = 100
	values["D"] = 500
	values["M"] = 1000

	numerals := strings.Split(s, "")

	total := 0

	i := 0
	for i < len(numerals) {
		numeral := numerals[i]
		value := values[numeral]

		nextNum := ""
		nextV := 0

		if i+1 < len(numerals) {
			nextNum = numerals[i+1]
			nextV = values[nextNum]
		}

		if nextV > value {
			log.Printf("Parse %s%s = %d", numeral, nextNum, (nextV - value))
			total += nextV - value
			i += 2
		} else {
			log.Printf("Parse %s = %d", numeral, value)
			total += value
			i++
		}

	}

	return total
}
