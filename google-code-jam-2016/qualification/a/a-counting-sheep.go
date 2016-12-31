package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const input string = "A-large-practice.in"

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
	n, e := strconv.Atoi(line)
	if e == nil {
		m := make(map[string]bool)
		if n != 0 {
			for i := 1; len(m) < 10 && i < 10000001; i++ {
				a := strconv.Itoa(n * i)
				for _, x := range a {
					m[string(x)] = true
					// Once we get all 10 characters return the current number
					if len(m) == 10 {
						return a
					}
				}
			}
		}
		return "INSOMNIA"
	} else {
		return e.Error()
	}
}
