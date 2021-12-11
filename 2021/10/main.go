package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	var score int
	var acScores = []int{}
	var scoreSys = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var acScoreSys = map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
	var closingCh = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	for _, line := range lines {
		oCh := regexp.MustCompile(`[\[\(\{<]`)
		line := strings.Split(line, "")

		var lifo = []string{}
		var corrupted bool
		for _, ch := range line {
			if oCh.MatchString(ch) {
				lifo = append(lifo, ch)
			} else {
				if lifo[len(lifo)-1] != closingCh[ch] {
					score += scoreSys[ch]
					corrupted = true
					break
				} else {
					lifo = lifo[0 : len(lifo)-1]
				}
			}
		}
		if !corrupted {
			var acscore int
			for i := len(lifo) - 1; i >= 0; i-- {
				acscore = acscore*5 + acScoreSys[lifo[i]]
			}
			acScores = append(acScores, acscore)
		}
	}
	fmt.Printf("Part 1: %d\n", score)

	sort.Ints(acScores)
	fmt.Printf("Part 2: %d\n", acScores[len(acScores)/2])
}
