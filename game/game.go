package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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

type Game struct {
	snake *Snake
}

func NewGame() *Game {
	return &Game{
		snake: NewSnake(boardSize/2, boardSize/2),
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.snake.ChangeDirection(up)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.snake.ChangeDirection(right)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.snake.ChangeDirection(down)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.snake.ChangeDirection(left)
	}
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardSize, boardSize
}