package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed in.txt
var in string

const (
	Left  = 0
	Right = 1
)

type Net map[string][2]string
type Path []int

type Instructions struct {
	Net  Net
	Path Path
}

func InstructionsFromString(s string) Instructions {
	parts := strings.Split(strings.TrimSpace(s), "\n\n")
	if len(parts) != 2 {
		panic("invalid input")
	}
	path := make(Path, 0)
	for _, rune := range parts[0] {
		if rune == 'L' {
			path = append(path, Left)
		} else if rune == 'R' {
			path = append(path, Right)
		} else {
			panic("invalid input")
		}
	}
	nodes := strings.Split(parts[1], "\n")
	net := make(Net)
	for _, node := range nodes {
		parts := strings.Split(node, " = (")
		if len(parts) != 2 {
			panic("invalid input")
		}
		source := parts[0]
		dest := strings.Split(parts[1], ", ")
		value := [2]string{dest[0], dest[1][:len(dest[1])-1]}
		net[source] = value
	}
	return Instructions{Path: path, Net: net}
}

func part1() {
	instructions := InstructionsFromString(in)
	cur := "AAA"
	curIdx := 0
	steps := 0
	for cur != "ZZZ" {
		direction := instructions.Path[curIdx]
		dest := instructions.Net[cur][direction]
		cur = dest
		curIdx = (curIdx + 1) % len(instructions.Path)
		steps++
	}
	fmt.Println(steps)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / GCD(lcm, nums[i])
	}
	return lcm
}

func part2() {
	instructions := InstructionsFromString(in)
	cur := make([]string, 0)
	for node := range instructions.Net {
		if node[2] == 'A' {
			cur = append(cur, node)
		}
	}
	results := make([]int, len(cur))
	for i := 0; i < len(cur); i++ {
		start := cur[i]
		curIdx := 0
		steps := 0
		for start[2] != 'Z' {
			direction := instructions.Path[curIdx]
			start = instructions.Net[start][direction]
			curIdx = (curIdx + 1) % len(instructions.Path)
			steps++
		}
		results[i] = steps
	}
	fmt.Println(LCM(results))
}

func main() {
	part1()
	part2()
}
