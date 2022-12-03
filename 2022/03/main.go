package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	prioritySum := 0

	for _, l := range lines {
		fh := map[byte]bool{}
		sh := map[byte]bool{}

		for i := range l {
			if i < len(l)/2 {
				fh[l[i]] = true
				continue
			}
			sh[l[i]] = true
			if _, ok := fh[l[i]]; ok {
				// A-Z
				if l[i] < 97 {
					prioritySum += int(l[i] - 38)
				} else {
					// a-z
					prioritySum += int(l[i] - 96)
				}
				break
			}
		}
	}

	fmt.Println("1:", prioritySum)

}
