package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const input string = "b-revenge-of-pancakes.input"

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

func processLine(line string) string {
	flipped := 0
	lastChar := ""
	for i := 0; i < len(line); i++ {
		nextChar := string(line[i])
		if nextChar != lastChar {
			flipped++
			lastChar = nextChar
		}
	}
	// We shouldn't have done the last flip if the string was "+"
	if lastChar == "+" {
		flipped--
	}
	return strconv.Itoa(flipped)
}
