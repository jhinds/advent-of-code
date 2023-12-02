package main

import (
	"fmt"
	"log"
	"strconv"
)

func dayOnePartOne() {
	lines := loadLines(1)
	total := 0
	for _, line := range lines {
		var nums []string
		for _, c := range line {
			ch := fmt.Sprintf("%c", c)
			if _, err := strconv.Atoi(ch); err == nil {
				nums = append(nums, ch)
			}
		}
		numStr := fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("error! %s", err.Error())
		}
		total += num
	}
	fmt.Println("answer: ", total)
}

func dayOnePartTwo() {
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	lines := loadLines(1)

	total := 0
	for _, line := range lines {
		var nums []string
		for i, c := range line {
			ch := fmt.Sprintf("%c", c)
			if _, err := strconv.Atoi(ch); err == nil {
				nums = append(nums, ch)
			} else {
				for numberWord, nint := range numMap {
					if len(numberWord)+i <= len(line) {
						if line[i:len(numberWord)+i] == numberWord {
							nums = append(nums, string(strconv.Itoa(nint)))
						}
					}
				}
			}
		}
		numStr := fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("error! %s", err.Error())
		}
		total += num
	}
	fmt.Println("answer: ", total)
}
