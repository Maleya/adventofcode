package main

import (
	"fmt"
	"math"
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
	return int(math.Round(ans))
}

func (m *monkey) dequeue_first_item() int {
	dequeue := m.items[0]
	m.items = m.items[1:]
	// fmt.Println(m.inspect_counter)
	m.inspect_counter += 1
	// fmt.Println(m.inspect_counter)
	return dequeue
}

func (m *monkey) enqueue_item(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) inspect_and_operate() int {
	// inspect, worry increse ,bordem.
	item := m.dequeue_first_item()
	result := m.operation(item)
	return m.get_bored(result, 3)
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
func (t *troop) total_inspect_tally() {
	for _, m := range t.members {
		count := m.inspect_counter
		fmt.Println(count)
	}
}

func (t *troop) perform_round() {
	for _, monkey := range t.members {
		for len(monkey.items) > 0 {
			// monkey.dequeue_first_item()
			item, recipient := monkey.perform_turn()
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
	// 	reminder: your queue pops are bad.

	// definitely not going to write a parser.
	m0 := monkey{
		name:  "Monkey 0",
		items: []int{79, 98},
		operation: func(old int) int {
			return old * 19
		},
		test:            divisble_by_n_fn(11),
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
	// fmt.Println(t)

	// m0 examples

	// THIS IS THE PROBLEM:
	fmt.Println("before:", t.members[0].inspect_counter)
	t.members[0].perform_turn()
	fmt.Println("after:", t.members[0].inspect_counter)

	// THE METHODS OF THE TROOP DONT WORK?
	// t.perform_round()
	// t.total_inspect_tally()
	// for _, m := range t.members {
	// 	fmt.Println(m)
	// }

	// fmt.Println("all items before pop:", m0.items)

	// // item := m0.items[0]
	// item := m0.dequeue_first_item()
	// fmt.Println("pre worry item:", item)
	// fmt.Println("all items after pop:", m0.items)
	// worry := m0.operation(item)
	// fmt.Println("post worry item:", worry)
	// bored := m0.get_bored(worry, 3)
	// fmt.Println("got bored:", bored)
	// result := m0.test(bored)
	// fmt.Println("result of test:", result)
	// throw := m0.throw_to[result]
	// fmt.Println("will throw to", throw)

	// input, _ := ioutil.ReadFile("example.txt")
	// input, _ := ioutil.ReadFile("example1.txt")
	// input, _ := ioutil.ReadFile("input.txt")
	// splitInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	// for i := 0; i < len(splitInput); i++ {
	// 	fmt.Println(splitInput[i])
	// 	// fmt.Sscanf(splitInput[i], "%s %d", &inst.command, &inst.arg)

	// }
	// fmt.Println("sum", sum)
	// make terminal 40 wide.

}
