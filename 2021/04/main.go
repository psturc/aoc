package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type board struct {
	rows  [][]string
	rank  int
	score int
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	draws := strings.Split(lines[0], ",")

	boards := make([]board, 0)
	var globalRank int

	lines = lines[2:]
	var b board
	for _, line := range lines {
		s := strings.Fields(line)
		if len(s) == 5 {
			b.rows = append(b.rows, s)
		} else {
			boards = append(boards, b)
			b = board{}
		}
	}

	for _, draw := range draws {
		for bix, b := range boards {
			if b.rank != 0 {
				continue
			}
			for _, row := range b.rows {
				for j, n := range row {
					if n == draw {
						row[j] = "x"
					}
				}
			}
			if wins, score := checkScore(b, draw); wins {
				globalRank++
				if globalRank == 1 {
					fmt.Printf("Part 1: %d\n", score)
				}
				boards[bix].score = score
				boards[bix].rank = globalRank
			}
		}
	}

	for _, bd := range boards {
		if bd.rank == globalRank {
			fmt.Printf("Part 2: %d\n", bd.score)
		}
	}

}

func checkScore(b board, draw string) (bool, int) {
	var x, y = make([]int, 5), make([]int, 5)
	var score int
	for i, r := range b.rows {
		for j, n := range r {
			if n == "x" {
				x[i] += 1
				y[j] += 1
			} else {
				nInt, _ := strconv.Atoi(n)
				score += nInt
			}
		}
	}
	for i := range x {
		if x[i] == 5 || y[i] == 5 {
			d, _ := strconv.Atoi(draw)
			return true, score * d
		}
	}
	return false, 0
}
