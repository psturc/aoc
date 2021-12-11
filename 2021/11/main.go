package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var flashCount int

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	var area = [][]int{}

	for _, line := range lines {
		area = append(area, utils.StrSliceToIntSlice(strings.Split(line, "")))
	}

	var i int = 1
	for {

		for y := range area {
			for x := range area[0] {
				if area[y][x] == 9 {
					area[y][x] = 0
				} else {
					area[y][x] = area[y][x] + 1
				}

			}
		}

		var visFlash = map[string]bool{}
		for y := range area {
			for x := range area[0] {
				if area[y][x] == 0 {
					flash(area, y, x, visFlash)
				}
			}
		}

		if i == 100 {
			fmt.Printf("Part 1: %d\n", flashCount)
		}
		if simultaneousFlash(area) {
			fmt.Printf("Part 2: %d\n", i)
			break
		}
		i++
	}

}

func flash(area [][]int, y, x int, visFlash map[string]bool) map[string]bool {
	coord := fmt.Sprintf("%d,%d", y, x)

	if area[y][x] == 0 {
		if _, ok := visFlash[coord]; ok {
			return visFlash
		}
	} else if area[y][x] == 9 {
		area[y][x] = 0
	} else {
		area[y][x]++
		return visFlash
	}
	flashCount++
	visFlash[coord] = true
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if y+i >= 0 && y+i < len(area) && x+j >= 0 && x+j < len(area[0]) {
				flash(area, y+i, x+j, visFlash)
			}
		}
	}
	return visFlash
}

func simultaneousFlash(area [][]int) bool {
	for _, row := range area {
		for _, v := range row {
			if v != 0 {
				return false
			}
		}
	}
	return true
}
