package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed in.txt
var in string

var CardValue1 = map[uint8]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

// Jokers are 1
var CardValue2 = map[uint8]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14,
}

type Hand struct {
	Cards [5]uint8
	Bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%v %v", string(h.Cards[:]), h.Bid)
}

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func (h Hand) Power(part int) int {
	count := make(map[uint8]int)
	if part == 1 {
		for _, c := range h.Cards {
			count[c]++
		}
	} else {
		jokers := 0
		for _, c := range h.Cards {
			if c != 'J' {
				count[c]++
			} else {
				jokers++
			}
		}
		highestCard := uint8('J')
		highest := 0
		for card, num := range count {
			if num > highest {
				highest = num
				highestCard = card
			}
		}
		count[highestCard] += jokers
	}

	counts := make(map[int]bool)
	for _, v := range count {
		counts[v] = true
	}
	if counts[5] {
		assert(len(count) == 1)
		return 7 // five of a kind
	}
	if counts[4] {
		assert(len(count) == 2)
		return 6 // four of a kind
	}
	if counts[3] {
		if counts[2] {
			assert(len(count) == 2)
			return 5 // full house
		}
		assert(len(count) == 3)
		return 4 // three of a kind
	}
	if counts[2] {
		if len(count) == 3 {
			return 3 // two pair
		}
		assert(len(count) == 4)
		return 2 // one pair
	}
	assert(len(count) == 5)
	return 1 // high card
}

func CmpHand(part int) func(a, b Hand) int {
	var CardValue map[uint8]int
	if part == 1 {
		CardValue = CardValue1
	} else {
		CardValue = CardValue2
	}
	inner := func(a, b Hand) int {
		aPow := a.Power(part)
		bPow := b.Power(part)
		if aPow != bPow {
			return aPow - bPow
		}
		for i := 0; i < 5; i++ {
			if a.Cards[i] != b.Cards[i] {
				aVal, ok := CardValue[a.Cards[i]]
				assert(ok)
				bVal, ok := CardValue[b.Cards[i]]
				assert(ok)
				return aVal - bVal
			}
		}
		return 0
	}
	return inner
}

func HandFromString(s string) Hand {
	res := strings.Split(s, " ")
	if len(res) != 2 {
		panic(fmt.Sprintf("got %v parts, expected 2, for string %v", len(res), s))
	}
	h := Hand{}
	for i := 0; i < 5; i++ {
		h.Cards[i] = s[i]
	}
	bid, err := strconv.Atoi(res[1])
	if err != nil {
		panic(fmt.Sprintf("failed to parse bid %v", res[1]))
	}
	h.Bid = bid
	return h
}

func solve(hands []Hand, part int) int {
	winnings := 0
	slices.SortFunc(hands, CmpHand(part))
	for i, h := range hands {
		winnings += h.Bid * (i + 1)
		//fmt.Printf("%v %v\n", h, h.Power(part))
	}
	return winnings
}

func main() {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		hands[i] = HandFromString(line)
	}
	p1 := solve(hands, 1)
	fmt.Println(p1)
	p2 := solve(hands, 2)
	fmt.Println(p2)

}
