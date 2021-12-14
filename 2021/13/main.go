package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	var m = map[int][]int{}
	var rules = [][]int{}

	for _, line := range lines {
		if strings.Contains(line, ",") {
			sp := strings.Split(line, ",")
			spInt := utils.StrSliceToIntSlice(sp)
			m[spInt[0]] = append(m[spInt[0]], []int{spInt[1]}...)

		}
		if strings.Contains(line, "=") {
			sp := strings.Fields(line)
			sp = strings.Split(sp[2], "=")
			v, _ := strconv.Atoi(sp[1])

			if sp[0] == "x" {
				rules = append(rules, []int{v, 0})
			} else {
				rules = append(rules, []int{0, v})
			}

		}
	}

	for i, rule := range rules {
		var dotsCount int
		for k, v := range m {
			if rule[0] == 0 {
				var ys = []int{}
				for _, y := range v {
					if y < rule[1] && !utils.ContainsInt(ys, y) {
						ys = append(ys, y)
					} else if y > rule[1] && !utils.ContainsInt(ys, 2*rule[1]-y) {
						ys = append(ys, 2*rule[1]-y)
					}
				}
				m[k] = ys
			} else {
				if k > rule[0] {
					newX := 2*rule[0] - k
					if _, ok := m[newX]; ok {
						for _, y := range v {
							if !utils.ContainsInt(m[newX], y) {
								m[newX] = append(m[newX], y)
							}
						}
					} else {
						m[newX] = v
					}
					delete(m, k)
				}
			}
		}
		for _, v := range m {
			dotsCount += len(v)
			sort.Ints(v)
		}
		if i == 0 {
			fmt.Printf("Part 1: %d\n", dotsCount)
		}
	}
	fmt.Printf("Part 2:\n\n")

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if _, ok := m[x]; ok {
				if utils.ContainsInt(m[x], y) {
					fmt.Print(" # ")
				} else {
					fmt.Print("   ")
				}
			}
		}
		fmt.Print("\n")
	}

}
