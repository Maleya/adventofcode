package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// seven is hardcoded as the number of hands
type camelcards struct {
	// all_hands       []hand
	strengthToIndex map[int][]int
	handTypes       map[int][]hand
}

func (c *camelcards) read_input(input_lines []string) {
	c.handTypes = make(map[int][]hand)

	for _, line := range input_lines {

		var h hand
		line_split := strings.Fields(line)
		h.bid, _ = strconv.Atoi(line_split[1])
		h.cards = line_split[0]
		h.classify_hand()
		c.handTypes[h.hand_strength] = append(c.handTypes[h.hand_strength], h)
	}
}

func (c *camelcards) scoreTieBreaks() {
	no_of_hands := 7
	card_strength := map[string]int{
		"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1,
	}

	// for i, stacked := range c.handTypes {
	for i := 0; i < no_of_hands; i++ {
		// len(c.handtypes) is the number of keys..thats wrong. and you dont know the order..
		stacked, ok := c.handTypes[i]

		if !ok {
			// fmt.Println(i, "not in map")
			continue
		}
		if len(stacked) == 1 {
			// fmt.Println(i, "single element")
			continue
		}
		// if duplicate hands of same strength exist:
		for hand_idx := 0; hand_idx < len(c.handTypes[i]); hand_idx++ {
			multi := 100000000000
			// multi := 10000
			hand_sum := 0
			for card_idx := 0; card_idx < len(c.handTypes[i][hand_idx].cards); card_idx++ {
				// t := c.handTypes[i][hand_idx].cards
				card := string(c.handTypes[i][hand_idx].cards[card_idx])
				hand_sum += card_strength[card] * multi
				multi = multi / 100
			}
			c.handTypes[i][hand_idx].tieBreakScore = hand_sum

		}
	}
}
func (c *camelcards) total_winnings() int {
	output := 0
	for i, stacked := range c.handTypes {
		if len(stacked) == 1 {
			bid := c.handTypes[i][0].bid
			rank := c.handTypes[i][0].rank
			prod := bid * rank
			output = output + prod

		}
		if len(stacked) > 1 {
			for j := 0; j < len(c.handTypes[i]); j++ {
				bid := c.handTypes[i][j].bid
				rank := c.handTypes[i][j].rank
				prod := bid * rank
				output = output + prod
			}
		}
	}
	return output
}

func (c *camelcards) rank_hands_new() {
	rank_iter := 1
	no_of_hands := 7
	for i := 0; i < no_of_hands; i++ {
		value, ok := c.handTypes[i]
		if !ok {
			continue
		}

		if len(value) == 1 {
			c.handTypes[i][0].rank = rank_iter
			fmt.Println(rank_iter, c.handTypes[i][0])
			rank_iter++
		}
		if len(value) > 1 {
			sort.SliceStable(c.handTypes[i], func(i, j int) bool {
				return value[i].tieBreakScore < value[j].tieBreakScore
			})

			for j := 0; j < len(value); j++ {
				c.handTypes[i][j].rank = rank_iter
				fmt.Println(rank_iter, c.handTypes[i][j])
				rank_iter++

			}

		}
	}
}

type hand struct {
	cards         string
	hand_name     string
	hand_strength int
	tieBreakScore int
	bid           int
	rank          int
}

func (h *hand) classify_hand() {
	card_count, highest_duplicate := count_occurances(h.cards)
	distinct_keys := len(card_count)

	switch {
	case distinct_keys == 1:
		h.hand_name = "Five of a kind"
		h.hand_strength = 6

	case distinct_keys == 2:
		if highest_duplicate == 4 {
			h.hand_name = "Four of a kind"
			h.hand_strength = 5
		} else {
			h.hand_name = "Full house"
			h.hand_strength = 4
		}
	case distinct_keys == 3:
		if highest_duplicate == 3 {
			h.hand_name = "Three of a kind"
			h.hand_strength = 3
		} else {
			h.hand_name = "Two pair"
			h.hand_strength = 2
		}
	case distinct_keys == 4:
		h.hand_name = "One pair"
		h.hand_strength = 1

	case distinct_keys == 5:
		h.hand_name = "High card"
		h.hand_strength = 0
	}

}

func count_occurances(cards string) (map[string]int, int) {
	highest_count := 0
	card_count := map[string]int{}
	for _, card := range cards {
		card_count[string(card)]++
		if card_count[string(card)] > highest_count {
			highest_count = card_count[string(card)]
		}
	}
	return card_count, highest_count
}
func parse_input_line(line string) hand {
	var h hand

	line_split := strings.Fields(line)
	h.bid, _ = strconv.Atoi(line_split[1])
	h.cards = line_split[0]
	h.classify_hand()
	return h

}

func main() {
	// fileName := "example.txt"
	// fileName := "example_1.txt"
	fileName := "input.txt"

	file, _ := os.Open(fileName)
	defer file.Close()
	content, _ := io.ReadAll(file)
	var camel_game camelcards

	splitInput := strings.Split(strings.TrimSpace(string(content)), "\n")

	camel_game.read_input(splitInput)
	camel_game.scoreTieBreaks()
	camel_game.rank_hands_new()
	part_a := camel_game.total_winnings()
	fmt.Println("part a", part_a)

	// fmt.Println(camel_game.handTypes)

}
