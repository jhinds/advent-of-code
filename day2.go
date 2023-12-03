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
	for _, set := range g.Sets {
		if set.Blue > blueMax {
			return false
		} else if set.Green > greenMax {
			return false
		} else if set.Red > redMax {
			return false
		}
	}
	return true
}

func (g *Game) MinimumSet() *Set {
	red := 0
	green := 0
	blue := 0
	for _, set := range g.Sets {
		if set.Blue > blue {
			blue = set.Blue
		}
		if set.Green > green {
			green = set.Green
		}
		if set.Red > red {
			red = set.Red
		}
	}
	return &Set{
		Red:   red,
		Blue:  blue,
		Green: green,
	}
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
func dayTwoPartTwo() {
	lines := loadLines(2)
	powerSum := 0
	for _, line := range lines {
		game := NewGame(line)
		set := game.MinimumSet()
		power := set.Green * set.Red * set.Blue
		powerSum += power
	}
	fmt.Println(powerSum)
}
