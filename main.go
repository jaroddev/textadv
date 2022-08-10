package main

import (
	"github.com/jaroddev/textadventure/data"
)

func main() {
	game := data.CreateStory()

	game.Play()
}
