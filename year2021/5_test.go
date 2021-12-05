package year2021

import (
	"testing"
)

type testCaseDay5 struct {
	input    []string
	expected int
}

func TestDay5Part1(t *testing.T) {
	var testCases = []testCaseDay5{
		{[]string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}, 5},
	}

	for _, testCase := range testCases {
		result := solveDay5Part1(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay5Part2(t *testing.T) {
	var testCases = []testCaseDay5{
		{[]string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}, 12},
		// {[]string{
		// 	"1,1 -> 3,3",
		// 	"9,7 -> 7,9",
		// }, 12},
	}

	for _, testCase := range testCases {
		result := solveDay5Part2(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}
