package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var data string

type Set struct {
	blue  int
	red   int
	green int
}

func SetFromString(s string) *Set {
	result := &Set{}
	colorStrings := strings.Split(s, ", ")
	for _, c := range colorStrings {
		parts := strings.Split(c, " ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("Could not split for 2 parts color string. Parts: %v, string: %v\n", parts, s))
		}
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("Could not parse color value <%v> as integer for string %v\n", parts[0], s))
		}

		color := parts[1]
		switch color {
		case "blue":
			result.blue = value
		case "red":
			result.red = value
		case "green":
			result.green = value
		default:
			panic(fmt.Sprintf("Found unknown color <%v> in string %v\n", color, s))
		}
	}
	return result
}

type Game struct {
	id   int
	sets []Set
}

func GameFromLine(line string) *Game {
	result := &Game{}
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("There are more than 2 parts split by ':' for line %v\n", line))
	}

	left := parts[0]
	id := strings.TrimPrefix(left, "Game ")
	var err error
	result.id, err = strconv.Atoi(id)
	if err != nil {
		panic(fmt.Sprintf("Game id <%v> was not parsed for left <%v> for line <%v>\n", id, left, line))
	}

	right := parts[1]
	setStrings := strings.Split(right, "; ")
	for _, s := range setStrings {
		result.sets = append(result.sets, *SetFromString(s))
	}
	return result
}

func part1() {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	idSum := 0
	games := []Game{}
	for _, l := range lines {
		g := *GameFromLine(l)
		games = append(games, g)

		ok := true
		for _, s := range g.sets {
			if s.red > maxRed {
				ok = false
			}
			if s.blue > maxBlue {
				ok = false
			}
			if s.green > maxGreen {
				ok = false
			}
		}
		if ok {
			idSum += g.id
		}
	}
	fmt.Println(idSum)
}

func part2() {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	totalPower := 0
	games := []Game{}
	for _, l := range lines {
		g := *GameFromLine(l)
		games = append(games, g)

		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, s := range g.sets {
			if s.blue > maxBlue {
				maxBlue = s.blue
			}
			if s.red > maxRed {
				maxRed = s.red
			}
			if s.green > maxGreen {
				maxGreen = s.green
			}
		}
		power := maxRed * maxBlue * maxGreen
		totalPower += power
	}
	fmt.Println(totalPower)
}

func main() {
	part1()
	part2()
}
