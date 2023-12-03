package main

import "testing"

func TestGame_MinimumSet(t *testing.T) {
	var tests = []struct {
		gameStr    string
		MinimumSet Set
	}{
		// the table itself
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Set{Red: 4, Green: 2, Blue: 6}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Set{Red: 1, Green: 3, Blue: 4}},
	}

	for _, tt := range tests {
		game := NewGame(tt.gameStr)
		if *game.MinimumSet() != tt.MinimumSet {
			t.Errorf("game: %#v, got %#v, expected %#v", game, game.MinimumSet(), tt.MinimumSet)
		}
	}
}
