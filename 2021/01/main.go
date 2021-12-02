package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	sliceOfNumbers := make([]int, 0)

	scanner.Scan()
	previousNumber, _ := strconv.Atoi(scanner.Text())
	sliceOfNumbers = append(sliceOfNumbers, previousNumber)

	largerMeasurements := 0

	for scanner.Scan() {
		newNumber, _ := strconv.Atoi(scanner.Text())
		sliceOfNumbers = append(sliceOfNumbers, newNumber)

		if newNumber > previousNumber {
			largerMeasurements++
		}

		previousNumber = newNumber
	}

	fmt.Printf("Part 1: %d\n", largerMeasurements)

	previousSlidingWindow := sliceOfNumbers[0] + sliceOfNumbers[1] + sliceOfNumbers[2]

	sumsLargerCount := 0

	for i := 1; i < len(sliceOfNumbers)-2; i++ {
		newSlidingWindow := sliceOfNumbers[i] + sliceOfNumbers[i+1] + sliceOfNumbers[i+2]

		if newSlidingWindow > previousSlidingWindow {
			sumsLargerCount++
		}

		previousSlidingWindow = newSlidingWindow

	}

	fmt.Printf("Part 2: %d\n", sumsLargerCount)
}
