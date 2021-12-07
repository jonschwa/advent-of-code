package year2021

import (
	"jonschwa/aoc/helpers"
	"strings"
	"testing"
)

type testCaseDay6 struct {
	input    []string
	numDays  int
	expected int
}

func TestDay6Part1(t *testing.T) {
	var testCases = []testCaseDay6{
		{[]string{
			"3,4,3,1,2",
		}, 18, 26},
		{[]string{
			"3,4,3,1,2",
		}, 80, 5934},
	}

	for _, testCase := range testCases {
		input := helpers.ConvertStringSliceToIntSlice(strings.Split(testCase.input[0], ","))
		result := solveDay6Part2(input, testCase.numDays)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay6Part2(t *testing.T) {
	var testCases = []testCaseDay6{
		{[]string{
			"3,4,3,1,2",
		}, 256, 26984457539},
	}
	for _, testCase := range testCases {
		input := helpers.ConvertStringSliceToIntSlice(strings.Split(testCase.input[0], ","))
		result := solveDay6Part2(input, testCase.numDays)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}
