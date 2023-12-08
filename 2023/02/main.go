package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

const maxRed int = 12
const maxGreen int = 13
const maxBlue int = 14

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	var idSum, powerSum int

	for i, line := range lines {
		sp := strings.Split(line, ":")
		sets := strings.Split(sp[1], ";")
		validGame := true

		var maxRedPlayed, maxGreenPlayed, maxBluePlayed int

		for _, set := range sets {
			var totalRed, totalGreen, totalBlue int

			for _, amount := range strings.Split(set, ", ") {
				numAndColor := strings.Fields(amount)
				numString, color := numAndColor[0], numAndColor[1]

				num, _ := strconv.Atoi(numString)

				switch color {
				case "red":
					totalRed += num
					if num > maxRedPlayed {
						maxRedPlayed = num
					}
				case "green":
					totalGreen += num
					if num > maxGreenPlayed {
						maxGreenPlayed = num
					}
				case "blue":
					totalBlue += num
					if num > maxBluePlayed {
						maxBluePlayed = num
					}
				}
			}

			if totalRed > maxRed || totalGreen > maxGreen || totalBlue > maxBlue {
				validGame = false
			}

		}
		if validGame {
			idSum += i + 1
		}
		powerSum += maxRedPlayed * maxGreenPlayed * maxBluePlayed
	}

	fmt.Printf("Part 1: %d\n", idSum)
	fmt.Printf("Part 2: %d\n", powerSum)
}
