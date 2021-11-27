package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"jonschwa/aoc/year2015"
	"jonschwa/aoc/year2021"
	"os"
	"os/exec"
)

func main() {
	yearPtr := flag.Int("y", 0, "Which year is the problem?")
	dayPtr := flag.Int("d", 0, "Which day is the problem?")
	testPtr := flag.Bool("test", false, "Is this a test run?")

	flag.Parse()
	year := *yearPtr
	day := *dayPtr
	test := *testPtr

	if year == 0 || day == 0 {
		fmt.Println("Invalid day/year")
		os.Exit(1)
	}

	if test {
		filename := fmt.Sprintf("year%d/%d.go", year, day)
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			fmt.Println("No solution found for that year/day.")
			os.Exit(1)
		}

		fmt.Printf("Running test for %d: Day %d...\n", year, day)
		test_filename := fmt.Sprintf("year%d/%d_test.go", year, day)
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			fmt.Println("No test file found for that year/day.")
			os.Exit(1)
		}
		solution := exec.Command("go", "test", filename, test_filename)
		var out bytes.Buffer
		solution.Stdout = &out
		err := solution.Run()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out.String())
	}

	switch year {
	case 2015:
		run2015(day)
	case 2021:
		run2021(day)
	default:
		fmt.Println("some default")
	}
}

func run2015(day int) {
	year2015.RunDay(day)
}

func run2021(day int) {
	year2021.RunDay(day)
}
