package year2021

import (
	"jonschwa/aoc/helpers"
	"strings"
	"testing"
)

type testCaseDay7 struct {
	input    []string
	expected float64
}

func TestDay7Part1(t *testing.T) {
	var testCases = []testCaseDay7{
		{[]string{
			"16,1,2,0,4,2,7,1,2,14",
		}, 37},
	}

	for _, testCase := range testCases {
		input := helpers.ConvertStringSliceToIntSlice(strings.Split(testCase.input[0], ","))
		result := solveDay7Part1(input)
		if result != testCase.expected {
			t.Errorf("Expected %f, but got: %f", testCase.expected, result)
		}
	}
}

func TestDay7Part2(t *testing.T) {
	var testCases = []testCaseDay7{
		{[]string{
			"16,1,2,0,4,2,7,1,2,14",
		}, 168},
	}
	for _, testCase := range testCases {
		input := helpers.ConvertStringSliceToIntSlice(strings.Split(testCase.input[0], ","))
		result := solveDay7Part2(input)
		if result != testCase.expected {
			t.Errorf("Expected %f, but got: %f", testCase.expected, result)
		}
	}
}
