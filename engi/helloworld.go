package main

import (
	"github.com/ajhager/engi"
)

// Game struct
type Game struct {
	*engi.Game
	bot   engi.Drawable
	batch *engi.Batch
	font  *engi.Font

	botX, botY float32
	keyMap     map[engi.Key]bool
}

// Preload .. Setup Data
func (game *Game) Preload() {
	engi.Files.Add("bot", "data/icon.png")
	engi.Files.Add("font", "data/font.png")
	game.batch = engi.NewBatch(engi.Width(), engi.Height())

	game.keyMap = make(map[engi.Key]bool)
	game.botX, game.botY = 512, 320

}

// Setup Routine
func (game *Game) Setup() {
	engi.SetBg(0x2d3779) // 背景色
	game.bot = engi.Files.Image("bot")
	game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)

}

// Render ring
func (game *Game) Render() {
	game.batch.Begin()
	game.batch.Draw(game.bot, game.botX, game.botY, 0.5, 0.5, 10, 10, 0, 0xffffff, 1)
	game.font.Print(game.batch, "ENGI", 475, 200, 0xffffff)
	game.batch.End()
}

// Key push event routine
func (game *Game) Key(key engi.Key, modifier engi.Modifier, action engi.Action) {

	switch action {
	case engi.PRESS:
		game.keyMap[key] = true
	case engi.RELEASE:
		game.keyMap[key] = false
	}
}

// Update
func (game *Game) Update(dt float32) {
	var dx, dy float32
	if game.keyMap[engi.ArrowUp] {
		dy = -500
	}
	if game.keyMap[engi.ArrowDown] {
		dy = 500
	}

	if game.keyMap[engi.ArrowLeft] {
		dx = -500
	}

	if game.keyMap[engi.ArrowRight] {
		dx = 500
	}
	dx, dy = dx*dt, dy*dt
	game.botX += dx
	game.botY += dy

}
func main() {
	engi.Open("Hello", 1024, 640, false, &Game{})
}
