package year2021

import (
	"testing"
)

type testCaseDay3 struct {
	input    []string
	expected int64
}

func TestDay3Part1(t *testing.T) {
	var testCases = []testCaseDay3{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}, 198},
	}

	for _, testCase := range testCases {
		result := solveDay3Part1(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay3Part2(t *testing.T) {
	var testCases = []testCaseDay3{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}, 230},
	}

	for _, testCase := range testCases {
		result := solveDay3Part2(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}
