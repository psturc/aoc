package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	prioritySum := 0
	threeRowsPrioritySum := 0

	group := map[int]map[byte]bool{
		0: {},
		1: {},
		2: {},
	}

	for i, l := range lines {
		fh := map[byte]bool{}
		lineSolved := false

		for j := range l {
			group[i%3][l[j]] = true
			if j < len(l)/2 {
				fh[l[j]] = true
				continue
			}
			if _, ok := fh[l[j]]; ok && !lineSolved {
				prioritySum += getPriorityValue(l[j])
				lineSolved = true
			}

		}

		if (i+1)%3 == 0 {
			for k := range group[0] {
				if group[1][k] && group[2][k] {
					threeRowsPrioritySum += getPriorityValue(k)
					group = map[int]map[byte]bool{
						0: {},
						1: {},
						2: {},
					}
					break
				}
			}
		}
	}

	fmt.Println("1:", prioritySum)
	fmt.Println("2:", threeRowsPrioritySum)
}

func getPriorityValue(b byte) int {
	if b < 97 {
		return int(b - 38)
	} else {
		// a-z
		return int(b - 96)
	}
}
