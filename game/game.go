package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	boardSize = 40.0
	cellSize = 1.0
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type Game struct {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardSize, boardSize
}