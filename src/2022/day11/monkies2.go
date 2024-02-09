package Monkies1

func main() {
	// definitely not going to write a parser.
	// definitely not going to write a parser.
	m0 := monkey{
		name:  "Monkey 0",
		items: []int{52, 60, 85, 69, 75, 75},
		operation: func(old int) int {
			return old * 17
		},
		test:            divisble_by_n_fn(13),
		throw_to:        map[bool]int{true: 6, false: 7},
		inspect_counter: 0,
	}
	m1 := monkey{
		name:  "Monkey 1",
		items: []int{96, 82, 61, 99, 82, 84, 85},
		operation: func(old int) int {
			return old + 8
		},
		test:            divisble_by_n_fn(7),
		throw_to:        map[bool]int{true: 0, false: 7},
		inspect_counter: 0,
	}
	m2 := monkey{
		name:  "Monkey 2",
		items: []int{95, 79},
		operation: func(old int) int {
			return old + 6
		},
		test:            divisble_by_n_fn(19),
		throw_to:        map[bool]int{true: 5, false: 3},
		inspect_counter: 0,
	}
	m3 := monkey{
		name:  "Monkey 3",
		items: []int{88, 50, 82, 65, 77},
		operation: func(old int) int {
			return old * 19
		},
		test:            divisble_by_n_fn(2),
		throw_to:        map[bool]int{true: 4, false: 1},
		inspect_counter: 0,
	}
	m4 := monkey{
		name:  "Monkey 4",
		items: []int{66, 90, 59, 90, 87, 63, 53, 88},
		operation: func(old int) int {
			return old + 7
		},
		test:            divisble_by_n_fn(5),
		throw_to:        map[bool]int{true: 1, false: 0},
		inspect_counter: 0,
	}
	m5 := monkey{
		name:  "Monkey 5",
		items: []int{92, 75, 62},
		operation: func(old int) int {
			return old * old
		},
		test:            divisble_by_n_fn(3),
		throw_to:        map[bool]int{true: 3, false: 4},
		inspect_counter: 0,
	}
	m6 := monkey{
		name:  "Monkey 6",
		items: []int{94, 86, 76, 67},
		operation: func(old int) int {
			return old + 1
		},
		test:            divisble_by_n_fn(11),
		throw_to:        map[bool]int{true: 5, false: 2},
		inspect_counter: 0,
	}
	m7 := monkey{
		name:  "Monkey 7",
		items: []int{57},
		operation: func(old int) int {
			return old + 2
		},
		test:            divisble_by_n_fn(17),
		throw_to:        map[bool]int{true: 6, false: 2},
		inspect_counter: 0,
	}

}
