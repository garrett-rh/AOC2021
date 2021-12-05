package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadFile is a utility to read in the challenge input.
// It accepts a file pointer, as given by the command line argument '-f'.
// Then, a list of strings is created from the input file using the new line as a delimeter.
func ReadFile(file_path *string) (file_lines []string) {
	file, err := os.Open(*file_path)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	// Takes the pointer created by os.Open
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		file_lines = append(file_lines, scanner.Text())
	}
	file.Close()

	return
}
