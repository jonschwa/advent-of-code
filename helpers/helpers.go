package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetDataForDay(year int, day int) []string {
	fileName := fmt.Sprintf("data/%d/%d.txt", year, day)

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")
	return sliceData
}

func ConvertStringSliceToIntSlice(strings []string) []int {
	var err error
	intData := make([]int, len(strings))
	for i, v := range strings {
		intData[i], err = strconv.Atoi(v)
		if err != nil {
			panic("unable to convert!")
		}
	}
	return intData
}
