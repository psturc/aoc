package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func main() {

	lines := utils.FileLinesToSlice("input.txt")

	topthree := map[int]int{1: 0, 2: 0, 3: 0}
	sum := 0

	for _, l := range lines {
		if l == "" {
			if sum >= topthree[1] {
				topthree[3] = topthree[2]
				topthree[2] = topthree[1]
				topthree[1] = sum
			} else if sum >= topthree[2] {
				topthree[3] = topthree[2]
				topthree[2] = sum
			} else if sum >= topthree[3] {
				topthree[3] = sum
			}
			sum = 0
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		sum += n
	}

	fmt.Println("1:", topthree[1])
	fmt.Println("2:", topthree[1]+topthree[2]+topthree[3])
}
