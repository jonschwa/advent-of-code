package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
)

func RunDay1() {
	data := helpers.GetDataForDay(2021, 1)
	intData := helpers.ConvertStringSliceToIntSlice(data)
	fmt.Println("Part 1:", solveDay1Part1(intData))
	fmt.Println("Part 2:", solveDay1Part2(intData))
}

func solveDay1Part1(data []int) int {
	depthIncreases := 0
	for i, depth := range data {
		if i == 0 {
			continue
		}
		if depth > data[i-1] {
			depthIncreases++
		}
	}
	return depthIncreases
}

func solveDay1Part2(data []int) int {
	totalIncreases := 0
	sum1 := 0
	sum2 := 0
	for i, depth := range data {
		if i < 3 {
			continue
		}
		sum1 = data[i-1] + data[i-2] + data[i-3]
		sum2 = depth + data[i-1] + data[i-2]

		if sum2 > sum1 {
			totalIncreases++
		}

	}
	return totalIncreases
}
