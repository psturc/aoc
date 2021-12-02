package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	x, depth := 0, 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		sl := strings.Split(lines[len(lines)-1], " ")

		switch sl[0] {
		case "forward":
			v, _ := strconv.Atoi(sl[1])
			x += v
		case "down":
			v, _ := strconv.Atoi(sl[1])
			depth += v
		case "up":
			v, _ := strconv.Atoi(sl[1])
			depth -= v
		}
	}

	fmt.Printf("Part 1: %d\n", x*depth)

	x, depth = 0, 0
	aim := 0

	for _, line := range lines {

		sl := strings.Split(line, " ")

		switch sl[0] {
		case "forward":
			v, _ := strconv.Atoi(sl[1])
			x += v
			depth += aim * v
		case "down":
			v, _ := strconv.Atoi(sl[1])
			aim += v
		case "up":
			v, _ := strconv.Atoi(sl[1])
			aim -= v
		}
	}

	fmt.Printf("Part 2: %d\n", x*depth)

}
