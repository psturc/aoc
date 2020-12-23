package main

import (
	"bufio"
	"fmt"
	"os"
)

var seats = [][]string{}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()

		seats = append(seats, []string{})
		for _, v := range line {
			seats[lineNum] = append(seats[lineNum], string(v))
		}
		lineNum++
	}
	keepApplyingRulesUntilStabilized()
	fmt.Println("part 1:", countOcuppiedSeats())
	//fmt.Println("part 2:", )

}

func keepApplyingRulesUntilStabilized() {
	for {
		changed := applyRules()
		if !changed {
			break
		}
	}
}

func applyRules() (changed bool) {
	newState := make([][]string, len(seats))
	for y := range seats {
		for x := range seats[y] {
			switch seats[y][x] {
			case "L":
				if canOccupy(y, x) {
					changed = true
					newState[y] = append(newState[y], "#")
					continue
				}
			case "#":
				if canEmpty(y, x) {
					changed = true
					newState[y] = append(newState[y], "L")
					continue
				}
			}
			newState[y] = append(newState[y], seats[y][x])
		}
	}
	seats = newState
	return
}

func countOcuppiedSeats() (count int) {
	for i := range seats {
		for _, v := range seats[i] {
			if v == "#" {
				count++
			}
		}
	}
	return
}

func canOccupy(y, x int) bool {
	yRange, xRange := getYXRanges(y, x)
	for _, yv := range yRange {
		for _, xv := range xRange {
			if seats[yv][xv] == "#" {
				return false
			}
		}
	}
	return true
}

func canEmpty(y, x int) bool {
	var ocuppiedSeats int
	yRange, xRange := getYXRanges(y, x)
	for _, yv := range yRange {
		for _, xv := range xRange {
			if seats[yv][xv] == "#" && (yv != y || xv != x) {
				ocuppiedSeats++
			}
		}
	}
	if ocuppiedSeats >= 4 {
		return true
	}
	return false
}

func getYXRanges(y, x int) (yRange []int, xRange []int) {
	if y > 0 {
		yRange = append(yRange, y-1)
	}
	if x > 0 {
		xRange = append(xRange, x-1)
	}
	yRange = append(yRange, y)
	xRange = append(xRange, x)
	if y < len(seats)-1 {
		yRange = append(yRange, y+1)
	}
	if x < len(seats[y])-1 {
		xRange = append(xRange, x+1)
	}

	return yRange, xRange
}
