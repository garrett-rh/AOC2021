package day3

import (
	"aoc2021/utils"
	"math"
	"strconv"
)

// PartOneSolution is a starter function for solving part one.
// Takes in the file as a list of strings.
// Each item in the list is the puzzle input, split at the newline.
// Then, a 'count' is returned which represents the puzzle answer.
func PartOneSolution(txtLines []string) (count int) {

	transposedList := utils.ListTransposer(txtLines)

	var binaryColumn string
	for _, arr := range transposedList {
		var binaryCounter int
		for _, item := range arr {
			// 48 becuase the ascii representation of 0
			// Counts the occurrences of 1 or 0
			if item == 48 {
				binaryCounter -= 1
			} else {
				binaryCounter += 1
			}
		}

		if binaryCounter > 0 {
			binaryColumn += "1"
		} else {
			binaryColumn += "0"
		}
	}

	epsilon, _ := strconv.ParseUint(binaryColumn, 2, len(binaryColumn))
	gamma := ^epsilon & uint64((math.Pow(2, float64(len(binaryColumn))) - 1))

	count = int(epsilon) * int(gamma)
	return
}
