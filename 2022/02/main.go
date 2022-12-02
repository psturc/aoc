package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

/*
0 0 = draw
0 1 = loss
0 2 = win

1 0 = win
1 1 = draw
1 2 = loss

2 0 = loss
2 1 = win
2 2 = draw
*/

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	score := 0
	score2 := 0

	op := map[string]int{"A": 0, "B": 1, "C": 2}
	me := map[string]int{"X": 0, "Y": 1, "Z": 2}

	for _, l := range lines {
		i, j := strings.Fields(l)[0], strings.Fields(l)[1]
		// 1
		if op[i] == me[j] {
			score += 3
		} else if (op[i]+1)%3 == me[j] {
			score += 6
		}
		score += me[j] + 1

		// 2
		switch me[j] {
		case 0: // lose
			score2 += 0 + (op[i]+2)%3 + 1
		case 1: // draw
			score2 += 3 + op[i] + 1
		case 2: // win
			score2 += 6 + (op[i]+1)%3 + 1
		}

	}

	fmt.Println("1:", score)
	fmt.Println("2:", score2)
}
