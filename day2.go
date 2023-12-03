package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	redMax   = 12
	greenMax = 13
	blueMax  = 14
)

type Set struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	ID   int
	Sets []Set
}

func NewGame(gameText string) *Game {
	numRe := regexp.MustCompile("[0-9]+")
	colorRe := regexp.MustCompile("[a-z]+")

	result := strings.Split(gameText, ":")
	gameId, err := strconv.Atoi(strings.TrimPrefix(result[0], "Game "))
	if err != nil {
		log.Fatalf("could not parse game: %s", err.Error())
	}
	//Game 2: 3 red, 1 blue, 2 green; 1 blue, 9 green; 1 red, 10 green
	strSets := strings.Split(result[1], ";")
	var sets []Set
	for _, strSet := range strSets {
		var set Set
		colors := strings.Split(strSet, ",")
		for _, color := range colors {
			strColorCount := numRe.FindAllString(color, -1)[0]
			colorCount, err := strconv.Atoi(strColorCount)
			if err != nil {
				log.Fatalf("could not parse color count: %s", err.Error())
			}
			switch color := colorRe.FindAllString(color, -1)[0]; color {
			case "green":
				set.Green += colorCount
			case "red":
				set.Red += colorCount
			case "blue":
				set.Blue += colorCount
			default:
				fmt.Println("this shouldn't happen, color: ", color)
			}
		}
		sets = append(sets, set)
	}

	return &Game{
		ID:   gameId,
		Sets: sets,
	}
}

func (g *Game) ValidGame() bool {
	var redTotal int
	var greenTotal int
	var blueTotal int

	for _, set := range g.Sets {
		redTotal += set.Red
		greenTotal += set.Green
		blueTotal += set.Blue
	}
	if blueTotal > blueMax {
		return false
	} else if greenTotal > greenMax {
		return false
	} else if redTotal > redMax {
		return false
	}
	return true
}

func dayTwoPartOne() {
	lines := loadLines(2)

	gameSum := 0
	for _, line := range lines {
		game := NewGame(line)
		if game.ValidGame() {
			gameSum += game.ID
		}
	}
	fmt.Println(gameSum)
}
func dayTwoPartTwo() {}
