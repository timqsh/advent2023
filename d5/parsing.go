package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AlmanacFromString(s string) Almanac {
	blocks := strings.Split(s, "\n\n")
	if len(blocks) != 8 {
		panic(fmt.Sprintf("Parsed %v blocks, expected 8\n", len(blocks)))
	}
	var mappings []Mapping
	for _, block := range blocks[1:] {
		mappings = append(mappings, MappingFromBlocks(block))
	}
	return Almanac{seeds: SeedsFromString(blocks[0]), mappings: mappings}
}

func SeedsFromString(s string) Seeds {
	result := Seeds{}
	s = strings.TrimPrefix(s, "seeds: ")
	for _, part := range strings.Split(s, " ") {
		val, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("Can't parse seed number <%v> from string %v", part, s))
		}
		result = append(result, val)
	}
	return result
}

func MappingFromBlocks(block string) Mapping {
	result := Mapping{}
	lines := strings.Split(block, "\n")
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		r := RangeFromString(line)
		result = append(result, r)
	}
	return result
}

func RangeFromString(s string) Range {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		panic(fmt.Sprintf("Got %v parts in range, expected 3 in line %v\n", len(parts), s))
	}
	var nums []int
	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("Can't parse range value <%v> in line %v\n", part, s))
		}
		nums = append(nums, val)
	}
	return Range{destStart: nums[0], sourceStart: nums[1], Len: nums[2]}
}
