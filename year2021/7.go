package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"math"
	"strings"
)

func RunDay7() {
	data := helpers.GetDataForDay(2021, 7)
	intData := helpers.ConvertStringSliceToIntSlice(strings.Split(data[0], ","))
	fmt.Printf("Part 1: %d\n", int(solveDay7Part1(intData)))
	fmt.Printf("Part 2: %d\n", int(solveDay7Part2(intData)))

}

func solveDay7Part1(data []int) float64 {
	var cheapest float64
	options := map[int]bool{}

	for _, val := range data {
		if !options[val] {
			options[val] = true
		}
	}
	for num, _ := range options {
		run := getFuelCostConst(num, data)
		if cheapest == 0 || run < cheapest {
			cheapest = run
		}
	}
	return cheapest
}

func solveDay7Part2(data []int) float64 {
	var cheapest float64
	var min int
	var max int
	for _, val := range data {
		if min == 0 || val < min {
			min = val
		}
		if max == 0 || val > max {
			max = val
		}
	}
	for i := min; i <= max; i++ {
		run := getFuelCostVariable(i, data)
		if cheapest == 0 || run < cheapest {
			cheapest = run
		}
	}
	return cheapest
}

func getFuelCostConst(pos int, data []int) float64 {
	var cost float64
	for _, num := range data {
		dist := float64(pos) - float64(num)
		cost += math.Abs(dist)
	}
	return cost
}

func getFuelCostVariable(pos int, data []int) float64 {
	var cost float64
	for _, num := range data {
		dist := math.Abs(float64(pos) - float64(num))
		var i float64
		for i = 1; i <= dist; i++ {
			cost += i
		}
	}
	return cost
}
