package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	arg       int
	visited   bool
}

var instructions = []instruction{}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		operation := parts[0]
		arg, _ := strconv.Atoi(parts[1])

		instructions = append(instructions, instruction{operation: operation, arg: arg})
	}

	accVal, _ := bootCodeAccValue(instructions)
	fmt.Println("part 1:", accVal)

	fixedCodeAccVal := fixedCodeAccValue()
	fmt.Println("part 2:", fixedCodeAccVal)
}

func bootCodeAccValue(instructions []instruction) (accValue int, looped bool) {
	var bootCodePosition int
	copyOfInstructions := make([]instruction, len(instructions))
	copy(copyOfInstructions, instructions)

	for {
		if bootCodePosition >= len(copyOfInstructions) {
			break
		}
		inst := copyOfInstructions[bootCodePosition]

		if !inst.visited {
			copyOfInstructions[bootCodePosition].visited = true
		} else {
			looped = true
			break
		}
		switch inst.operation {
		case "acc":
			accValue += inst.arg
			bootCodePosition++
		case "jmp":
			bootCodePosition += inst.arg
		case "nop":
			bootCodePosition++
		}
	}
	return
}

func fixedCodeAccValue() (accValue int) {
	var looped bool = true
	updatedInstructions := make([]instruction, len(instructions))

	for i := range instructions {
		if instructions[i].operation == "jmp" || instructions[i].operation == "nop" {
			copy(updatedInstructions, instructions)

			if instructions[i].operation == "jmp" {
				updatedInstructions[i].operation = "nop"
			} else {
				updatedInstructions[i].operation = "jmp"
			}
			accValue, looped = bootCodeAccValue(updatedInstructions)
			if !looped {
				return
			}
		}
	}
	return
}
