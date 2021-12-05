package main

import (
	parts "aoc2021/day_three/parts"
	"aoc2021/utils"
	"flag"
	"fmt"
)

// Main handles command line inputs and starting the process of completing the challenges.
func main() {
	filePtr := flag.String("f", "input.txt", "File to read challenge input from.")
	flag.Parse()

	fileInput := utils.ReadFile(filePtr)

	answerOneCount := parts.PartOneSolution(fileInput)
	answerTwoCount := parts.PartTwoSolution(fileInput)
	fmt.Printf("Answer 1: %d\n", answerOneCount)
	fmt.Printf("Answer 2: %d\n", answerTwoCount)
}
