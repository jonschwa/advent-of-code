package year2021

func RunDay(day int) {
	handlers := map[int]func(){
		1: RunDay1,
		2: RunDay2,
		3: RunDay3,
		4: RunDay4,
		5: RunDay5,
		6: RunDay6,
		7: RunDay7,
		8: RunDay8,
		9: RunDay9,
	}

	handlers[day]()
}
