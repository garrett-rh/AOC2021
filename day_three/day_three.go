package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func read_file(file_path *string) (file_lines []string, length int) {
	file, err := os.Open(*file_path)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	scanner := bufio.NewScanner(file) // Takes the pointer created by os.Open
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if length == 0 { // Go inits ints to 0
			length = len(scanner.Text())
		}
		file_lines = append(file_lines, scanner.Text())
	}

	file.Close()
	return
}

func part1_solution(txtlines []string, length int) (count int) {

	// Transpose the list to get the rows into columns
	var transposed_list [][]byte
	for i := 0; i < length; i++ {
		var list_item []byte // Placed here so that the values are emptied after every iteration
		for _, line := range txtlines {
			list_item = append(list_item, line[i])
		}
		transposed_list = append(transposed_list, list_item)
	}

	var binary_column string
	for _, arr := range transposed_list {
		var binary_counter int // Placed inside for loop to reset to 0 at every iteration
		for _, item := range arr {
			if item == 48 { // 48 becuase the ascii representation of 0
				binary_counter -= 1
			} else {
				binary_counter += 1
			}
		}

		if binary_counter > 0 {
			binary_column += "1"
		} else {
			binary_column += "0"
		}
	}

	epsilon, _ := strconv.ParseUint(binary_column, 2, len(binary_column))
	gamma := ^epsilon & uint64((math.Pow(2, float64(len(binary_column))) - 1)) // Calculates the bitmask

	count = int(epsilon) * int(gamma)
	return
}

func part2_solution(txtlines []string, length int) (count int) {
	// Transpose the list to get the rows into columns
	var transposed_list [][]byte
	for i := 0; i < length; i++ {
		var list_item []byte // Placed here so that the values are emptied after every iteration
		for _, line := range txtlines {
			list_item = append(list_item, line[i])
		}
		transposed_list = append(transposed_list, list_item)
	}

	fmt.Println(txtlines)
	//txtlines2 = txtlines
	for _, arr := range transposed_list {
		var list_of_ones []int
		var list_of_zeros []int
		for idx2, ascii := range arr {
			if ascii == 48 {
				list_of_zeros = append(list_of_zeros, idx2)
			} else {
				list_of_ones = append(list_of_ones, idx2)
			}
		}
		if len(list_of_zeros) < len(list_of_ones) {
			for _, item := range list_of_zeros {
				txtlines[item] = "null"
			}
		} else {
			for _, item := range list_of_ones {
				txtlines[item] = "null"
			}
		}

		var counter [2]int
		for idx, item := range txtlines {
			if item != "null" {
				counter[0] += 1
				counter[1] = idx
			}
		}
		if counter[0] == 1 {
			fmt.Println(counter[1])
			counter[0] = 0
			break
		}

		for idx, item := range txtlines {
			if item != "null" {
				counter[0] += 1
				if counter[0] == 1 {
					counter[1] = idx
					break
				}
			}
		}

	}
	fmt.Println(txtlines) //, string_not_in_slice("null", txtlines))
	return
}

// Helper function for part 2
//func string_not_in_slice(a string, list []string) int {
//	for idx, item := range list {
//		if a != item {
//			return idx
//		} else {
//			return
//		}
//	}
//}

func main() {
	file_ptr := flag.String("f", "input.txt", "File to read challenge input from.")
	flag.Parse()

	answer_one_count := part1_solution(read_file(file_ptr))
	answer_two_count := part2_solution(read_file(file_ptr))
	fmt.Printf("Answer 1: %d\n", answer_one_count)
	fmt.Printf("Answer 2: %d\n", answer_two_count)
}
