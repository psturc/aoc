package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	g := utils.StrSliceToIntSlice(strings.Split(lines[0], ","))
	gmap := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	for _, v := range g {
		gmap[v]++
	}

	// Part 1: for i := 0; i < 80; i++ {
	for i := 0; i < 256; i++ {
		newMap := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
		for j := 8; j >= 0; j-- {
			if j > 0 {
				newMap[j-1] = gmap[j]
			} else {
				newMap[6] += gmap[j]
				newMap[8] += gmap[j]
			}
		}
		gmap = newMap
	}

	var count int
	for _, v := range gmap {
		count += v
	}
	fmt.Println(count)
}
