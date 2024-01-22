package main

import (
	"github.com/aaronmahlke/goo/gameloop"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := gameloop.Setup()

	// Game loop
	for !rl.WindowShouldClose() {
		gameloop.Loop(&game)
	}
}
