package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var passportModel = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func main() {
	passport := make(map[string]string)
	var passportsValid, passportsDataValid int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 {
			for _, v := range strings.Split(line, " ") {
				parts := strings.Split(v, ":")
				passport[parts[0]] = parts[1]
			}
		} else {
			if isPassportValid(passport) {
				passportsValid++
				if isPassportsDataValid(passport) {
					passportsDataValid++
				}
			}
			passport = make(map[string]string)
		}
	}
	fmt.Println("part 1", passportsValid)
	fmt.Println("part 2", passportsDataValid)

}

func isPassportValid(passport map[string]string) bool {
	if len(passport) < 7 {
		return false
	}

	for _, v := range passportModel {
		if _, ok := passport[v]; !ok && v != "cid" {
			return false
		}
	}
	return true
}

func isPassportsDataValid(passport map[string]string) bool {
	for k, v := range passport {
		switch k {
		case "byr":
			n, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if n < 1920 || n > 2002 {
				return false
			}
		case "iyr":
			n, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if n < 2010 || n > 2020 {
				return false
			}

		case "eyr":
			n, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if n < 2020 || n > 2030 {
				return false
			}
		case "hgt":
			match, err := regexp.MatchString("^(1([5-8][0-9]|9[0-3])cm)$|^((59|[6][0-9]|7[0-6])in)$", v)
			if err != nil {
				panic(err)
			}
			if !match {
				return false
			}
		case "hcl":
			match, err := regexp.MatchString("^#([0-9]|[a-f]){6}$", v)
			if err != nil {
				panic(err)
			}
			if !match {
				return false
			}
		case "ecl":
			match, err := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", v)
			if err != nil {
				panic(err)
			}
			if !match {
				return false
			}
		case "pid":
			match, err := regexp.MatchString("^[0-9]{9}$", v)
			if err != nil {
				panic(err)
			}
			if !match {
				return false
			}
		}
	}
	return true
}
