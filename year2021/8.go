package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strconv"
	"strings"
)

type SegmentSeq struct {
	signals      []string
	outputs      []string
	charIsKnown  map[string]bool
	decodedChars [10]string
	charToInt    map[string]int
}

func (s *SegmentSeq) getUniqueChars() {
	for _, val := range s.signals {
		if len(val) == 2 {
			s.decodedChars[1] = val
			s.charIsKnown[val] = true
			s.charToInt[val] = 1
		} else if len(val) == 3 {
			s.decodedChars[7] = val
			s.charIsKnown[val] = true
			s.charToInt[val] = 7
		} else if len(val) == 4 {
			s.decodedChars[4] = val
			s.charIsKnown[val] = true
			s.charToInt[val] = 4
		} else if len(val) == 7 {
			s.decodedChars[8] = val
			s.charIsKnown[val] = true
			s.charToInt[val] = 8
		}
	}
}

func (s *SegmentSeq) get6Chars(chars []string) {
	//6 is the only one that does not contain 1
	// 9 is the one containing both 1 and 4
	// otherwise 0
	for _, char := range chars {
		oneChars := 0
		fourChars := 0

		for _, c := range s.decodedChars[1] {
			for _, find := range char {
				if c == find {
					oneChars++
				}
			}
		}
		for _, c := range s.decodedChars[4] {
			for _, find := range char {
				if c == find {
					fourChars++
				}
			}
		}
		if oneChars == len(s.decodedChars[1]) {
			if fourChars == len(s.decodedChars[4]) {
				s.decodedChars[9] = char
				s.charIsKnown[char] = true
				s.charToInt[char] = 9
			} else {
				s.decodedChars[0] = char
				s.charIsKnown[char] = true
				s.charToInt[char] = 0
			}
		} else {
			s.decodedChars[6] = char
			s.charIsKnown[char] = true
			s.charToInt[char] = 6
		}
	}
}

func (s *SegmentSeq) get5Chars(chars []string) {
	for _, char := range chars {
		oneChars := 0  // 3 has all of the same chars as 1
		nineChars := 0 // 5 has 5 chars in common w/ 9
		// 2 has neither

		for _, c := range s.decodedChars[1] {
			for _, find := range char {
				if c == find {
					oneChars++
				}
			}
		}
		for _, c := range s.decodedChars[9] {
			for _, find := range char {
				if c == find {
					nineChars++
				}
			}
		}

		if oneChars == len(s.decodedChars[1]) {
			s.decodedChars[3] = char
			s.charIsKnown[char] = true
			s.charToInt[char] = 3
		} else if nineChars == 5 {
			s.decodedChars[5] = char
			s.charIsKnown[char] = true
			s.charToInt[char] = 5
		} else {
			s.decodedChars[2] = char
			s.charIsKnown[char] = true
			s.charToInt[char] = 2
		}
	}
}

func (s SegmentSeq) derive5digits(seq string) int {
	num2Match := 0
	num3Match := 0
	num5Match := 0

	for _, digit := range seq {
		for _, m := range s.decodedChars[2] {
			if digit == m {
				num2Match++
			}
		}
		for _, m := range s.decodedChars[3] {
			if digit == m {
				num3Match++
			}
		}
		for _, m := range s.decodedChars[5] {
			if digit == m {
				num5Match++
			}
		}
	}

	if num5Match == 5 {
		return 5
	} else if num2Match == 5 {
		return 2
	} else if num3Match == 5 {
		return 3
	}
	panic("couldn't resolve 5 digit num")
}

func (s SegmentSeq) derive6digits(seq string) int {
	num6Match := 0
	num9Match := 0
	num0Match := 0

	for _, digit := range seq {
		for _, m := range s.decodedChars[6] {
			if digit == m {
				num6Match++
			}
		}
		for _, m := range s.decodedChars[9] {
			if digit == m {
				num9Match++
			}
		}
		for _, m := range s.decodedChars[0] {
			if digit == m {
				num0Match++
			}
		}
	}

	if num6Match == 6 {
		return 6
	} else if num9Match == 6 {
		return 9
	} else if num0Match == 6 {
		return 0
	}
	panic("couldn't resolve 6 digit num")
}

func (s *SegmentSeq) getOutputAsInt() int {
	digits := []int{}
	for _, seq := range s.outputs {
		var digit int
		if s.charIsKnown[seq] {
			digit = s.charToInt[seq]
		} else if len(seq) == 2 {
			digit = 1
		} else if len(seq) == 3 {
			digit = 7
		} else if len(seq) == 4 {
			digit = 4
		} else if len(seq) == 7 {
			digit = 8
		} else if len(seq) == 5 {
			digit = s.derive5digits(seq)
		} else if len(seq) == 6 {
			digit = s.derive6digits(seq)
		}
		digits = append(digits, digit)
	}

	stringNum := ""
	for _, d := range digits {
		stringNum += fmt.Sprint(d)
	}
	composedInt, err := strconv.Atoi(stringNum)
	if err != nil {
		panic(err)
	}
	return composedInt
}

func RunDay8() {
	data := helpers.GetDataForDay(2021, 8)
	fmt.Println("Part 1:", solveDay8Part1(data))
	fmt.Println("Part 2:", solveDay8Part2(data))
}

func solveDay8Part1(data []string) int {
	var numDigits int
	for _, line := range data {
		numDigits += getNumUniqueDigitsInLine(line)
	}
	return numDigits
}

func solveDay8Part2(data []string) int {
	var intVal int
	for _, line := range data {
		num := getValueOfDigitsInLine(line)
		intVal += num
	}
	return intVal
}

func getSegmentSeqFromLine(line string) SegmentSeq {
	arr := strings.Fields(line)
	foundMap := make(map[string]bool)
	var knownInts [10]string
	charToInt := make(map[string]int)
	return SegmentSeq{
		arr[:10],
		arr[11:],
		foundMap,
		knownInts,
		charToInt,
	}
}

func getNumUniqueDigitsInLine(line string) int {
	var numDigits int
	segmentSeq := getSegmentSeqFromLine(line)
	for _, val := range segmentSeq.outputs {
		if len(val) == 2 || len(val) == 3 || len(val) == 4 || len(val) == 7 {
			numDigits++
		}
	}
	return numDigits
}

func getValueOfDigitsInLine(line string) int {
	// first, get the unique chars...
	segmentSeq := getSegmentSeqFromLine(line)
	segmentSeq.getUniqueChars()

	// then lets get 6 chars first, and then 5 chars
	fiveChars := []string{}
	sixChars := []string{}
	for _, val := range segmentSeq.signals {
		if !segmentSeq.charIsKnown[val] {
			if len(val) == 6 {
				sixChars = append(sixChars, val)
			}
			if len(val) == 5 {
				fiveChars = append(fiveChars, val)
			}
		}
	}
	segmentSeq.get6Chars(sixChars)
	segmentSeq.get5Chars(fiveChars)
	return segmentSeq.getOutputAsInt()
}
