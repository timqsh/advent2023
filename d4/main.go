package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed in.txt
var s string

type Card struct {
	id   int
	win  map[int]bool
	have map[int]bool
}

func CardFromString(line string) Card {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("Can't split line into 2 parts. Got %v, line %v", parts, line))
	}
	id, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Card ")))
	if err != nil {
		panic(fmt.Sprintf("Can't parse id from %v, err: %v line: %v", parts[0], err, line))
	}
	parts = strings.Split(parts[1], " | ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("Can't split right part into 2 parts. Got %v, line %v", parts, line))
	}
	win := numsFromString(parts[0])
	have := numsFromString(parts[1])
	return Card{id: id, win: win, have: have}
}

func numsFromString(s string) map[int]bool {
	result := map[int]bool{}
	for _, s := range strings.Split(strings.TrimSpace(s), " ") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		w, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Can't parse number from %v, err: %v", s, err))
		}
		result[w] = true
	}
	return result
}

func (c Card) Matches() int {
	result := 0
	for key := range c.have {
		_, found := c.win[key]
		if found {
			result += 1
		}
	}
	return result
}

func main() {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	cards := []Card{}
	for _, l := range lines {
		c := CardFromString(l)
		cards = append(cards, c)
	}

	copies := []int{}
	for range cards {
		copies = append(copies, 1)
	}
	part1 := 0
	for curIdx, c := range cards {
		m := c.Matches()
		if m != 0 {
			part1 += 1 << (m - 1)
		}
		curCopies := copies[curIdx]
		for i := curIdx + 1; i <= curIdx+m && i < len(cards); i++ {
			copies[i] += curCopies
		}
	}
	part2 := 0
	for _, val := range copies {
		part2 += val
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
