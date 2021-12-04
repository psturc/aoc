package utils

import (
	"bufio"
	"os"
	"strconv"
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
