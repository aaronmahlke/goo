package gameloop

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW = 800
	screenH = 450
)

type CharacterTexture struct {
	active rl.Texture2D
	Up     rl.Texture2D
	Down   rl.Texture2D
	Left   rl.Texture2D
	Right  rl.Texture2D
}

type Player struct {
	sourceRect rl.Rectangle
	destRect   rl.Rectangle
	pos        rl.Vector2
	width      int
	height     int
	texture    CharacterTexture
	frame      int
	isMoving   bool
}

type Game struct {
	player     Player
	frameCount int
}

func Setup() Game {
	rl.InitWindow(screenW, screenH, "Goo")
	rl.SetTargetFPS(60)

	var player Player

	// Setup Player
	player.texture.Up = rl.LoadTexture("assets/ACharUp.png")
	player.texture.Down = rl.LoadTexture("assets/ACharDown.png")
	player.texture.Left = rl.LoadTexture("assets/ACharLeft.png")
	player.texture.Right = rl.LoadTexture("assets/ACharRight.png")

	player.width = 48 / 2
	player.height = 48 / 2

	player.isMoving = false

	player.pos = rl.Vector2{
		X: screenW / 2.0,
		Y: -screenH / 10,
	}

	// source Rect
	player.sourceRect = rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(player.width),
		Height: float32(player.height),
	}

	// destination Rect
	player.destRect = rl.Rectangle{
		X:      screenW / 2.0,
		Y:      screenH / 2.0,
		Width:  float32(player.width * 3),
		Height: float32(player.height * 3),
	}

	game := Game{
		player: player,
	}

	return game
}

func Loop(game *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	rl.DrawTexturePro(
		game.player.texture.active,
		game.player.sourceRect,
		game.player.destRect,
		game.player.pos,
		0,
		rl.White,
	)
	rl.EndDrawing()

	game.player.isMoving = false
	game.frameCount++

	if rl.IsKeyDown(rl.KeyRight) {
		game.player.pos.X -= 5
		game.player.texture.active = game.player.texture.Right
		game.player.isMoving = true
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		game.player.pos.X += 5
		game.player.texture.active = game.player.texture.Left
		game.player.isMoving = true
	}

	if rl.IsKeyDown(rl.KeyUp) {
		game.player.pos.Y += 5
		game.player.texture.active = game.player.texture.Up
		game.player.isMoving = true
	}

	if rl.IsKeyDown(rl.KeyDown) {
		game.player.pos.Y -= 5
		game.player.texture.active = game.player.texture.Down
		game.player.isMoving = true
	}

	if game.player.isMoving {
		animatePlayer(game, &game.player)
	} else {
		game.player.sourceRect.X = 0
	}
}

func animatePlayer(game *Game, player *Player) {
	if game.frameCount%8 == 1 {
		player.frame++
		player.frame %= 2
	}

	player.sourceRect.X = float32(player.width * player.frame)
}
