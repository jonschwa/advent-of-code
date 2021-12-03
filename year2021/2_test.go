package year2021

import (
	"testing"
)

type testCaseDay2 struct {
	input    []string
	expected int
}

func TestDay2Part1(t *testing.T) {
	var testCases = []testCaseDay2{
		{[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, 150},
	}

	for _, testCase := range testCases {
		testCommands := convertLinesToCommands(testCase.input)
		result := solveDay2Part1(testCommands)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay2Part2(t *testing.T) {
	var testCases = []testCaseDay2{
		{[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, 900},
	}

	for _, testCase := range testCases {
		testCommands := convertLinesToCommands(testCase.input)
		result := solveDay2Part2(testCommands)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}

}
