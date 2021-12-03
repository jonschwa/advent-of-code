package year2021

func RunDay(day int) {
	handlers := map[int]func(){
		1: RunDay1,
		2: RunDay2,
	}

	handlers[day]()
}
