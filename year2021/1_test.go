package year2021

import (
	"testing"
)

type testCase struct {
	input    []int
	expected int
}

func TestDay1Part1(t *testing.T) {

	var testCases = []testCase{
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
	}

	for _, testCase := range testCases {
		result := solveDay1Part1(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay2Part2(t *testing.T) {
	var testCases = []testCase{
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 5},
	}

	for _, testCase := range testCases {
		result := solveDay1Part2(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}
