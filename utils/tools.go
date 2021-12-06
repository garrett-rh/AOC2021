package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
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

// listTransposer takes a list and flips it on its side.
// A list of [abc abc abc] becomes [aaa bbb ccc].
func ListTransposer(txtLines []string) (transposedList []string) {

	length := len(txtLines[0])
	for i := 0; i < length; i++ {
		// listItem will hold the items in the row -> column transtion
		// using the above example, listItem will hold [aaa], then [bbb], ...
		var listItem []string
		for _, line := range txtLines {
			listItem = append(listItem, string(line[i]))
		}
		//Then, the items are joined and returned back in the form they were presented in
		transposedList = append(transposedList, strings.Join(listItem, ""))
	}
	return
}
