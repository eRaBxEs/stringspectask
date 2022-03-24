package mylib

import (
	"strconv"
	"strings"
)

// testValidity : took a little bit over 20 mins
func testValidity(in string) (out bool) {
	if len(in) == 0 {
		return out
	}

	if !strings.ContainsAny(in, "-") {
		return out
	}

	for _, r := range in {

		if !strings.ContainsAny("abcdefghijklmnopqrstuvwxyz0123456789", strings.ToLower(string(r))) {
			return out
		}

	}

	out = true

	return out
}

// averageNumber : took about 15 mins
func averageNumber(in string) (out float32) {
	sliceContainer := strings.Split(in, "-")
	numbersCount := 0
	numbersSum := 0

	for _, st := range sliceContainer {

		if strings.ContainsAny("0123456789", st) {
			// count the numbers
			numbersCount += len(st)

			// loop through the substrings that contains numbers
			for _, num := range st {
				intVar, err := strconv.Atoi(string(num))
				if err != nil {
					return float32(-1)
				}
				numbersSum += intVar
			}
		}

	}

	return float32(numbersSum / numbersCount)

}
