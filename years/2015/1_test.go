package main

import (
	"testing"
)

type testCasePart1 struct {
	data          string
	expectedFloor int
}

type testCasePart2 struct {
	data         string
	basementChar int
}

func TestDay1Part1(t *testing.T) {
	var testCases = []testCasePart1{
		{
			data:          "(())",
			expectedFloor: 0,
		},
		{
			data:          "()()",
			expectedFloor: 0,
		},
		{
			data:          "(((",
			expectedFloor: 3,
		},
		{
			data:          "(()(()(",
			expectedFloor: 3,
		},
		{
			data:          "(((",
			expectedFloor: 3,
		},
		{
			data:          "))(((((",
			expectedFloor: 3,
		},
		{
			data:          "())",
			expectedFloor: -1,
		},
		{
			data:          "))(",
			expectedFloor: -1,
		},
		{
			data:          ")))",
			expectedFloor: -3,
		},
		{
			data:          ")())())",
			expectedFloor: -3,
		},
	}

	for _, testCase := range testCases {
		floor := solveDay1Part1(testCase.data)
		if floor != testCase.expectedFloor {
			t.Errorf("Expected %d, but got: %d", testCase.expectedFloor, floor)
		}
	}
}

func TestDay1Part2(t *testing.T) {
	var testCases = []testCasePart2{
		{
			data:         ")",
			basementChar: 1,
		},
		{
			data:         "()())",
			basementChar: 5,
		},
	}

	for _, testCase := range testCases {
		char := solveDay1Part2(testCase.data)
		if char != testCase.basementChar {
			t.Errorf("Expected %d, but got: %d", testCase.basementChar, char)
		}
	}
}
