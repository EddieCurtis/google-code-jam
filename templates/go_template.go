package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const input string = "input_file.txt"

func main() {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		lines := strings.Split(string(data), "\n")
		// Remove first value as that just defines number of input lines
		for i, next := range lines[1:] {
			if next != "" {
				result := processLine(next)
				fmt.Printf("Case #%d: %s\n", i+1, result)
			}
		}
	}
}

func processLine(line string) {
	// Implement solution here
	return ""
}
