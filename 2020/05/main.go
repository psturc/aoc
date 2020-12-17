package main

import (
	"bufio"
	"fmt"
	"os"
)

var highestSeatID int
var seats [128][8]int

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		seatID := getSeatID(line)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Println("part 1:", highestSeatID)
	fmt.Println("part 2:", getMySeatID())
}

func getSeatID(line string) int {
	rowRange := []int{0, 127}
	columnRange := []int{0, 7}

	for _, v := range line {
		switch string(v) {
		case "F":
			rowRange[1] = rowRange[0] + (rowRange[1]-rowRange[0])/2
		case "B":
			rowRange[0] = rowRange[1] - (rowRange[1]-rowRange[0])/2
		case "L":
			columnRange[1] = columnRange[0] + (columnRange[1]-columnRange[0])/2
		case "R":
			columnRange[0] = columnRange[1] - (columnRange[1]-columnRange[0])/2
		}
	}
	seats[rowRange[0]][columnRange[0]] = 1
	return rowRange[0]*8 + columnRange[0]
}

func getMySeatID() int {
	var seatExists bool
	for i, row := range seats {
		for j, column := range row {
			if column == 0 && seatExists {
				return i*8 + j
			}
			if !seatExists && column == 1 {
				seatExists = true
			}
		}
	}
	return 0
}
