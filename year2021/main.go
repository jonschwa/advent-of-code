package year2021

func RunDay(day int) {
	handlers := map[int]func(){
		1: RunDay1,
		2: RunDay2,
		3: RunDay3,
	}

	handlers[day]()
}
