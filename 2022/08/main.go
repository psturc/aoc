package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var grid = [][]int{}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	part1 := 0
	part2 := 0

	for _, line := range lines {
		grid = append(grid, utils.StrSliceToIntSlice(strings.Split(line, "")))
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			var visible bool = false
			var scenicScore int = 1
			if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
				part1++
				continue
			}
			scenicScore *= inspectTree(x, y, 0, -1, &visible)
			scenicScore *= inspectTree(x, y, 0, 1, &visible)
			scenicScore *= inspectTree(x, y, -1, 0, &visible)
			scenicScore *= inspectTree(x, y, 1, 0, &visible)
			if visible {
				part1++
			}
			if scenicScore > part2 {
				part2 = scenicScore
			}
		}
	}

	fmt.Println("1:", part1)
	fmt.Println("2:", part2)
}

func inspectTree(x, y, dx, dy int, visible *bool) int {
	treeHeight := grid[y][x]
	tempScore := 0

	for i, j := x+dx, y+dy; i >= 0 && j >= 0 && i < len(grid[y]) && j < len(grid); i, j = i+dx, j+dy {
		tempScore++
		if treeHeight <= grid[j][i] {
			break
		}
		if i+dx == len(grid[y]) || j+dy == len(grid) || i+dx < 0 || j+dy < 0 {
			*visible = true
		}
	}

	return tempScore
}
