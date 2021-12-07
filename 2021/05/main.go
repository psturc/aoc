package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	diagram := map[string]int{}

	for _, l := range lines {
		r := regexp.MustCompile(`,| -> `)
		s := r.Split(l, -1)
		coords := make([]int, len(s))
		for i, v := range s {
			coords[i], _ = strconv.Atoi(v)
		}
		x1, y1, x2, y2 := coords[0], coords[1], coords[2], coords[3]

		if x1 == x2 {
			for {
				p := fmt.Sprintf("%d,%d", x1, y1)
				if _, ok := diagram[p]; !ok {
					diagram[p] = 1
				} else {
					diagram[p]++
				}
				if y1 == y2 {
					break
				} else if y1 > y2 {
					y1--
				} else {
					y1++
				}
			}
		} else if y1 == y2 {
			for {
				p := fmt.Sprintf("%d,%d", x1, y1)
				if _, ok := diagram[p]; !ok {
					diagram[p] = 1
				} else {
					diagram[p]++
				}
				if x1 == x2 {
					break
				} else if x1 > x2 {
					x1--
				} else {
					x1++
				}
			}
			// Part 2
		} else if utils.AbsoluteValue(x1-x2) == utils.AbsoluteValue(y1-y2) {
			for {
				p := fmt.Sprintf("%d,%d", x1, y1)
				if _, ok := diagram[p]; !ok {
					diagram[p] = 1
				} else {
					diagram[p]++
				}
				if x1 == x2 && y1 == y2 {
					break
				}
				if x1 > x2 {
					x1--
				} else if x1 < x2 {
					x1++
				}
				if y1 > y2 {
					y1--
				} else if y1 < y2 {
					y1++
				}
			}
		}
	}
	var res int
	for _, v := range diagram {
		if v > 1 {
			res++
		}
	}
	//fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", res)
}
