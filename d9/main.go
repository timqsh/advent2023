package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed in.txt
var in string

func main() {
	lines := strings.Split(strings.TrimSpace(in), "\n")
	total := 0
	totalBefore := 0
	for _, line := range lines {
		nums := []int{}
		for _, num := range strings.Split(line, " ") {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums = append(nums, val)
		}
		steps := [][]int{}
		steps = append(steps, nums)
		cur := nums
		for {
			allZero := true
			for _, num := range cur {
				if num != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				break
			}
			next := []int{}
			for i := 0; i < len(cur)-1; i++ {
				diff := cur[i+1] - cur[i]
				next = append(next, diff)
			}
			steps = append(steps, next)
			cur = next
		}
		valsBefore := []int{0}
		for i := len(steps) - 2; i >= 0; i-- {
			val := steps[i][len(steps[i])-1] + steps[i+1][len(steps[i+1])-1]
			steps[i] = append(steps[i], val)
			valBefore := steps[i][0] - valsBefore[len(valsBefore)-1]
			valsBefore = append(valsBefore, valBefore)
		}
		predict := steps[0][len(steps[0])-1]
		before := valsBefore[len(valsBefore)-1]
		total += predict
		totalBefore += before
	}
	fmt.Println(total)
	fmt.Println(totalBefore)
}
