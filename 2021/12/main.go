package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var (
	foundPaths = map[string]bool{}
	combs      = map[string][]string{}
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	for _, line := range lines {
		ls := strings.Split(line, "-")
		l, r := ls[0], ls[1]

		mapCombinations(combs, l, r)
		mapCombinations(combs, r, l)
	}

	search("start", []string{})
	fmt.Printf("Part 1: %d\n", len(foundPaths))

	foundPaths = map[string]bool{}
	searchV2("start", []string{})
	fmt.Printf("Part 2: %d\n", len(foundPaths))
}

func mapCombinations(m map[string][]string, l, r string) {
	if l != "end" && r != "start" {
		if _, ok := m[l]; !ok {
			m[l] = []string{r}
		} else {
			if !utils.Contains(m[l], r) {
				m[l] = append(m[l], r)
			}
		}
	}
}

func search(p string, path []string) {
	path = append(path, p)

	for _, v := range combs[p] {
		if utils.IsLowercase(v) && utils.Contains(path, v) {
			continue
		}
		if v == "end" {
			if _, ok := foundPaths[strings.Join(path, ",")]; !ok {
				foundPaths[strings.Join(append(path, "end"), ",")] = true
			}
		}
		search(v, path)
	}

}

func searchV2(p string, path []string) {
	path = append(path, p)

	for _, v := range combs[p] {
		if utils.IsLowercase(v) && utils.Contains(path, v) {
			if visitedSmallCaveTwice(path) {
				continue
			}

		}
		if v == "end" {
			if _, ok := foundPaths[strings.Join(path, ",")]; !ok {
				foundPaths[strings.Join(append(path, "end"), ",")] = true
			}
		}
		searchV2(v, path)
	}

}

func visitedSmallCaveTwice(slc []string) bool {
	var visited = map[string]bool{}
	for _, c := range slc {
		if _, ok := visited[c]; ok && utils.IsLowercase(c) {
			return true
		} else {
			visited[c] = true
		}
	}
	return false
}
