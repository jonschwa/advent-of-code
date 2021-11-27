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

// A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
// A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
func TestDay2Part2(t *testing.T) {
	var testCases = []testCase{
		{[]string{"2x3x4"}, 34},
		{[]string{"1x1x10"}, 14},
	}

	for _, testCase := range testCases {
		result := solveDay2Part2(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %d, but got: %d", testCase.expected, result)
		}
	}
}
