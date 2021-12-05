package day3

import (
	"math"
	"strconv"
)

// PartOneSolution is a starter function for solving part one.
// Takes in the file as a list of strings.
// Each item in the list is the puzzle input, split at the newline.
// Then, a 'count' is returned which represents the puzzle answer.
func PartOneSolution(txtLines []string) (count int) {

	transposedList := listTransposer(txtLines)

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

// listTransposer takes a list and flips it on its side.
// A list of [abc abc abc] becomes [[aaa] [bbb] [ccc]].
func listTransposer(txtLines []string) (transposedList [][]byte) {

	length := len(txtLines[0])
	for i := 0; i < length; i++ {
		// listItem will hold the items in the row -> column transtion
		// using the above example, listItem will hold [aaa], then [bbb], ...
		var listItem []byte
		for _, line := range txtLines {
			listItem = append(listItem, line[i])
		}
		transposedList = append(transposedList, listItem)
	}
	return
}
