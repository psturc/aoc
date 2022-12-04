package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	pairOverlaps := 0
	partialOverlaps := 0

	for _, l := range lines {
		sp := strings.FieldsFunc(l, func(r rune) bool { return r == '-' || r == ',' })

		fstr, fm := generateStringAndMapSequences(sp[0], sp[1])
		sstr, sm := generateStringAndMapSequences(sp[2], sp[3])

		for k := range fm {
			if sm[k] {
				partialOverlaps++
				if strings.Contains(fstr, sstr) || strings.Contains(sstr, fstr) {
					pairOverlaps++
				}
				break
			}
		}
	}
	fmt.Println("1:", pairOverlaps)
	fmt.Println("2:", partialOverlaps)
}

func generateStringAndMapSequences(start, end string) (seq string, m map[int]bool) {
	seq = ","
	m = map[int]bool{}
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	for i := startInt; i <= endInt; i++ {
		seq += strconv.Itoa(i) + ","
		m[i] = true
	}
	return seq, m
}
