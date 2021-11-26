package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func GetDataForDay(year int, day int) ([]string, error) {
	filename := fmt.Sprintf("data/%d/%d.txt", year, day)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
