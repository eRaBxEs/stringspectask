package mylib

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"
const digitset = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// testValidity : took a little bit over 20 mins
func testValidity(in string) (out bool, err error) {
	if len(in) == 0 {
		return out, nil
	}

	if !strings.ContainsAny(in, "-") {
		return out, nil
	}

	for i, r := range in {

		if i == 0 {
			if !strings.Contains("abcdefghijklmnopqrstuvwxyz0123456789", strings.ToLower(string(r))) {
				return out, nil
			}
		}

	}

	out = true

	return out, nil
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
		return fmt.Sprintf("%s-%s-%s-%s", StringWithDigitset(2, digitset), StringWithCharset(4, charset), StringWithDigitset(2, digitset), StringWithCharset(4, charset))
	}
	return fmt.Sprintf("%s-%s-%s-%s", StringWithCharset(4, charset), StringWithDigitset(2, digitset), StringWithDigitset(2, digitset), StringWithCharset(4, charset))
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringWithDigitset(length int, digitset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = digitset[seededRand.Intn(len(digitset))]
	}
	return string(b)
}
