package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	var udc int
	for _, line := range lines {
		r := strings.Split(line, "|")[1]

		o := strings.Fields(r)

		for _, seg := range o {
			if (len(seg) >= 2 && len(seg) <= 4) || len(seg) == 7 {
				udc++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", udc)

	var result int
	for _, line := range lines {
		var nums = map[int]string{}
		var segmap = map[string]string{
			"a": "",
			"b": "",
			"c": "",
			"d": "",
			"e": "",
			"f": "",
			"g": "",
		}
		ls := strings.Split(line, "|")
		l, r := ls[0], ls[1]

		for _, v := range strings.Fields(l) {
			switch len(v) {
			case 2:
				nums[1] = v
			case 4:
				nums[4] = v
			case 3:
				nums[7] = v
			case 7:
				nums[8] = v
			}
		}
		// Figure out "a"
		for _, ch := range nums[7] {
			if !strings.Contains(nums[1], string(ch)) {
				segmap["a"] = string(ch)
				break
			}
		}
		// Figure out "c"
	out:
		for _, v := range strings.Fields(l) {

			if len(v) == 6 {
				for _, ch := range nums[1] {
					if !strings.Contains(v, string(ch)) {
						segmap["c"] = string(ch)
						nums[6] = v
						break out
					}
				}
			}
		}

		// Figure out "f"
		if string(nums[1][0]) == segmap["c"] {
			segmap["f"] = string(nums[1][1])
		} else {
			segmap["f"] = string(nums[1][0])
		}

		// Figure out 2, 3 and 5
		for _, v := range strings.Fields(l) {

			if len(v) == 5 {
				if !strings.Contains(v, segmap["f"]) {
					nums[2] = v
				} else if !strings.Contains(v, segmap["c"]) {
					nums[5] = v
				} else {
					nums[3] = v
				}
			}
		}

		// Figure out 9 and 0
		for _, v := range strings.Fields(l) {

			if len(v) == 6 && nums[6] != v {
				for _, ch := range nums[4] {
					if !strings.Contains(v, string(ch)) {
						nums[0] = v
						break
					}
				}
				if nums[0] != v {
					nums[9] = v
				}
			}
		}

		// Figure out "b"
		for _, ch := range nums[9] {
			if !strings.Contains(nums[3], string(ch)) {
				segmap["b"] = string(ch)
				break
			}
		}

		// Figure out "d" and "e"
		for _, ch := range nums[8] {
			if !strings.Contains(nums[0], string(ch)) {
				segmap["d"] = string(ch)
			}
			if !strings.Contains(nums[9], string(ch)) {
				segmap["e"] = string(ch)
			}
		}

		// Figure out "g"
		keys := ""
		for _, v := range segmap {
			keys += v
		}
		for k := range segmap {
			if !strings.Contains(keys, k) {
				segmap["g"] = k
				break
			}
		}

		// Decode the number and sum
		var dnstr string = ""
		var dnint int
		for _, n := range strings.Fields(r) {
			dnstr += fmt.Sprintf("%d", decodeNumber(n, segmap))
		}

		dnint, _ = strconv.Atoi(dnstr)
		result += dnint

	}
	fmt.Printf("Part 2: %d\n", result)
}

func decodeNumber(numStr string, segmap map[string]string) int {
	if len(numStr) == 2 {
		return 1
	}
	if len(numStr) == 4 {
		return 4
	}
	if len(numStr) == 3 {
		return 7
	}
	if len(numStr) == 7 {
		return 8
	}
	if len(numStr) == 5 {
		if strings.Contains(numStr, segmap["e"]) {
			return 2
		}
		if strings.Contains(numStr, segmap["b"]) {
			return 5
		}
		return 3
	}
	if len(numStr) == 6 {
		if !strings.Contains(numStr, segmap["e"]) {
			return 9
		}
		if !strings.Contains(numStr, segmap["d"]) {
			return 0
		}
	}
	return 6
}
