package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	var area = [][]int{}

	var sum int
	basinSizes := make([]int, 0)
	for _, line := range lines {
		r := utils.StrSliceToIntSlice(strings.Split(line, ""))
		area = append(area, r)
	}

	for y := range area {
		for x := range area[0] {
			x1, x2, y1, y2 := x-1, x+1, y-1, y+1

			if x1 >= 0 && area[y][x] >= area[y][x1] {
				continue
			}
			if x2 < len(area[0]) && area[y][x] >= area[y][x2] {
				continue
			}
			if y1 >= 0 && area[y][x] >= area[y1][x] {
				continue
			}
			if y2 < len(area) && area[y][x] >= area[y2][x] {
				continue
			}

			sum += area[y][x] + 1
			vis := findBasin(area, y, x, map[string]bool{})
			basinSizes = append(basinSizes, len(vis))
		}
	}

	fmt.Printf("Part 1: %d\n", sum)

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Printf("Part 2: %d\n", basinSizes[0]*basinSizes[1]*basinSizes[2])

}

func findBasin(area [][]int, y, x int, vis map[string]bool) map[string]bool {
	coord := fmt.Sprintf("%d,%d", y, x)
	if _, ok := vis[coord]; ok {
		return vis
	} else {
		vis[coord] = true
	}

	x1, x2, y1, y2 := x-1, x+1, y-1, y+1
	if x1 >= 0 && area[y][x1] != 9 {
		vis = mergeMaps(vis, findBasin(area, y, x1, vis))
	}
	if x2 < len(area[0]) && area[y][x2] != 9 {
		vis = mergeMaps(vis, findBasin(area, y, x2, vis))
	}
	if y1 >= 0 && area[y1][x] != 9 {
		vis = mergeMaps(vis, findBasin(area, y1, x, vis))
	}
	if y2 < len(area) && area[y2][x] != 9 {
		vis = mergeMaps(vis, findBasin(area, y2, x, vis))
	}

	return vis
}

func mergeMaps(old, new map[string]bool) map[string]bool {
	for k, v := range new {
		old[k] = v
	}
	return old
}
