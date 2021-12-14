package year2021

import (
	"testing"
)

type testCaseDay9 struct {
	input    []string
	expected int
}

func TestDay9Part1(t *testing.T) {
	var testCases = []testCaseDay9{
		{[]string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		}, 15},
	}

	for _, testCase := range testCases {
		result := solveDay9Part1(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay9Part2(t *testing.T) {
	var testCases = []testCaseDay9{
		{[]string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		}, 1134},
	}

	for _, testCase := range testCases {
		result := solveDay9Part2(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}
