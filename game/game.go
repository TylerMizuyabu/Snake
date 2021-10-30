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

type Game struct {
	gameStart bool
	gameOver  bool
	snake     *Snake
	apple     *Apple
}

func NewGame() *Game {
	return &Game{
		gameStart: false,
		gameOver:  false,
		snake:     NewSnake(boardSize/2, boardSize/2),
		apple:     NewApple(float64(rand.Intn(boardSize)), float64(rand.Intn(boardSize))),
	}
}

func (g *Game) Update() error {
	if !g.gameStart {
		g.gameStart = true
		g.snake.StartMove()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
		return nil
	} else if g.gameOver {
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

	g.gameOver = g.snake.HitWall() || g.snake.HitItself()
	if g.gameOver {
		g.snake.StopMove()
	}

	if g.snake.head.HasCollided(g.apple.Entity) {
		g.snake.AddBody()
		g.snake.SpeedUp()
		g.apple.SetCoordinates(float64(rand.Intn(boardSize)), float64(rand.Intn(boardSize)))
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
	g.gameStart = false
	g.snake = NewSnake(boardSize/2, boardSize/2)
	g.apple = NewApple(rand.Float64()*boardSize, rand.Float64()*boardSize)
}
