package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type bags map[string][]bag
type bag struct {
	name     string
	quantity int
}

var mapOfBags = bags{}
var wantedBags = map[string]bool{}

func main() {
	file, _ := os.Open("input.txt")
	reReplace := regexp.MustCompile(`\s?bag[s]?\s?|\s?contain no other|[.]`)
	reSplit := regexp.MustCompile(`\s?contain\s?|\s?,\s?`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = reReplace.ReplaceAllString(line, "")
		parts := reSplit.Split(line, -1)

		bagName := parts[0]
		mapOfBags[bagName] = []bag{}

		for i := 1; i < len(parts); i++ {
			bn := string(parts[i][2:])
			bq, _ := strconv.Atoi(string(parts[i][0]))
			mapOfBags[bagName] = append(mapOfBags[bagName], bag{name: bn, quantity: bq})
			if bn == "shiny gold" {
				wantedBags[bagName] = true
			}
		}
	}

	for n := range mapOfBags {
		b := bag{name: n}
		if b.doesContainShinyGoldBag() {
			wantedBags[n] = true
		}
	}

	fmt.Println("part 1:", len(wantedBags))
	fmt.Println("part 2:", bag{name: "shiny gold"}.getNumberOfInsideBags())
}

func (b bag) doesContainShinyGoldBag() bool {
	for _, v := range mapOfBags[b.name] {
		if v.doesContainShinyGoldBag() {
			wantedBags[v.name] = true
			wantedBags[b.name] = true
		}
	}
	if val, ok := wantedBags[b.name]; ok && val {
		return true
	}

	return false
}

func (b bag) getNumberOfInsideBags() int {
	var total int
	for _, v := range mapOfBags[b.name] {
		total += v.quantity + v.quantity*v.getNumberOfInsideBags()
	}
	return total

}
