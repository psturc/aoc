package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FileLinesToSlice(filename string) (slice []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	return
}

func StrSliceToIntSlice(in []string) []int {
	out := make([]int, len(in))
	for i, s := range in {
		out[i], _ = strconv.Atoi(s)
	}
	return out
}

func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Contains(slc []string, str string) bool {
	for _, v := range slc {
		if v == str {
			return true
		}
	}
	return false
}

func ContainsInt(slc []int, i int) bool {
	for _, v := range slc {
		if v == i {
			return true
		}
	}
	return false
}

func IsLowercase(str string) bool {
	return strings.ToLower(str) == str
}
