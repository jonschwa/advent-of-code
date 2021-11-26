package main

import (
	"fmt"
	"jonschwa/aoc/helpers"
)

func main() {
	data, err := helpers.GetDataForDay(2015, 1)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}
	fmt.Println("Part 1:", solveDay1Part1(data[0]))
	fmt.Println("Part 2:", solveDay1Part2(data[0]))

}

func solveDay1Part1(data string) int {
	var floor int = 0
	for _, char := range data {
		if char == []rune("(")[0] {
			floor++
		} else if char == []rune(")")[0] {
			floor--
		}
	}
	return floor
}

func solveDay1Part2(data string) int {
	var floor int = 0
	for i, char := range data {
		if char == []rune("(")[0] {
			floor++
		} else if char == []rune(")")[0] {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return floor
}
