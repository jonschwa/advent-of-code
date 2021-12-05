package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"regexp"
)

type Grid [][]int

type Coords struct {
	x int
	y int
}

type Line struct {
	origin   Coords
	endpoint Coords
}

func (g Grid) Print() {
	for _, row := range g {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Print("\n")
	}
}

func (g Grid) drawHorizontalLine(line Line) {
	var start int
	var end int
	if line.origin.y < line.endpoint.y {
		start = line.origin.y
		end = line.endpoint.y
	} else {
		start = line.endpoint.y
		end = line.origin.y
	}
	for i := start; i <= end; i++ {
		g[i][line.origin.x]++
	}
}

func (g Grid) drawVerticalLine(line Line) {
	var start int
	var end int
	if line.origin.x < line.endpoint.x {
		start = line.origin.x
		end = line.endpoint.x
	} else {
		start = line.endpoint.x
		end = line.origin.x
	}
	for i := start; i <= end; i++ {
		g[line.origin.y][i]++
	}
}

func (g Grid) drawDiagonalLine(line Line) {
	var dX int
	var dY int

	if line.origin.x < line.endpoint.x {
		dX = 1
	} else {
		dX = -1
	}

	if line.origin.y < line.endpoint.y {
		dY = 1
	} else {
		dY = -1
	}

	yVal := line.origin.y
	if dX > 0 {
		for i := line.origin.x; i <= line.endpoint.x; i++ {
			g[yVal][i]++
			yVal += dY
		}
	} else {
		for i := line.origin.x; i >= line.endpoint.x; i-- {
			g[yVal][i]++
			yVal += dY
		}
	}
}

// for a line, get the max X/Y Coordinates for grid sizing
func (l Line) getMaxCoords() (int, int) {
	var maxX int
	var maxY int

	if l.origin.x > l.endpoint.x {
		maxX = l.origin.x
	} else {
		maxX = l.endpoint.x
	}

	if l.origin.y > l.endpoint.y {
		maxY = l.origin.y
	} else {
		maxY = l.endpoint.y
	}

	return maxX, maxY
}

func createLinesAndGrid(data []string) ([]Line, Grid) {
	re, _ := regexp.Compile(`(\d{1,3}),(\d{1,3})\s->\s(\d{1,3}),(\d{1,3})`)
	lines := make([]Line, 0)
	gridWidth := 0
	gridHeight := 0
	for _, line := range data {

		// idx 0 is the og string, then the rest are the subpattern matches
		matches := re.FindStringSubmatch(line)
		intMatches := helpers.ConvertStringSliceToIntSlice(matches[1:])
		originCoords := Coords{x: intMatches[0], y: intMatches[1]}
		endpointCoords := Coords{x: intMatches[2], y: intMatches[3]}

		line := Line{originCoords, endpointCoords}
		lines = append(lines, line)

		lMaxX, lMaxY := line.getMaxCoords()
		if gridWidth < lMaxX {
			gridWidth = lMaxX
		}
		if gridHeight < lMaxY {
			gridHeight = lMaxY
		}
	}

	grid := make([][]int, 0)

	var i int
	for i = 0; i <= gridHeight; i++ {
		grid = append(grid, make([]int, gridWidth+1))
	}

	return lines, grid
}

func RunDay5() {
	data := helpers.GetDataForDay(2021, 5)
	fmt.Println("Part 1:", solveDay5Part1(data))
	fmt.Println("Part 2:", solveDay5Part2(data))

}

func solveDay5Part1(data []string) int {
	// parse data into slice of lines
	lines, grid := createLinesAndGrid(data)

	for _, line := range lines {
		// For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.
		if line.origin.x == line.endpoint.x {
			grid.drawHorizontalLine(line)
		} else if line.origin.y == line.endpoint.y {
			grid.drawVerticalLine(line)
		}
	}

	// find places where more than 2 lines intersect
	numDangerSpots := 0
	for _, row := range grid {
		for _, num := range row {
			if num > 1 {
				numDangerSpots++
			}
		}
	}
	return numDangerSpots
}

func solveDay5Part2(data []string) int {
	// parse data into slice of lines
	lines, grid := createLinesAndGrid(data)
	for _, line := range lines {
		// Because of the limits of the hydrothermal vent mapping system, the lines in your list will only ever be horizontal, vertical,
		// or a diagonal line at exactly 45 degrees (slope +-1)
		if line.origin.x == line.endpoint.x {
			grid.drawHorizontalLine(line)
		} else if line.origin.y == line.endpoint.y {
			grid.drawVerticalLine(line)
		} else {
			grid.drawDiagonalLine(line)
		}
	}
	// find places where more than 2 lines intersect
	numDangerSpots := 0
	for _, row := range grid {
		for _, num := range row {
			if num > 1 {
				numDangerSpots++
			}
		}
	}
	return numDangerSpots
}
