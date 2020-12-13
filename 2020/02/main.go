package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordObject struct {
	min       int
	max       int
	char      string
	password  string
	charCount int
}

func (pwdO passwordObject) isValidPartOne() bool {
	for i := range pwdO.password {
		if string(pwdO.password[i]) == pwdO.char {
			pwdO.charCount++
		}
	}
	if pwdO.charCount < pwdO.min {
		return false
	}
	if pwdO.charCount > pwdO.max {
		return false
	}
	return true
}

func (pwdO passwordObject) isValidPartTwo() bool {

	if string(pwdO.password[pwdO.min-1]) != pwdO.char && string(pwdO.password[pwdO.max-1]) != pwdO.char {
		return false
	}

	if string(pwdO.password[pwdO.min-1]) == pwdO.char && string(pwdO.password[pwdO.max-1]) == pwdO.char {
		return false
	}
	return true
}

func main() {
	var passwordsValidPartOne, passwordsValidPartTwo int

	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		pwdO := newPasswordObject(line)

		if pwdO.isValidPartOne() {
			passwordsValidPartOne++
		}

		if pwdO.isValidPartTwo() {
			passwordsValidPartTwo++
		}
	}

	fmt.Println("part 1: ", passwordsValidPartOne)
	fmt.Println("part 2: ", passwordsValidPartTwo)
}

func newPasswordObject(line string) passwordObject {

	line = strings.ReplaceAll(line, "-", " ")
	line = strings.ReplaceAll(line, ":", "")

	parts := strings.Fields(line)

	if len(parts) != 4 {
		panic(fmt.Sprintln("invalid input: len of parts: ", len(parts)))
	}

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	pwdO := passwordObject{min: min, max: max, char: parts[2], password: parts[3]}

	return pwdO
}
