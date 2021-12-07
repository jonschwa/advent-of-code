package year2021

import (
	"fmt"
	"jonschwa/aoc/helpers"
	"strings"
)

type LanternFish struct {
	timer int
}

func (lf LanternFish) IncreaseDay() (LanternFish, bool) {
	var newFish bool
	if lf.timer == 0 {
		lf.timer = 6
		newFish = true
	} else {
		lf.timer--
	}
	return lf, newFish
}

func RunDay6() {
	data := helpers.GetDataForDay(2021, 6)
	intData := helpers.ConvertStringSliceToIntSlice(strings.Split(data[0], ","))
	fmt.Println("Part 1:", solveDay6Part1(intData, 80))
	fmt.Println("Part 2:", solveDay6Part2(intData, 256))

}

func solveDay6Part1(data []int, days int) int {
	allFish := createFishContainer(data)
	for i := 1; i <= days; i++ {
		allFish = getNextDay(allFish)
	}
	return len(allFish)
}

func solveDay6Part2(data []int, days int) int {
	fishMap := createFishAgeMap(data)
	for i := 1; i <= days; i++ {
		fishMap = updateFishAgeMap(fishMap)
	}
	var fishCount int
	for _, num := range fishMap {
		fishCount += num
	}
	return fishCount
}

func createFishContainer(data []int) []LanternFish {
	fish := make([]LanternFish, 0)
	for _, timer := range data {
		fish = append(fish, LanternFish{timer})
	}
	return fish
}

func getNextDay(fish []LanternFish) []LanternFish {
	var nextDay []LanternFish
	for _, fish := range fish {
		updated, newFish := fish.IncreaseDay()
		nextDay = append(nextDay, updated)
		if newFish {
			nextDay = append(nextDay, LanternFish{8})
		}
	}
	return nextDay
}

func createFishAgeMap(data []int) map[int]int {
	fishMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
	for _, age := range data {
		fishMap[age]++
	}
	return fishMap
}

func updateFishAgeMap(fishMap map[int]int) map[int]int {
	// these fish are done and going to become 6s and 8s
	completedFish := fishMap[0]
	for i := 0; i < 8; i++ {
		fishMap[i] = fishMap[i+1]
	}
	//big life changes for the special aged fish
	fishMap[6] += completedFish
	fishMap[8] = completedFish

	return fishMap
}
