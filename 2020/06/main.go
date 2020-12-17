package main

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	size       int
	yesAnswers map[string]int
}

func newGroup() group {
	return group{size: 0, yesAnswers: make(map[string]int)}
}

// https://adventofcode.com/2020/day/6
func main() {

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var groupID int
	groups := make([]group, 1)
	groups[0] = newGroup()

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 {
			groups[groupID].size++
			for _, v := range line {
				groups[groupID].yesAnswers[string(v)]++
			}
		} else {
			groups = append(groups, newGroup())
			groupID++
		}
	}

	allYesAnswersCount, groupYesAnswersCount := countAnswers(groups)

	fmt.Println("part 1:", allYesAnswersCount)
	fmt.Println("part 2:", groupYesAnswersCount)
}

func countAnswers(groups []group) (allYeses int, groupYeses int) {
	for _, g := range groups {
		allYeses += len(g.yesAnswers)
		for _, c := range g.yesAnswers {
			if g.size == c {
				groupYeses++
			}
		}
	}
	return
}
