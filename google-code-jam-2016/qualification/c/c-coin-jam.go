package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const input string = "c-coin-jam.input"

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
	// Implement solution here
	parts := strings.Split(line, " ")
	length, _ := strconv.Atoi(parts[0])
	limit, _ := strconv.Atoi(parts[1])
	count := 0
	ret := ""

Outer:
	for i := 0; count < limit; i++ {
		jamcoin := jamcoin(int64(i), length)
		var factors [10]string
		for j := 2; j < 11; j++ {
			// Convert jamcoin to decimal using base j
			num, _ := strconv.ParseInt(jamcoin, j, 64)
			factor := firstfactor(float64(num))
			if factor == 0 {
				continue Outer
			}
			factors[j-2] = strconv.FormatFloat(factor, 'f', 0, 64)
		}
		ret = ret + "\n" + jamcoin + " " + strings.Trim(fmt.Sprint(factors), "[]")
		count++
	}
	return ret
}

func firstfactor(num float64) float64 {
	for i := 2; float64(i) < num/2; i++ {
		remainder := math.Remainder(num, float64(i))
		if remainder == 0 {
			return float64(i)
		}
	}
	return 0
}

// Returns the jamcoin of index i with a specified length
// e.g.
// i = 0, length = 4 -> 1001
// i = 1, length = 4 -> 1011
// i = 2, length = 4 -> 1101
func jamcoin(i int64, length int) string {
	a := strconv.FormatInt(i, 2)
	format := "1%0" + strconv.Itoa(length-2) + "s1"
	return fmt.Sprintf(format, a)
}
