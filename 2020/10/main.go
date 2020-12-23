package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var differences = map[int]int{}
var adapters = []int{}
var dp = map[int]int{}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	adapters = append(adapters, 0)

	for scanner.Scan() {
		line := scanner.Text()

		lineToNum, _ := strconv.Atoi(line)
		adapters = append(adapters, int(lineToNum))
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	fmt.Println("part 1:", multiplyDifferences())
	fmt.Println("part 2:", countArrangements(0))
	fmt.Println(adapters)
	fmt.Println(dp)

}

func multiplyDifferences() int {

	for i := range adapters[:len(adapters)-1] {
		differences[adapters[i+1]-adapters[i]]++
	}
	return differences[1] * differences[3]
}

func countArrangements(i int) (count int) {
	if i == len(adapters)-1 {
		return 1
	}

	if val, ok := dp[i]; ok {
		return val
	}
	for j := i + 1; j < len(adapters); j++ {
		if adapters[j]-adapters[i] <= 3 {
			count += countArrangements(j)
		} else {
			break
		}
	}
	dp[i] = count
	return count
}
