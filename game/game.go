package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	boardSize = 38.0
	cellSize  = 1.0
)

type Game struct {
	gameStart bool
	gameOver  bool
	board     *Board
	snake     *Snake
	apple     *Apple
}

func NewGame() *Game {
	return &Game{
		gameStart: false,
		gameOver:  false,
		board:     NewBoard(boardSize, boardSize),
		snake:     NewSnake(boardSize/2, boardSize/2),
		apple:     NewApple(float64(cellSize+rand.Intn(int(boardSize)-1)), float64(cellSize+rand.Intn(int(boardSize)-1))),
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

	g.gameOver = g.snake.HitWall(cellSize, cellSize) || g.snake.HitItself()
	if g.gameOver {
		g.snake.StopMove()
	}

	if g.snake.head.HasCollided(g.apple.Entity) {
		g.snake.AddBody()
		g.snake.SpeedUp()
		g.apple.SetCoordinates(float64(cellSize+rand.Intn(int(boardSize)-1)), float64(cellSize+rand.Intn(int(boardSize)-1)))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
	g.apple.Draw(screen)
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardSize+2*cellSize, boardSize+2*cellSize
}

func (g *Game) reset() {
	g.gameOver = false
	g.gameStart = false
	g.snake = NewSnake(cellSize+boardSize/2, cellSize+boardSize/2)
	g.apple = NewApple(float64(cellSize+rand.Intn(int(boardSize)-1)), float64(cellSize+rand.Intn(int(boardSize)-1)))
}
