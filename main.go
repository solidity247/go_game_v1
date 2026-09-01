package main

import (
	"os"
	"strings"

	"github.com/solidity247/go_game_v1/game"
)

var handleCommand func(c string) string

func main() {
	initGame()
}

func initGame() {
	myGame := game.New()
	handleCommand = myGame.RunRawCommand
	if !strings.Contains(os.Args[0], "test") {
		myGame.Play()
	}
}
