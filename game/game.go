package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	boardSize = 40.0
	cellSize  = 1.0
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type Game struct {
	gameOver bool
	snake    *Snake
	apple    *Apple
}

var i = 0

func NewGame() *Game {
	return &Game{
		gameOver: false,
		snake:    NewSnake(boardSize/2, boardSize/2),
		apple:    NewApple(rand.Float64()*boardSize, rand.Float64()*boardSize),
	}
}


func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
		return nil
	}else if g.gameOver {
		return nil
	}

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

	if i % 4 == 0 {
		g.snake.Move()
	}
	i++
	g.gameOver = g.snake.HitWall() || g.snake.HitItself()
	if g.snake.head.HasCollided(g.apple.Entity) {
		g.snake.AddBody()
		g.apple.SetCoordinates(rand.Float64()*boardSize, rand.Float64()*boardSize)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.apple.Draw(screen)
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardSize, boardSize
}

func (g *Game) reset() {
	g.gameOver = false
	g.snake = NewSnake(boardSize/2, boardSize/2)
	g.apple = NewApple(rand.Float64()*boardSize, rand.Float64()*boardSize)
}