package Monkies1

func main() {
	// definitely not going to write a parser.
	m0_part_a := monkey{
		name:  "Monkey 0",
		items: []int{79, 98},
		operation: func(old int) int {
			return old * 19
		},
		test:            divisble_by_n_fn(11),
		throw_to:        map[bool]int{true: 2, false: 3},
		inspect_counter: 0,
	}
	m1_part_a := monkey{
		name:  "Monkey 1",
		items: []int{54, 65, 75, 74},
		operation: func(old int) int {
			return old + 6
		},
		test:            divisble_by_n_fn(19),
		throw_to:        map[bool]int{true: 2, false: 0},
		inspect_counter: 0,
	}
	m2_part_a := monkey{
		name:  "Monkey 2",
		items: []int{79, 60, 97},
		operation: func(old int) int {
			return old * old
		},
		test:            divisble_by_n_fn(13),
		throw_to:        map[bool]int{true: 1, false: 3},
		inspect_counter: 0,
	}
	m3_part_a := monkey{
		name:  "Monkey 3",
		items: []int{74},
		operation: func(old int) int {
			return old + 3
		},
		test:            divisble_by_n_fn(17),
		throw_to:        map[bool]int{true: 0, false: 1},
		inspect_counter: 0,
	}
}
