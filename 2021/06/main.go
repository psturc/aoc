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
		var v int
		for j := 8; j >= 0; j-- {
			if j == 8 {
				v = gmap[8]
				gmap[8] = gmap[0]
			} else if j > 0 {
				vv := v
				v = gmap[j]
				gmap[j] = vv
			} else {
				gmap[6] += gmap[0]
				gmap[0] = v
			}
		}
	}

	var count int
	for _, v := range gmap {
		count += v
	}
	fmt.Println(count)
}
