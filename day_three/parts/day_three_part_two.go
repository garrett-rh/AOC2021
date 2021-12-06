package day3

import (
	"strconv"
)

// Stole this from https://github.com/krmaxwell because I suck.
// I did end up writing it all in instead of copying & pasting if that nets me any bonus points.
func PartTwoSolution(txtLines []string) (count int64) {

	count = oxygenGeneratorRating(txtLines) * carbonDioxideGeneratorRating(txtLines)
	return
}

func oxygenGeneratorRating(txtLines []string) (count int64) {

	for i := 0; i < len(txtLines[0]); i++ { // Iterates through using the length to later reference i as a column index
		mostCommonBit := commonBit(i, txtLines, 1) // Gets most common bit for oxygen rating with a flag of 1 for oxygen
		var newLines []string
		for _, line := range txtLines {
			if line[i] == byte(mostCommonBit) { // checks if the binary for the line[column] is equal to the most common bit for the column
				newLines = append(newLines, line) // Creates the new list to store the line and keep it for the next iteration
			}
		}
		txtLines = newLines     // this replaces txtLines with the passing lines so that the next loop will have an updated list
		if len(txtLines) == 1 { // Prevents index out of range on line 17
			break
		}
	}
	count, _ = strconv.ParseInt(txtLines[0], 2, 0) // converts binary to decimal
	return
}

func carbonDioxideGeneratorRating(txtLines []string) (count int64) {

	for i := 0; i < len(txtLines[0]); i++ {
		leastCommonBit := commonBit(i, txtLines, 0) // this line only cares for the rows that fall into least common
		var newLines []string
		for _, line := range txtLines {
			if line[i] == byte(leastCommonBit) { // if the line item is equal to the least common bit, add it to the list
				newLines = append(newLines, line)
			}
		}
		txtLines = newLines
		if len(txtLines) == 1 { // Prevents index out of range on line 35
			break
		}
	}
	count, _ = strconv.ParseInt(txtLines[0], 2, 0)
	return
}

func commonBit(pos int, txtLines []string, flag byte) (bit rune) {
	var zero int
	var one int

	for _, binary := range txtLines {
		if binary[pos] == '0' { // iterates through the lines of input using the passed in 'pos' var as a column reference and searches for 0
			zero++
		} else {
			one++
		}
	}
	if flag == 1 && one >= zero { // This is for the oxygen generator & for the oxygen generator's bit of interest to be 1
		return '1'
	} else if flag == 0 && zero > one { // This is for the co2 generator & for the bit of interest to be 1
		return '1'
	}
	return '0' // All other cases result in the bit of interest to be 0
}
