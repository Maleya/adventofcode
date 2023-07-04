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
	bordem := m.get_bored(result, 3)
	return bordem
}

func (m *monkey) perform_turn() (int, int) {
	item := m.inspect_and_operate()
	pass_to := m.throw_to[m.test(item)]
	fmt.Println(m.name, "passes", item, "to", pass_to)
	return item, pass_to

}

type troop struct {
	members []monkey
}

// troop methods
func (t *troop) total_inspect_tally() []int {
	counts := make([]int, len(t.members))
	for _, m := range t.members {
		counts = append(counts, m.inspect_counter)
	}
	sort.Ints(counts)
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
	// definitely not going to write a parser.
	m0 := monkey{
		name:  "Monkey 0",
		items: []int{79, 98},
		operation: func(old int) int {
			return old * 19
		},
		test:            divisble_by_n_fn(23),
		throw_to:        map[bool]int{true: 2, false: 3},
		inspect_counter: 0,
	}
	m1 := monkey{
		name:  "Monkey 1",
		items: []int{54, 65, 75, 74},
		operation: func(old int) int {
			return old + 6
		},
		test:            divisble_by_n_fn(19),
		throw_to:        map[bool]int{true: 2, false: 0},
		inspect_counter: 0,
	}
	m2 := monkey{
		name:  "Monkey 2",
		items: []int{79, 60, 97},
		operation: func(old int) int {
			return old * old
		},
		test:            divisble_by_n_fn(13),
		throw_to:        map[bool]int{true: 1, false: 3},
		inspect_counter: 0,
	}
	m3 := monkey{
		name:  "Monkey 3",
		items: []int{74},
		operation: func(old int) int {
			return old + 3
		},
		test:            divisble_by_n_fn(17),
		throw_to:        map[bool]int{true: 0, false: 1},
		inspect_counter: 0,
	}

	t := troop{members: []monkey{m0, m1, m2, m3}}

	rounds := 20
	for ii := 0; ii < rounds; ii++ {
		t.perform_round()
	}
	passes := t.total_inspect_tally()
	fmt.Println("ans", passes[len(passes)-1]*passes[len(passes)-2])

	// print the troop
	for i := 0; i < len(t.members); i++ {
		fmt.Println(t.members[i])
	}

}
