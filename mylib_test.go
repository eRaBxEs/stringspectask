package mylib

import (
	"testing"
)

func TestTestValidity(t *testing.T) {
	myTest := []struct {
		inputName      string
		expectedOutput bool
		expectedError  error
	}{
		{inputName: "23-ab-48-caba-56-haha", expectedOutput: true, expectedError: nil},
	}

	for _, tt := range myTest {

		actualOutput, actualError := testValidity(tt.inputName)

		if actualError != tt.expectedError {
			t.Error("should not get an error")
		}
		if actualOutput != tt.expectedOutput {
			t.Errorf("got %v, want %v", actualOutput, tt.expectedOutput)
		}
	}
}

func TestAverageNumber(t *testing.T) {
	myTest := []struct {
		inputName      string
		expectedOutput int
		expectedError  error
	}{
		{inputName: "23-ab-48-caba-56-haha", expectedOutput: 5, expectedError: nil},
	}

	for _, tt := range myTest {

		actualOutput, actualError := averageNumber(tt.inputName)

		if actualError != tt.expectedError {
			t.Error("should not get an error")
		}
		if actualOutput != tt.expectedOutput {
			t.Errorf("got %v, want %v", actualOutput, tt.expectedOutput)
		}
	}
}

func TestWholeStory(t *testing.T) {
	myTest := []struct {
		inputName      string
		expectedOutput string
		expectedError  error
	}{
		{inputName: "23-ab-48-caba-56-haha", expectedOutput: "ab caba haha", expectedError: nil},
	}

	for _, tt := range myTest {

		actualOutput, actualError := wholeStory(tt.inputName)

		if actualError != tt.expectedError {
			t.Error("should not get an error")
		}
		if actualOutput != tt.expectedOutput {
			t.Errorf("got %v, want %v", actualOutput, tt.expectedOutput)
		}
	}
}

func TestStoryStats(t *testing.T) {
	myTest := []struct {
		inputName                 string
		expectedOutputShort       string
		expectedOutputLong        string
		expectedOutputAvg         int
		expectedOutputListWithAvg []string
		expectedError             error
	}{
		{inputName: "23-ab-48-caba-56-haha", expectedOutputShort: "ab", expectedOutputLong: "caba",
			expectedOutputAvg: 3, expectedOutputListWithAvg: []string{}, expectedError: nil},
	}

	for _, tt := range myTest {

		actualOutputShort, actualOutputLong, actualOutputAvg, actualOutputListWithAvg, actualError := storyStats(tt.inputName)

		if actualError != tt.expectedError {
			t.Error("should not get an error")
		}
		if actualOutputShort != tt.expectedOutputShort {
			t.Errorf("got %v, want %v", actualOutputShort, tt.expectedOutputShort)
		}
		if actualOutputLong != tt.expectedOutputLong {
			t.Errorf("got %v, want %v", actualOutputLong, tt.expectedOutputLong)
		}
		if actualOutputAvg != tt.expectedOutputAvg {
			t.Errorf("got %v, want %v", actualOutputAvg, tt.expectedOutputAvg)
		}

		if len(actualOutputListWithAvg) != len(tt.expectedOutputListWithAvg) {
			t.Errorf("got length of %v, want length of %v", actualOutputListWithAvg, tt.expectedOutputListWithAvg)
		}
	}
}
