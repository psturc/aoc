package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

var (
	binQueue     string
	packetLength int
	versionsum   int
)

func main() {
	input := utils.FileLinesToSlice("input.txt")
	input = strings.Split(input[0], "")

	/*
		1. Parse all hex numbers to generate 8 bits
		2. Determinte version (3 bits)
		3. Determine type ID (3 bits)
		4. Keep remaining 2 bits for later
		4. If ID == 4
			1. Parse
	*/

	for _, hexch := range input {
		dec, _ := strconv.ParseInt(hexch, 16, 0)
		binQueue += fmt.Sprintf("%04b", dec)
	}

	for {
		versionsum += binToDec(binQueue[0:3])
		id := binToDec(binQueue[3:6])

		binQueue = binQueue[6:]
		packetLength = 6

		if id == 4 {
			getLiteralValue()
		}

		fmt.Println("binQueue:", binQueue, "packageLength:", packetLength)
		if packetLength%4 != 0 {
			binQueue = binQueue[4-packetLength%4:]
		}

		packetLength = 0
		if len(binQueue) == 0 {
			fmt.Println("THE END!")
			break
		}
	}

	fmt.Printf("Part 1: %d\n", versionsum)

}

func binToDec(bin string) int {
	dec, _ := strconv.ParseInt(bin, 2, 0)
	return int(dec)
}

func getLiteralValue() {
	var valueQueue string
	for {
		group := binQueue[0:5]
		binQueue = binQueue[5:]
		valueQueue += group[1:]

		packetLength += 5
		if string(group[0]) == "0" {
			break
		}
	}
	fmt.Println("literal value:", binToDec(valueQueue))

}
