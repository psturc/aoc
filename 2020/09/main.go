package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	series := make([]uint64, 0)
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		lineToNum, _ := strconv.Atoi(line)
		series = append(series, uint64(lineToNum))
	}

	invalidNumber := findFirstRuleViolation(series)
	fmt.Println("part 1:", invalidNumber)
	fmt.Println("part 2:", findEncryptionWeakness(series, invalidNumber))

}

func findFirstRuleViolation(series []uint64) uint64 {
	for i := 25; i < len(series); i++ {
		if !isValidPreamble(series[i-25:i], series[i]) {
			return series[i]
		}
	}
	return 0
}

func isValidPreamble(preamble []uint64, testedNumber uint64) (valid bool) {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if i == j {
				continue
			}
			if preamble[i]+preamble[j] == testedNumber {
				return true
			}
		}
	}
	return
}

func findEncryptionWeakness(series []uint64, invalidNumber uint64) uint64 {
	sumCandidates := make([]uint64, 0)
	for _, v := range series {
		sumCandidates = append(sumCandidates, v)

		for {
			if sum(sumCandidates) > invalidNumber {
				sumCandidates = sumCandidates[1:]
				continue
			}
			break
		}
		if sum(sumCandidates) == invalidNumber {
			sort.Slice(sumCandidates, func(i, j int) bool { return sumCandidates[i] < sumCandidates[j] })
			return sumCandidates[0] + sumCandidates[len(sumCandidates)-1]
		}

	}
	return 0
}

func sum(slice []uint64) (sum uint64) {
	for _, v := range slice {
		sum += v
	}
	return
}
