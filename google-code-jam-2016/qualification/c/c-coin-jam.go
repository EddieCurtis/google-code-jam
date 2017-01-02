package main

import (
	"fmt"
	"io/ioutil"
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
	var factor float64
	// Use big rather than int64 because the numbers will be too large for the 32 digit exercise
	x := new(big.Int)
	x.SetString(jamcoin, base)
	// Check if num is prime with a probability of 0.9
	if x.ProbablyPrime(10) {
		factor = 0
	} else {
		factor = firstfactor(x, stop)
	}
	c <- Ret{base - 2, factor}
}

func firstfactor(num *big.Int, stop chan bool) float64 {
	for i := 2; i < 10000; i++ {
		select {
		case <-stop:
			// If a signal is sent to the stop channel, just return 0
			return 0
		default:
			z := new(big.Int).Set(num)
			if z.Rem(z, big.NewInt(int64(i))).Int64() == 0 {
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
