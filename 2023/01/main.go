package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	sumPartOne := 0
	sumPartTwo := 0
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range lines {
		var numsPartOne []int
		var numsPartTwo []int
		digitIndex := map[int]int{}
		for i, digit := range digits {
			firstIndex := strings.Index(line, digit)
			if firstIndex != -1 {
				digitIndex[firstIndex] = i + 1
			}
			lastIndex := strings.LastIndex(line, digit)
			if lastIndex != -1 && firstIndex != lastIndex {
				digitIndex[lastIndex] = i + 1
			}
		}
		for i, r := range line {
			if val, ok := digitIndex[i]; ok {
				numsPartTwo = append(numsPartTwo, val)
				continue
			}
			if number, err := strconv.Atoi(string(r)); err == nil {
				numsPartOne = append(numsPartOne, number)
				numsPartTwo = append(numsPartTwo, number)
			}
		}

		sumPartOne += numsPartOne[0]*10 + numsPartOne[len(numsPartOne)-1]
		sumPartTwo += numsPartTwo[0]*10 + numsPartTwo[len(numsPartTwo)-1]
	}
	fmt.Printf("Part 1: %d\n", sumPartOne)
	fmt.Printf("Part 2: %d\n", sumPartTwo)
}
