package mylib

import (
	"fmt"
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

// wholeStory : took like 10 mins
func wholeStory(in string) (out string) {
	sliceContainer := strings.Split(in, "-")

	var wordContainer string

	for i, st := range sliceContainer {

		if strings.ContainsAny("abcdefghijklmnopqrstuvwxyz", strings.ToLower(st)) {

			if i == 1 {
				wordContainer = st
			} else {
				wordContainer = fmt.Sprintf("%s %s", wordContainer, st)
			}

		}

	}

	return wordContainer

}

// storyStats : it took a about 30 mins
func storyStats(in string) (short, long string, avg float32, listWithAvg []string) {
	sliceContainer := strings.Split(in, "-")
	shortWordLength := 0
	longWordLength := 0
	wordsCount := 0
	wordsSum := 0

	for i, st := range sliceContainer {

		if strings.ContainsAny("abcdefghijklmnopqrstuvwxyz", strings.ToLower(st)) {
			if i == 0 {
				shortWordLength = len(st)
				longWordLength = len(st)
			}
			if len(st) < shortWordLength {
				shortWordLength = len(st)
				short = st

			}
			if len(st) > longWordLength {
				longWordLength = len(st)
				long = st
			}

			wordsCount += 1
			wordsSum = len(st)
		}
	}

	avg = float32(wordsSum / wordsCount)

	for _, st := range sliceContainer {
		if strings.ContainsAny("abcdefghijklmnopqrstuvwxyz", strings.ToLower(st)) {
			if len(st) == int(avg) {
				listWithAvg = append(listWithAvg, st)
			}
		}
	}

	return
}

// generate : this will likely take me longer than 30mins because of the spec string pattern
func generate(in bool) (out string) {
	if in {

	}
	return out
}
