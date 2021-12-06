package day4

import (
	"strings"
)

type BingoCard [5][5]int
type BingoState [5][5]bool
type Bingo struct {
	Layout BingoCard
	State  BingoState
}

func PartOneSolution(fileInput []string) int {

	balls := strings.Split(fileInput[0], ",")

	return 0
}
