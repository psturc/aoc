package main

import (
	"fmt"
	"strconv"

	"aoc/utils"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var gamma, epsilon string
	counter := make([]int, len(lines[0]))

	for _, line := range lines {
		for i := range line {
			// string(49) == "1"
			if line[i] == 49 {
				counter[i]++
			}
		}
	}

	for _, c := range counter {
		if c > len(lines)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaDec * epsilonDec)
}

func part2(lines []string) int {
	var o2, co2 string

	o2 = getRating(lines, "o2")
	co2 = getRating(lines, "co2")

	o2Dec, _ := strconv.ParseInt(o2, 2, 64)
	co2Dec, _ := strconv.ParseInt(co2, 2, 64)

	return int(o2Dec * co2Dec)
}

func getRating(lines []string, ratingType string) string {
	var candidates = []string{}
	candidates = lines[:]

	for {
		for i := range lines[0] {
			zeros := make([]string, 0)
			ones := make([]string, 0)

			for _, line := range candidates {
				// string(49) == "1"
				if line[i] == 49 {
					ones = append(ones, line)
				} else {
					zeros = append(zeros, line)
				}
			}

			if ratingType == "o2" {
				if len(ones) >= len(zeros) {
					candidates = ones[:]
				} else {
					candidates = zeros[:]
				}
			} else {
				if len(zeros) <= len(ones) {
					candidates = zeros[:]
				} else {
					candidates = ones[:]
				}
			}

			if len(candidates) == 1 {
				return candidates[0]
			}
		}
	}
}
