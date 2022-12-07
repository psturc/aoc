package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.FileLinesToSlice("input.txt")[0]
	chars := []byte{}
	var partOneSolved bool

	for i := range input {
		for j := range chars {
			if chars[j] == input[i] {
				chars = chars[j+1:]
				break
			}
		}
		chars = append(chars, input[i])

		if !partOneSolved && len(chars) == 4 {
			fmt.Println("1: ", i+1)
			partOneSolved = true
		}
		if len(chars) == 14 {
			fmt.Println("2: ", i+1)
			break
		}
	}
}
