package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed in.txt
var data string

type numPos struct {
	start int
	end   int
	row   int
}

func isCoordFilled(row, col int, lines []string) bool {
	if row < 0 || col < 0 || row >= len(lines) || col >= len(lines[0]) {
		return false
	}
	return lines[row][col] != '.' && !unicode.IsDigit(rune(lines[row][col]))
}

func isPartNum(num numPos, lines []string) bool {
	for i := num.start - 1; i <= num.end+1; i++ {
		if isCoordFilled(num.row-1, i, lines) {
			return true
		}
		if isCoordFilled(num.row+1, i, lines) {
			return true
		}
	}
	if isCoordFilled(num.row, num.start-1, lines) {
		return true
	}
	if isCoordFilled(num.row, num.end+1, lines) {
		return true
	}
	return false
}

func getNums(lines []string) []numPos {
	nums := []numPos{}
	for row, line := range lines {
		inNum := false
		start := -1
		for col, r := range line {
			if !inNum && unicode.IsDigit(r) {
				inNum = true
				start = col
			} else if inNum && !unicode.IsDigit(r) {
				nums = append(nums, numPos{row: row, start: start, end: col - 1})
				inNum = false
				start = -1
			}
		}
		if inNum {
			nums = append(nums, numPos{row: row, start: start, end: len(line) - 1})
		}
	}
	return nums
}

func part1() {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	nums := getNums(lines)

	total := 0
	for _, n := range nums {
		if isPartNum(n, lines) {
			res, err := strconv.Atoi(lines[n.row][n.start : n.end+1])
			if err != nil {
				panic(err)
			}
			// fmt.Println(res)
			total += res
		}
	}
	fmt.Println(total)
}

type gearPos struct {
	row int
	col int
}

type gearStats struct {
	count int
	ratio int
}

func updateGears(row, col int, lines []string, gears map[gearPos]gearStats, val int) {
	if row < 0 || col < 0 || row >= len(lines) || col >= len(lines[0]) {
		return
	}
	if lines[row][col] != '*' {
		return
	}
	pos := gearPos{row: row, col: col}
	gearStats, found := gears[pos]
	gearStats.count += 1
	if !found {
		gearStats.ratio = val
	} else {
		gearStats.ratio *= val
	}
	gears[pos] = gearStats
}

func getGears(lines []string, nums []numPos) map[gearPos]gearStats {
	gears := map[gearPos]gearStats{}
	for _, num := range nums {
		val, err := strconv.Atoi(lines[num.row][num.start : num.end+1])
		if err != nil {
			panic(err)
		}
		for i := num.start - 1; i <= num.end+1; i++ {
			updateGears(num.row-1, i, lines, gears, val)
			updateGears(num.row+1, i, lines, gears, val)
		}
		updateGears(num.row, num.start-1, lines, gears, val)
		updateGears(num.row, num.end+1, lines, gears, val)
	}
	return gears
}

func part2() {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	nums := getNums(lines)
	gears := getGears(lines, nums)

	total := 0
	for _, stats := range gears {
		if stats.count == 2 {
			total += stats.ratio
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
