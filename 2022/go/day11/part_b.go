package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	name            string
	items           []int
	operation       func(old int) int
	test            func(old int) bool
	throw_to        map[bool]int
	inspect_counter int
}

// monkey methods:
func (m *monkey) get_bored(item int, divide_by int) int {
	ans := float64(item / divide_by)
	return int(ans)
}

func (m *monkey) dequeue_first_item() int {
	dequeue := m.items[0]
	m.items = m.items[1:]
	m.inspect_counter += 1
	return dequeue
}

func (m *monkey) enqueue_item(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) inspect_and_operate() int {
	// inspect, worry increse ,bordem.
	item := m.dequeue_first_item()
	result := m.operation(item)
	calming := result % 9699690 //(13 * 7 * 19 * 2 * 5 * 3 * 11 * 17) // LCM of divisors
	// calming := result % 96577
	// bordem := m.get_bored(calming, 3)
	// bordem := m.get_bored(result, 3)
	return calming
}

func (m *monkey) perform_turn() (int, int) {
	item := m.inspect_and_operate()
	pass_to := m.throw_to[m.test(item)]
	// fmt.Println(m.name, "passes", item, "to", pass_to)
	return item, pass_to
}

type troop struct {
	members []monkey
}

// troop methods
func (t *troop) total_inspect_tally() []int {
	counts := make([]int, len(t.members))
	for i, m := range t.members {
		counts[i] = m.inspect_counter
		// counts = append(counts, m.inspect_counter)
	}
	// sort.Ints(counts)
	return counts

}

func (t *troop) perform_round() {
	for m := 0; m < len(t.members); m++ {
		// fmt.Println("m", m, t.members[m])
		for len(t.members[m].items) > 0 {
			// fmt.Println(len(t.members[m].items))
			item, recipient := t.members[m].perform_turn()
			t.members[recipient].enqueue_item(item)

		}
	}
}

// other
func divisble_by_n_fn(n int) func(old int) bool {
	return func(old int) bool {
		return old%n == 0
	}
}

func main() {
	// m0_example := monkey{
	// 	name:  "Monkey 0",
	// 	items: []int{79, 98},
	// 	operation: func(old int) int {
	// 		return old * 19
	// 	},
	// 	test:            divisble_by_n_fn(23),
	// 	throw_to:        map[bool]int{true: 2, false: 3},
	// 	inspect_counter: 0,
	// }
	// m1_example := monkey{
	// 	name:  "Monkey 1",
	// 	items: []int{54, 65, 75, 74},
	// 	operation: func(old int) int {
	// 		return old + 6
	// 	},
	// 	test:            divisble_by_n_fn(19),
	// 	throw_to:        map[bool]int{true: 2, false: 0},
	// 	inspect_counter: 0,
	// }
	// m2_example := monkey{
	// 	name:  "Monkey 2",
	// 	items: []int{79, 60, 97},
	// 	operation: func(old int) int {
	// 		return old * old
	// 	},
	// 	test:            divisble_by_n_fn(13),
	// 	throw_to:        map[bool]int{true: 1, false: 3},
	// 	inspect_counter: 0,
	// }
	// m3_example := monkey{
	// 	name:  "Monkey 3",
	// 	items: []int{74},
	// 	operation: func(old int) int {
	// 		return old + 3
	// 	},
	// 	test:            divisble_by_n_fn(17),
	// 	throw_to:        map[bool]int{true: 0, false: 1},
	// 	inspect_counter: 0,
	// }

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
	// t := troop{members: []monkey{m0_example, m1_example, m2_example, m3_example}}
	t := troop{members: []monkey{m0, m1, m2, m3, m4, m5, m6, m7}}
	fmt.Println("troop size:", len(t.members))
	// t := t_real

	// fmt.Println("old", t_example)
	// fmt.Println("new", t_real)
	// fmt.Println("")

	rounds := 10000
	for ii := 0; ii < rounds; ii++ {
		t.perform_round()
	}
	passes := t.total_inspect_tally()
	fmt.Println("tally unsorted:", passes)
	// fmt.Println("tally test2:", t_example.total_inspect_tally())
	// fmt.Println("tally test3:", t_real.total_inspect_tally())

	sort.Ints(passes)
	fmt.Println("tally sorted:", passes)
	fmt.Println("ans", passes[len(passes)-1]*passes[len(passes)-2])

	// print the troop
	for i := 0; i < len(t.members); i++ {
		fmt.Println(t.members[i])
	}

}
