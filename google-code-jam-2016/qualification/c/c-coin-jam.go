package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
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

type Ret struct {
	Index  int
	Factor float64
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
		// Channel for the return values
		c := make(chan Ret, 9)
		// Channel to stop processing
		stop := make(chan bool, 1)
		for j := 2; j < 11; j++ {
			findFactor(j, c, stop, jamcoin)
		}
		for k := 0; k < 9; k++ {
			r := <-c
			if r.Factor < 1 {
				stop <- true
				continue Outer
			}
			factors[r.Index] = strconv.FormatFloat(r.Factor, 'f', 0, 64)
		}
		close(c)
		close(stop)
		ret = ret + "\n" + jamcoin + " " + strings.Trim(fmt.Sprint(factors), "[]")
		count++
	}
	return ret
}

func findFactor(base int, c chan Ret, stop chan bool, jamcoin string) {
	// Convert jamcoin to decimal using given base
	num, _ := strconv.ParseInt(jamcoin, base, 64)
	var factor float64
	x := big.NewInt(num)
	// Check if num is prime with a probability of 0.9
	if x.ProbablyPrime(10) {
		factor = 0
	} else {
		factor = firstfactor(float64(num), stop)
	}
	c <- Ret{base - 2, factor}
}

func firstfactor(num float64, stop chan bool) float64 {
	for i := 2; float64(i) < num/2; i++ {
		select {
		case <-stop:
			fmt.Println("Stopping goroutine")
			// If a signal is sent to the stop channel, just return 0
			return 0
		default:
			remainder := math.Remainder(num, float64(i))
			if remainder == 0 {
				return float64(i)
			}
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
