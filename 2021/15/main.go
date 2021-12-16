package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	area        = [][]int{}
	hugeArea    = [][]int{}
	areaFactory = map[int][][]int{}
	toVisit     = map[string]int{}
	visualize   = [][]string{}
	vList       = map[string]*vertex{}
	up          = []int{-1, 0}
	right       = []int{0, 1}
	down        = []int{1, 0}
	left        = []int{0, -1}
	goal        int
)

type vertex struct {
	value     int
	neighbors []string
	dist      int
	visited   bool
	prev      string
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	for _, line := range lines {
		lineInt := utils.StrSliceToIntSlice(strings.Split(line, ""))
		area = append(area, lineInt)
	}

	areaFactory[0] = area
	for i := 1; i < 9; i++ {
		areaFactory[i] = increaseArea(areaFactory[i-1])
	}

	step := 0
	multiplier := 1
	// Part 2: multiplier := 5
	for i := 0; i < multiplier; i++ {
		hugeArea = append(hugeArea, areaFactory[i]...)
		for j := i + 1; j < multiplier+i; j++ {
			newArea := areaFactory[j]
			for x := step; x < step+len(area); x++ {
				hugeArea[x] = append(hugeArea[x], newArea[x-step]...)
			}
		}
		step += len(area)
	}

	area = hugeArea

	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area); x++ {
			coord := fmt.Sprintf("%d,%d", y, x)

			var neighbors = []string{}
			for _, v := range [][]int{up, right, down, left} {
				if y+v[0] >= 0 && x+v[1] >= 0 && y+v[0] < len(area) && x+v[1] < len(area) {
					neighbors = append(neighbors, fmt.Sprintf("%d,%d", y+v[0], x+v[1]))
				}
			}
			vList[coord] = &vertex{value: area[y][x], neighbors: neighbors, dist: math.MaxInt}

		}
	}
	goal = len(area) - 1
	start := time.Now()

	coord := "0,0"
	coordGoal := fmt.Sprintf("%d,%d", goal, goal)
	vList[coord].dist = 0

	for {

		vList[coord].visited = true
		vx := vList[coord]

		if coord == coordGoal {
			fmt.Println(vx.dist)
			break
		}

		for _, n := range vx.neighbors {
			nb := vList[n]
			newDist := vx.dist + vList[n].value
			if !nb.visited && newDist < nb.dist {
				vList[n].dist = newDist
				vList[n].prev = coord
				toVisit[n] = newDist
			}
		}

		coord = getMinDist()
		delete(toVisit, coord)

	}

	fmt.Println(time.Since(start))

	for y := 0; y < len(area); y++ {
		visualize = append(visualize, []string{})
		for x := 0; x < len(area); x++ {
			visualize[y] = append(visualize[y], "-")
		}
	}

	toPrint := fmt.Sprintf("%d,%d", goal, goal)
	for {
		sp := strings.Split(toPrint, ",")
		y, _ := strconv.Atoi(sp[0])
		x, _ := strconv.Atoi(sp[1])

		visualize[y][x] = "O"
		toPrint = vList[toPrint].prev
		if toPrint == "" {
			break
		}
	}

}

func getMinDist() string {
	var id string
	var minDist int = math.MaxInt
	for k, v := range toVisit {
		if v < minDist {
			minDist = v
			id = k
		}
	}
	return id
}

func increaseArea(a [][]int) [][]int {
	var newArea = [][]int{}
	for i, r := range a {
		newArea = append(newArea, []int{})
		for _, v := range r {
			var n = 1
			if v < 9 {
				n = v + 1
			}
			newArea[i] = append(newArea[i], n)
		}
	}
	return newArea
}
