package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

var (
	rules          = map[string]string{}
	pairsFrequency = map[string]int{}
	letterQuantity = map[string]int{}
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	t := strings.Split(lines[0], "")

	lines = lines[2:]

	for _, line := range lines {
		sp := strings.Split(line, " -> ")
		rules[sp[0]] = sp[1]
	}

	start := time.Now()

	for i := 0; i < len(t)-1; i++ {
		pairsFrequency[t[i]+t[i+1]]++
	}

	// Part 1: for step := 1; step <= 10; step++ {
	for step := 1; step <= 40; step++ {
		newPf := map[string]int{}
		for k, v := range pairsFrequency {
			sp := strings.Split(k, "")
			newPf[sp[0]+rules[k]] += v
			newPf[rules[k]+sp[1]] += v

			pairsFrequency = newPf
		}
	}

	for k, v := range pairsFrequency {
		letterQuantity[string(k[0])] += v
	}
	// Add +1 for the last letter of the template
	letterQuantity[string(t[len(t)-1])] += 1

	var high = 0
	var low = math.MaxInt
	for _, v := range letterQuantity {
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
	}

	fmt.Println(high-low, time.Since(start))
}
