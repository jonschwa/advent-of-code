package year2015

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strconv"
	"strings"
)

type Gift struct {
	l             int
	w             int
	h             int
	areas         []int
	minArea       int
	paperRequired int
}

func (g *Gift) calculateAreas() {
	side1 := g.l * g.w
	side2 := g.w * g.h
	side3 := g.h * g.l

	g.areas = []int{side1, side2, side3}
	var minArea int
	var totalArea int

	for i, area := range g.areas {
		totalArea += 2 * area
		if i == 0 || area < minArea {
			minArea = area
		}
	}

	g.paperRequired = totalArea + minArea
	g.minArea = minArea
}

func createGift(dimensions []string) Gift {
	l, err := strconv.Atoi(dimensions[0])
	if err != nil {
		fmt.Println("Error:", err)
	}
	w, err := strconv.Atoi(dimensions[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	h, err := strconv.Atoi(dimensions[2])
	if err != nil {
		fmt.Println("Error:", err)
	}

	gift := Gift{l: l, w: w, h: h}
	return gift
}

func RunDay2() {
	data, err := helpers.GetDataForDay(2015, 2)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}
	fmt.Println("Part 1:", solveDay2Part1(data))
	fmt.Println("Part 2:", solveDay2Part2(data))

}

func solveDay2Part1(data []string) int {
	totalPaperNeeded := 0
	for _, dim := range data {
		dimensions := strings.Split(dim, "x")
		gift := createGift(dimensions)
		gift.calculateAreas()
		totalPaperNeeded += gift.paperRequired
	}
	return totalPaperNeeded

}

func solveDay2Part2(data []string) int {
	return 0
}
