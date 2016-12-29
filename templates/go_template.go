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
		for _, next := range lines[1:] {
			if next != "" {
				processLine(next)
			}
		}
	}
}

func processLine(line string) {
	fmt.Println(line)
}
