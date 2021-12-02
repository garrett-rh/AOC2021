package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func read_file(file_path *string) (file_lines []string) {
	file, err := os.Open(*file_path)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	scanner := bufio.NewScanner(file) // Takes the pointer created by os.Open
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		file_lines = append(file_lines, scanner.Text())
	}
	file.Close()
	return
}

func part1_solution(txtlines []string) (count int) {
	var direction []byte
	var amount []int
	var horizontal int
	var depth int
	for _, line := range txtlines {
		int1, _ := strconv.Atoi(string(line[len(line)-1])) // Takes the byte given by the last char of line & converts to string to int
		amount = append(amount, int1)
		direction = append(direction, line[0])
	}

	for idx, cmd := range direction {
		switch cmd {
		case 'f':
			horizontal += amount[idx]
		case 'd':
			depth += amount[idx]
		case 'u':
			depth -= amount[idx]
		}
	}

	count = depth * horizontal

	return
}

func part2_solution(txtlines []string) (count int) {
	var direction []byte
	var amount []int
	var horizontal int
	var depth int
	var aim int

	for _, line := range txtlines {
		int1, _ := strconv.Atoi(string(line[len(line)-1])) // Takes the byte given by the last char of line & converts to string to int
		amount = append(amount, int1)
		direction = append(direction, line[0])
	}

	for idx, cmd := range direction {
		switch cmd {
		case 'f':
			horizontal += amount[idx]
			depth += aim * amount[idx]
		case 'd':
			aim += amount[idx]
		case 'u':
			aim -= amount[idx]
		}
	}

	count = depth * horizontal

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
