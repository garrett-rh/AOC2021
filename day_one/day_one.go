package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func read_file(file_path *string) (file_lines []int64) {
	file, err := os.Open(*file_path)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(file) // Takes the pointer created by os.Open
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		int1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		file_lines = append(file_lines, int1)
	}
	file.Close()

	return
}

func part1_solution(txtlines []int64) (part_one_count int) {
	part_one_count = 0
	for idx, line := range txtlines {
		if idx == 0 { // Prevents out of range indexing
			continue
		} else if line > txtlines[idx-1] {
			part_one_count += 1
		}
	}
	return
}

func part2_solution(txtlines []int64) (part_two_count int) {
	var last_three_line_sum int64
	var new_three_line_sum int64
	for idx, line := range txtlines {
		if idx == 0 { // Prevents out of range indexing
			continue
		} else if idx > (len(txtlines) - 3) { // Breaks if there aren't enough measurements left
			break
		} else {
			new_three_line_sum = line + txtlines[idx+1] + txtlines[idx+2]
			last_three_line_sum = txtlines[idx-1] + line + txtlines[idx+1]
		}
		if new_three_line_sum > last_three_line_sum {
			part_two_count += 1
		}
		last_three_line_sum, new_three_line_sum = 0, 0 // Reset line sums to not interfere w/ next comparison
	}
	return
}

func main() {
	file_ptr := flag.String("f", "input.txt", "File to read challenge input from.")
	flag.Parse()

	answer_one_count := part1_solution(read_file(file_ptr))
	answer_two_count := part2_solution(read_file(file_ptr))
	fmt.Printf("Answer 1: %d\n", answer_one_count)
	fmt.Printf("Answer 2: %d\n", answer_two_count)
}
