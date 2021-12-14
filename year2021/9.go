package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"sort"
	"strconv"
)

type CaveMap [][]int
type BasinMap map[string]bool

func (c CaveMap) checkIfLowPoint(x int, y int) bool {
	var isLowPoint bool
	num := c[x][y]
	maxX := len(c) - 1
	maxY := len(c[0]) - 1

	if x == 0 {
		// if this is the top row, only search below and next to
		if y == 0 {
			return c[x][y+1] > num && c[x+1][y] > num
		} else if 0 < y && y < maxY {
			return c[x][y+1] > num && c[x+1][y] > num && c[x][y-1] > num
		} else if y == maxY {
			return c[x][y-1] > num && c[x+1][y] > num
		}
	} else if x < maxX {
		// otherwise, search above, below, and next to
		if y == 0 {
			return c[x][y+1] > num && c[x+1][y] > num && c[x-1][y] > num
		} else if 0 < y && y < maxY {
			return c[x][y+1] > num && c[x+1][y] > num && c[x-1][y] > num && c[x][y-1] > num
		} else if y == maxY {
			return c[x][y-1] > num && c[x+1][y] > num && c[x-1][y] > num
		}
	} else if x == maxX {
		// if this is the bottom row, only search next to and above
		if y == 0 {
			return c[x][y+1] > num && c[x-1][y] > num
		} else if 0 < y && y < maxY {
			return c[x][y+1] > num && c[x-1][y] > num && c[x][y-1] > num
		} else if y == maxY {
			return c[x][y-1] > num && c[x-1][y] > num
		}
	}
	return isLowPoint
}

func (c CaveMap) getBasinSizeOfLowPoint(x int, y int) int {
	// maintain a map of found points so we don't double up
	basinMap := &BasinMap{}

	// starting from the low point, check the neighbors
	c.addNeighborsToBasinMap(x, y, basinMap)

	// if those neighbors are not 9, add them and put them on the queue to check their neighbors
	return len(*basinMap)
}

func (c CaveMap) addNeighborsToBasinMap(x int, y int, basinMap *BasinMap) {
	num := c[x][y]
	maxX := len(c) - 1
	maxY := len(c[0]) - 1

	if num < 9 {
		basinAddr := fmt.Sprintf("x%dy%d", x, y)
		if !(*basinMap)[basinAddr] {
			(*basinMap)[basinAddr] = true
			if x == 0 {
				// if this is the top row, only search below and next to
				if y == 0 {
					c.addNeighborsToBasinMap(x, y+1, basinMap)
					c.addNeighborsToBasinMap(x+1, y, basinMap)
				} else if 0 < y && y < maxY {
					c.addNeighborsToBasinMap(x, y+1, basinMap)
					c.addNeighborsToBasinMap(x+1, y, basinMap)
					c.addNeighborsToBasinMap(x, y-1, basinMap)
				} else if y == maxY {
					c.addNeighborsToBasinMap(x, y-1, basinMap)
					c.addNeighborsToBasinMap(x+1, y, basinMap)
				}
			} else if x < maxX {
				if y == 0 {
					c.addNeighborsToBasinMap(x, y+1, basinMap)
					c.addNeighborsToBasinMap(x+1, y, basinMap)
					c.addNeighborsToBasinMap(x-1, y, basinMap)
				} else if 0 < y && y < maxY {
					c.addNeighborsToBasinMap(x, y+1, basinMap)
					c.addNeighborsToBasinMap(x+1, y, basinMap)
					c.addNeighborsToBasinMap(x-1, y, basinMap)
					c.addNeighborsToBasinMap(x, y-1, basinMap)
				} else if y == maxY {
					c.addNeighborsToBasinMap(x-1, y, basinMap)
					c.addNeighborsToBasinMap(x, y-1, basinMap)
					c.addNeighborsToBasinMap(x+1, y, basinMap)
				}
			} else if x == maxX {
				if y == 0 {
					c.addNeighborsToBasinMap(x, y+1, basinMap)
					c.addNeighborsToBasinMap(x-1, y, basinMap)
				} else if 0 < y && y < maxY {
					c.addNeighborsToBasinMap(x, y+1, basinMap)
					c.addNeighborsToBasinMap(x-1, y, basinMap)
					c.addNeighborsToBasinMap(x, y-1, basinMap)
				} else if y == maxY {
					c.addNeighborsToBasinMap(x, y-1, basinMap)
					c.addNeighborsToBasinMap(x-1, y, basinMap)
				}
			}
		}
	}
}

func RunDay9() {
	data := helpers.GetDataForDay(2021, 9)
	fmt.Printf("Part 1: %d\n", int(solveDay9Part1(data)))
	fmt.Printf("Part 2: %d\n", int(solveDay9Part2(data)))
}

func solveDay9Part1(data []string) int {
	caveMap := newCaveMap(data)
	sumLowPoints := getSumLowPoints(caveMap)
	return sumLowPoints
}

func solveDay9Part2(data []string) int {
	caveMap := newCaveMap(data)
	largestBasinSum := getLargestBasinProduct(caveMap)
	return largestBasinSum
}

func newCaveMap(data []string) CaveMap {
	caveMap := CaveMap{}
	for _, row := range data {
		intRow := []int{}
		for _, char := range row {
			intVal, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			intRow = append(intRow, intVal)
		}
		caveMap = append(caveMap, intRow)
	}
	return caveMap
}

func getSumLowPoints(c CaveMap) int {
	var sumLowPoints int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			if c.checkIfLowPoint(i, j) {
				sumLowPoints += c[i][j] + 1
			}
		}
	}
	return sumLowPoints
}

func getLargestBasinProduct(c CaveMap) int {
	var basins []int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			if c.checkIfLowPoint(i, j) {
				basins = append(basins, c.getBasinSizeOfLowPoint(i, j))
			}
		}
	}
	// sort the array of basins and return product of largest 3
	sort.Sort(sort.Reverse(sort.IntSlice(basins[:])))

	return basins[0] * basins[1] * basins[2]
}
