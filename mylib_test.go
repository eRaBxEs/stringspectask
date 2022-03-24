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
