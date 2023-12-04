package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed in.txt
var data string

type Num struct {
	start int
	end   int
	row   int
	val   int
}

type Pos struct {
	row int
	col int
}

func (n Num) neighbors(lines []string) []Pos {
	allNeighbors := []Pos{}
	for i := n.start - 1; i <= n.end+1; i++ {
		allNeighbors = append(allNeighbors, Pos{n.row - 1, i})
		allNeighbors = append(allNeighbors, Pos{n.row + 1, i})
	}
	allNeighbors = append(allNeighbors, Pos{n.row, n.start - 1})
	allNeighbors = append(allNeighbors, Pos{n.row, n.end + 1})

	validNeighbors := []Pos{}
	for _, n := range allNeighbors {
		if 0 <= n.row && n.row < len(lines) && 0 <= n.col && n.col < len(lines[0]) {
			validNeighbors = append(validNeighbors, n)
		}
	}
	return validNeighbors
}

func getNums(lines []string) []Num {
	nums := []Num{}
	r := regexp.MustCompile(`\d+`)
	for row, line := range lines {
		matches := r.FindAllStringIndex(line, -1)
		for _, match := range matches {
			val, err := strconv.Atoi(lines[row][match[0]:match[1]])
			if err != nil {
				panic(err)
			}
			pos := Num{row: row, start: match[0], end: match[1] - 1, val: val}
			nums = append(nums, pos)
		}
	}
	return nums
}

func part1() {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	nums := getNums(lines)

	total := 0
	for _, n := range nums {
		for _, neigh := range n.neighbors(lines) {
			if lines[neigh.col][neigh.row] != '.' {
				total += n.val
				break
			}
		}
	}
	fmt.Println(total)
}

func getGearParts(lines []string, nums []Num) map[Pos][]int {
	gearParts := map[Pos][]int{}
	for _, num := range nums {
		for _, n := range num.neighbors(lines) {
			if lines[n.row][n.col] == '*' {
				gearParts[n] = append(gearParts[n], num.val)
			}
		}
	}
	return gearParts
}

func part2() {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	nums := getNums(lines)
	gearParts := getGearParts(lines, nums)

	total := 0
	for _, parts := range gearParts {
		if len(parts) == 2 {
			total += parts[0] * parts[1]
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
