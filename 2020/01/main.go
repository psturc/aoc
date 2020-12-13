package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arrOfNum := make([]int, 0)
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		arrOfNum = append(arrOfNum, num)
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	for i := range arrOfNum {
		for j := range arrOfNum {
			if arrOfNum[i]+arrOfNum[j] == 2020 {
				fmt.Println(arrOfNum[i], arrOfNum[j])
				fmt.Println(arrOfNum[i] * arrOfNum[j])
			}
			for k := range arrOfNum {
				if arrOfNum[i]+arrOfNum[j]+arrOfNum[k] == 2020 {
					fmt.Println(arrOfNum[i], arrOfNum[j], arrOfNum[k])
					fmt.Println(arrOfNum[i] * arrOfNum[j] * arrOfNum[k])
				}
			}
		}
	}

}
