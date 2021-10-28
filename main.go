package main

import (
	"github.com/TylerMizuyabu/snake/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Snake")
	game := game.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
