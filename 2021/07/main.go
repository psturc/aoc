package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	l := utils.FileLinesToSlice("input.txt")[0]

	h := utils.StrSliceToIntSlice(strings.Split(l, ","))

	sort.Ints(h)
	median := h[len(h)/2]

	var fuelm int
	for _, v := range h {
		fuelm += utils.AbsoluteValue(median - v)
	}

	dist := make([]int, 0)
	for i := h[0]; i < h[len(h)-1]; i++ {
		var sum int = 0
		for _, v := range h {
			sum += cost(utils.AbsoluteValue(v - i))
		}
		dist = append(dist, sum)
	}

	sort.Ints(dist)

	fmt.Printf("Part 1: %d\n", fuelm)
	fmt.Printf("Part 2: %d\n", dist[0])
}

func cost(i int) int {
	return i * (i + 1) / 2
}
