package mylib

import (
	"fmt"
	"math"
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
func averageNumber(in string) (out int, err error) {
	sliceContainer := strings.Split(in, "-")
	numbersCount := 0.0
	numbersSum := 0.0

	for _, st := range sliceContainer {

		if strings.ContainsAny("0123456789", st) {
			// count the numbers
			numbersCount += float64(len(st))

			// loop through the substrings that contains numbers
			for _, num := range st {
				intVar, err := strconv.Atoi(string(num))
				if err != nil {
					return -1, err
				}
				numbersSum += float64(intVar)
			}
		}

	}

	avg := math.Round(numbersSum / numbersCount)
	return int(avg), nil

}

// wholeStory : took like 10 mins
func wholeStory(in string) (out string, err error) {
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

	return wordContainer, nil

}

// storyStats : it took a about 30 mins
func storyStats(in string) (short, long string, avg int, listWithAvg []string, err error) {
	sliceContainer := strings.Split(in, "-")
	shortWordLength := 100
	longWordLength := 0
	wordsCount := 0.0
	wordsSum := 0.0
	var onlyWords []string

	for _, st := range sliceContainer {

		if strings.ContainsAny("abcdefghijklmnopqrstuvwxyz", strings.ToLower(st)) {

			onlyWords = append(onlyWords, strings.ToLower(st))
		}

		wordsCount += 1
		wordsSum += float64(len(st))

	}

	shortWordLength, longWordLength = len(onlyWords[0]), len(onlyWords[0])
	short, long = onlyWords[0], onlyWords[0]

	for _, st := range onlyWords {

		if len(st) > longWordLength {
			longWordLength = len(st)
			long = st
		}

		if len(st) < shortWordLength {
			shortWordLength = len(st)
			short = st

		}

	}

	avg = int(math.Round(wordsSum / wordsCount))

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
