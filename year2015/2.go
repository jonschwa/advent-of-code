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

func (g *Gift) getRequiredRibbon() int {
	sides := []int{g.l, g.h, g.w}
	maxIdx := 0
	biggestSide := 0

	// iterate, get the idx we want to pop (man this sucks in golang)
	for i, size := range sides {
		if i == 0 || size > biggestSide {
			biggestSide = size
			maxIdx = i
		}
	}

	// get the two smallest sides and calc perimeter
	sides[maxIdx] = sides[len(sides)-1] // Copy last element to index i.
	sides[len(sides)-1] = 0             // Erase last element (write zero value).
	minSides := sides[:len(sides)-1]    // new slice = truncated slice

	perimeter := (minSides[0] + minSides[1]) * 2
	volume := g.l * g.w * g.h
	requiredRibbon := perimeter + volume
	return requiredRibbon
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
	data := helpers.GetDataForDay(2015, 2)
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
	totalRibbonNeeded := 0
	for _, dim := range data {
		dimensions := strings.Split(dim, "x")
		gift := createGift(dimensions)
		gift.calculateAreas()
		totalRibbonNeeded += gift.getRequiredRibbon()
	}
	return totalRibbonNeeded
}
