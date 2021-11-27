package year2015

import (
	"testing"
)

type testCase struct {
	input    []string
	expected int
}

func TestDay2Part1(t *testing.T) {

	var testCases = []testCase{
		{[]string{"2x3x4"}, 58},
		{[]string{"1x1x10"}, 43},
	}

	for _, testCase := range testCases {
		result := solveDay2Part1(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}

func TestDay2Part2(t *testing.T) {

}
