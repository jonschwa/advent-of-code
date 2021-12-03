package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strconv"
)

type Container [][2]int

func RunDay3() {
	data, err := helpers.GetDataForDay(2021, 3)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	fmt.Println("Part 1:", solveDay3Part1(data))
	fmt.Println("Part 2:", solveDay3Part2(data))
}

func getRatesFromContainer(c Container) (int64, int64) {
	var epsilonRate, gammaRate string
	for _, counts := range c {
		if counts[0] > counts[1] {
			epsilonRate += "0"
			gammaRate += "1"
		} else {
			epsilonRate += "1"
			gammaRate += "0"
		}
	}

	epsilonInt, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		panic(err)
	}
	gammaInt, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		panic(err)
	}

	return epsilonInt, gammaInt
}

func solveDay3Part1(data []string) int64 {
	container := make(Container, len(data[0]))
	for _, chars := range data {
		for j, char := range chars {
			intChar, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			container[j][intChar]++
		}
	}

	epsilonRate, gammaRate := getRatesFromContainer(container)
	return epsilonRate * gammaRate
}

func solveDay3Part2(data []string) int64 {
	return 0
}
