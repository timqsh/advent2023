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

var CardValue = map[uint8]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
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

func (h Hand) Power() int {
	count := make(map[uint8]int)
	for _, c := range h.Cards {
		count[c]++
	}
	counts := make(map[int]bool)
	for _, v := range count {
		counts[v] = true
	}
	if counts[5] {
		assert(len(count) == 1)
		return 10
	}
	if counts[4] {
		assert(len(count) == 2)
		return 9
	}
	if counts[3] {
		if counts[2] {
			assert(len(count) == 2)
			return 8
		}
		assert(len(count) == 3)
		return 7
	}
	if counts[2] {
		if len(count) == 3 {
			return 6
		}
		assert(len(count) == 4)
		return 5
	}
	assert(len(count) == 5)
	return 4
}

func CmpHand(a, b Hand) int {
	aPow := a.Power()
	bPow := b.Power()
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

func main() {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		hands[i] = HandFromString(line)
	}
	winnings := 0
	slices.SortFunc(hands, CmpHand)
	for i, h := range hands {
		winnings += h.Bid * (i + 1)
		//fmt.Printf("%v %v\n", h, h.Power())
	}
	//248629404 is to high
	//248137488 is to low
	fmt.Println(winnings)
}
