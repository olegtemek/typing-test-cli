package main

import (
	"github.com/olegtemek/typing-test-cli/internal/game"
	"github.com/olegtemek/typing-test-cli/internal/mode"
	"github.com/olegtemek/typing-test-cli/internal/utils"
	"github.com/olegtemek/typing-test-cli/internal/word"
)

func main() {
	utils.ClearTerminal()

	fileName, err := mode.Init()
	if err != nil {
		panic(err)
	}

	words, err := word.Init(fileName)
	if err != nil {
		panic(err)
	}

	game := game.New(words)

	game.Start()
	game.Stats()
}
