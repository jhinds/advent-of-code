package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// loadLines will load the data for a day into a string slice
func loadLines(day int) []string {
	body, err := os.Open(fmt.Sprintf("data/day%d.txt", day))
	defer body.Close()
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fileScanner := bufio.NewScanner(body)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
