package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var setup map[int][]string
var setup2 map[int][]string

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	setup = doSetup(lines)
	setup2 = doSetup(lines)
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, l := range lines {
		matches := re.FindStringSubmatch(l)
		if len(matches) > 0 {
			howMuch, _ := strconv.Atoi(matches[len(matches)-3])
			from, _ := strconv.Atoi(matches[len(matches)-2])
			to, _ := strconv.Atoi(matches[len(matches)-1])

			move(howMuch, from, to)
		}

	}

	var part1, part2 string
	for i := 1; i <= len(setup); i++ {
		part1 += setup[i][len(setup[i])-1]
		part2 += setup2[i][len(setup2[i])-1]
	}

	fmt.Println("1:", part1)
	fmt.Println("2:", part2)
}

func doSetup(lines []string) map[int][]string {
	setup := map[int][]string{}

	for _, line := range lines {
		if !strings.Contains(line, "[") {
			break
		}
		p := 0
		for i := 1; p < len(line); i++ {
			if line[p+1] != 32 {
				setup[i] = append([]string{string(line[p+1])}, (setup[i])...)
			}
			p += 4
		}
	}
	return setup
}

func move(howMuch, from, to int) {
	// part1 solution
	part1 := []string{}
	for i := 1; i <= howMuch; i++ {
		part1 = append(part1, setup[from][len(setup[from])-i])
	}
	setup[from] = setup[from][0 : len(setup[from])-howMuch]
	setup[to] = append(setup[to], part1...)

	// part2 solution
	part2 := setup2[from][len(setup2[from])-howMuch:]
	setup2[from] = setup2[from][0 : len(setup2[from])-howMuch]
	setup2[to] = append(setup2[to], part2...)
}
