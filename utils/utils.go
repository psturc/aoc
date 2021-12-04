package utils

import (
	"bufio"
	"os"
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
