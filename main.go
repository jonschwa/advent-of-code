package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
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

	// todo ensure file exists
	filename := fmt.Sprintf("years/%d/%d.go", year, day)
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No solution found for that year/day.")
		os.Exit(1)
	}

	// Solution may be a test run, or an actual run depending on the flags
	var solution *exec.Cmd
	if test {
		fmt.Printf("Running test for %d: Day %d...\n", year, day)
		test_filename := fmt.Sprintf("years/%d/%d_test.go", year, day)
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			fmt.Println("No test file found for that year/day.")
			os.Exit(1)
		}
		solution = exec.Command("go", "test", "-v", filename, test_filename)
	} else {
		fmt.Printf("Running solution for %d: Day %d...\n", year, day)
		solution = exec.Command("go", "run", filename)
	}
	var out bytes.Buffer
	solution.Stdout = &out
	err := solution.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out.String())

}
