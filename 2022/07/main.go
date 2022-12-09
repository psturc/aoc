package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Dir struct {
	parent   string
	children []string
	size     int
}

const totalDiskSpace int = 70000000
const requiredFreeSpace int = 30000000

var dirs = map[string]Dir{}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	var currentDir string
	var part1 int
	var part2 int = math.MaxInt

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			sp := strings.Fields(line)
			switch sp[1] {
			case "cd":
				if sp[2] == ".." {
					currentDir = dirs[currentDir].parent
				} else if sp[2] != "/" {
					currentDir = currentDir + sp[2] + "/"
				} else {
					currentDir = "/"
				}
			case "ls":
				continue
			}
		} else if strings.HasPrefix(line, "dir") {
			newDirName := currentDir + strings.Fields(line)[1] + "/"
			dir := dirs[currentDir]
			dir.children = append(dir.children, newDirName)
			dirs[currentDir] = dir
			dirs[newDirName] = Dir{parent: currentDir}
		} else {
			fileSize, err := strconv.Atoi(strings.Fields(line)[0])
			if err != nil {
				panic(err)
			}
			updateParents(currentDir, int(fileSize))

		}
	}

	var usedSpace = dirs["/"].size

	for name := range dirs {
		if dirs[name].size <= 100000 {
			part1 += dirs[name].size
		}

		if dirs[name].size < part2 && (totalDiskSpace-requiredFreeSpace) >= (usedSpace-dirs[name].size) {
			part2 = dirs[name].size
		}
	}
	fmt.Println("1:", part1)
	fmt.Println("2:", part2)
}

func updateParents(name string, size int) {
	if dirs[name].parent != "" {
		updateParents(dirs[name].parent, size)
	}
	dir := dirs[name]
	dir.size += size
	dirs[name] = dir

}
