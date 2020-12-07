package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://adventofcode.com/2020/day/3
func main() {
	var treesMet int
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileLines := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	treesMet = treesHitsWithSlope(fileLines, 3, 1)
	fmt.Println("part one: ", treesMet)

	treesMet *= treesHitsWithSlope(fileLines, 1, 1)
	treesMet *= treesHitsWithSlope(fileLines, 5, 1)
	treesMet *= treesHitsWithSlope(fileLines, 7, 1)
	treesMet *= treesHitsWithSlope(fileLines, 1, 2)

	fmt.Println("part two: ", treesMet)
}

func treesHitsWithSlope(fileLines []string, xStep int, yStep int) (treesHits int) {
	xPosition := 0
	yPosition := 0

	lineLength := len(fileLines[0])

	for i, line := range fileLines {

		if i%yStep != 0 {
			continue
		}

		if string(line[xPosition%lineLength]) == "#" {
			treesHits++
		}

		xPosition += xStep
		yPosition += yStep
	}
	return
}
