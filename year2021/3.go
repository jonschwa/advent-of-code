package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strconv"
)

type BitCriteria [2]int
type Container []BitCriteria

func RunDay3() {
	data, err := helpers.GetDataForDay(2021, 3)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	fmt.Println("Part 1:", solveDay3Part1(data))
	fmt.Println("Part 2:", solveDay3Part2(data))
}

func getPowerRatesFromContainer(c Container) (int64, int64) {
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

func makeContainer(data []string) Container {
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
	return container
}

func solveDay3Part1(data []string) int64 {
	container := makeContainer(data)
	epsilonRate, gammaRate := getPowerRatesFromContainer(container)
	return epsilonRate * gammaRate
}

func getBitCriteria(idx int, data []string, ratingType string) int {
	var bitCriteria int
	var bitCriteriaContainer BitCriteria
	for _, chars := range data {
		intChar, err := strconv.Atoi(string(chars[idx]))
		if err != nil {
			panic(err)
		}
		if intChar == 0 {
			bitCriteriaContainer[0]++
		} else {
			bitCriteriaContainer[1]++
		}
	}
	if ratingType == "o2" {
		// o2 rating uses the most common (or 1 as the tie breaker)
		if bitCriteriaContainer[0] > bitCriteriaContainer[1] {
			bitCriteria = 0
		} else if bitCriteriaContainer[0] < bitCriteriaContainer[1] {
			bitCriteria = 1
		} else {
			bitCriteria = 1
		}
	} else if ratingType == "co2" {
		// co2 rating uses the least common (or 0 as the tie breaker)
		if bitCriteriaContainer[0] > bitCriteriaContainer[1] {
			bitCriteria = 1
		} else if bitCriteriaContainer[0] < bitCriteriaContainer[1] {
			bitCriteria = 0
		} else {
			bitCriteria = 0
		}
	}

	return bitCriteria
}

func getUpdatedData(idx int, bitCriteria int, data []string) []string {
	newData := data[:0]
	for _, chars := range data {
		intChar, err := strconv.Atoi(string(chars[idx]))
		if err != nil {
			panic(err)
		}
		if intChar == bitCriteria {
			newData = append(newData, chars)
		}
	}
	return newData
}

func filterBitsByCriteria(idx int, data []string, ratingType string) string {
	i := idx
	bitCriteria := getBitCriteria(i, data, ratingType)
	newData := getUpdatedData(idx, bitCriteria, data)

	if len(newData) > 1 {
		i++
		return filterBitsByCriteria(i, newData, ratingType)
	}
	return newData[0]
}

func solveDay3Part2(data []string) int64 {
	o2Data := make([]string, len(data))
	co2Data := make([]string, len(data))

	// making copies because data is a reference type, and gets mutated in the recursive functions above,
	// which is v annoying. TODO - there must be a more idiomatic way to do this
	copy(o2Data, data)
	copy(co2Data, data)

	o2Rating := filterBitsByCriteria(0, o2Data, "o2")
	co2Rating := filterBitsByCriteria(0, co2Data, "co2")

	o2Int, err := strconv.ParseInt(o2Rating, 2, 64)
	if err != nil {
		panic(err)
	}
	co2Int, err := strconv.ParseInt(co2Rating, 2, 64)
	if err != nil {
		panic(err)
	}

	return o2Int * co2Int
}
