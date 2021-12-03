package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strconv"
	"strings"
)

type Command struct {
	dir string
	mag int
}

type Position struct {
	hor   int
	depth int
	aim   int
}

func (p *Position) runCommandV1(c Command) {
	if c.dir == "forward" {
		p.hor = p.hor + c.mag
	} else if c.dir == "up" {
		p.depth = p.depth - c.mag
	} else if c.dir == "down" {
		p.depth = p.depth + c.mag
	}
}

func (p *Position) runCommandV2(c Command) {
	if c.dir == "forward" {
		p.hor = p.hor + c.mag
		p.depth = p.depth + (p.aim * c.mag)
	} else if c.dir == "up" {
		p.aim = p.aim - c.mag
	} else if c.dir == "down" {
		p.aim = p.aim + c.mag
	}
}

func convertLinesToCommands(data []string) []Command {
	commands := make([]Command, len(data))
	for i, line := range data {
		fields := strings.Fields(line)
		dir, strMag := fields[0], fields[1]
		mag, err := strconv.Atoi(strMag)
		if err != nil {
			panic("error making commands")
		}
		commands[i] = Command{dir, mag}
	}
	return commands
}

func RunDay2() {
	data, err := helpers.GetDataForDay(2021, 2)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	commands := convertLinesToCommands(data)
	fmt.Println("Part 1:", solveDay2Part1(commands))
	fmt.Println("Part 2:", solveDay2Part2(commands))
}

func solveDay2Part1(data []Command) int {
	var position Position
	for _, command := range data {
		position.runCommandV1((command))
	}
	return position.hor * position.depth
}

func solveDay2Part2(data []Command) int {
	var position Position
	for _, command := range data {
		position.runCommandV2((command))
	}
	return position.hor * position.depth
}
