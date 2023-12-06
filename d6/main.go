package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed in.txt
var in string

type Race struct {
	Time     int
	Distance int
}

func Parse(s string, part int) []Race {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) != 2 {
		panic(fmt.Sprintf("Got %d lines, expected 2, for input:\n%s\n", len(lines), s))
	}
	if part == 2 {
		lines[0] = strings.TrimPrefix(lines[0], "Time: ")
		lines[1] = strings.TrimPrefix(lines[1], "Distance: ")
		lines[0] = strings.ReplaceAll(lines[0], " ", "")
		lines[1] = strings.ReplaceAll(lines[1], " ", "")
	}
	digitReg := regexp.MustCompile(`\d+`)
	times := digitReg.FindAllString(lines[0], -1)
	distances := digitReg.FindAllString(lines[1], -1)
	if len(times) != len(distances) {
		panic(fmt.Sprintf("Got %d times, %d distances, expected equal numbers, for input:\n%s\n", len(times), len(distances), s))
	}
	result := []Race{}
	for i := range times {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic(fmt.Sprintf("Got bad time %s, for input:\n%s\n", times[i], s))
		}
		dist, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(fmt.Sprintf("Got bad distance %s, for input:\n%s\n", distances[i], s))
		}
		result = append(result, Race{time, dist})
	}
	return result
}

func WaysToBeat(races []Race) int {
	waysToBeatAll := 1
	for _, race := range races {
		waysToBeatRace := 0
		for pressTime := 0; pressTime <= race.Time; pressTime++ {
			remaining := race.Time - pressTime
			speed := pressTime
			distance := speed * remaining
			if distance > race.Distance {
				waysToBeatRace++
			}
		}
		waysToBeatAll *= waysToBeatRace
	}
	return waysToBeatAll
}

func main() {
	races1 := Parse(in, 1)
	fmt.Println(WaysToBeat(races1))

	races2 := Parse(in, 2)
	fmt.Println(WaysToBeat(races2))
}
