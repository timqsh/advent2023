package main

import (
	_ "embed"
	"fmt"
	"math"
	"sync"
)

//go:embed in.txt
var input string

type Range struct {
	destStart   int
	sourceStart int
	Len         int
}

type Mapping []Range

func (m Mapping) Convert(source int) int {
	result := source
	for _, r := range m {
		if r.sourceStart <= source && source < r.sourceStart+r.Len {
			offset := source - r.sourceStart
			result = r.destStart + offset
			break
		}
	}
	return result
}

type Seeds []int

type Almanac struct {
	seeds    Seeds
	mappings []Mapping
}

func (a Almanac) SeedToLocation(seed int) int {
	current := seed
	for _, m := range a.mappings {
		current = m.Convert(current)
	}
	return current
}

func main() {
	alm := AlmanacFromString(input)
	part1 := math.MaxInt
	for _, seed := range alm.seeds {
		location := alm.SeedToLocation(seed)
		if location < part1 {
			part1 = location
		}
	}
	fmt.Println(part1)

	var results = make(chan int, 1000)
	wg := sync.WaitGroup{}
	for i := 0; i < len(alm.seeds); i += 2 {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			curMin := math.MaxInt
			start := alm.seeds[i]
			count := alm.seeds[i+1]
			for j := start; j < start+count; j++ {
				location := alm.SeedToLocation(j)
				if location < curMin {
					curMin = location
				}
			}
			results <- curMin
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	part2 := math.MaxInt
	for val := range results {
		if val < part2 {
			part2 = val
		}
	}
	// 2:54.71 without parallel
	// 55.6 sec with parallel
	fmt.Println(part2)
}
