package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func readLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(content)
	text = strings.TrimSpace(text)
	lines := strings.Split(text, "\n")
	return lines
}

func firstAndLastDigits(s string) int {
	first := -1
	last := -1
	for _, char := range s {
		if unicode.IsDigit(char) {
			digit := int(char - '0')
			if first == -1 {
				first = digit
			}
			last = digit
		}
	}
	return first*10 + last
}

func part1(filename string) {
	lines := readLines(filename)
	total := 0
	for _, line := range lines {
		val := firstAndLastDigits(line)
		total += val
	}
	fmt.Println(total)
}

func mapKeys(m map[string]int) []string {
	result := make([]string, 0)
	for k := range m {
		result = append(result, k)
	}
	return result
}

func part2(filename string) {
	var words = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	for i := 1; i < 10; i++ {
		words[strconv.Itoa(i)] = i
	}
	pattern := "("
	pattern += strings.Join(mapKeys(words), "|")
	pattern += ")"
	r := regexp.MustCompile(pattern)
	rLast := regexp.MustCompile(".*" + pattern)

	lines := readLines(filename)
	total := 0
	for _, line := range lines {
		first := r.FindString(line)
		if first == "" {
			panic("didn't found first match in string" + line)
		}
		results := rLast.FindStringSubmatch(line)
		last := results[len(results)-1]
		if last == "" {
			panic("didn't found last match in string" + line)
		}

		firstVal, ok := words[first]
		if !ok {
			panic("unknown pattern for first: " + first)
		}
		lastVal, ok := words[last]
		if !ok {
			panic("unknown pattern for last: <" + first + "> in string " + line)
		}
		total += firstVal*10 + lastVal
	}
	fmt.Println(total)
}

func main() {
	part1("d1/in.txt")
	part2("d1/in.txt")
}
