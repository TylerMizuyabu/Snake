package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	boardSize = 38
	cellSize  = 1
)

var (
	appleCount = 0
)

type Game struct {
	gameStart bool
	gameOver  bool
	board     *Board
	snake     *Snake
	apple     *Apple
}

// TODO: Contemplate on the current way things are being handled in terms of boundary detection
// and the coordinate system (now that there are dynamic sizes). Maybe the board struct should
// handle the snakes collision detection with the wall.

func NewGame() *Game {
	board := NewBoard(boardSize, boardSize)
	minX, minY := board.GetWallWidth()
	return &Game{
		gameStart: false,
		gameOver:  false,
		board:     board,
		snake:     NewSnake(minX+boardSize/2, minY+boardSize/2),
		apple: NewApple(minX, boardSize, minY, boardSize),
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

	minX, minY := g.board.GetWallWidth()
	g.gameOver = g.snake.HitWall(minX, minY) || g.snake.HitItself()
	if g.gameOver {
		g.snake.StopMove()
	}

	if g.snake.head.HasCollided(g.apple.Entity) {
		g.snake.AddBody()
		g.snake.SpeedUp()
		/////////////////////////////////////////
		// TODO: I don't really like how this is done
		if appleCount += 1; appleCount%5 == 0 {
			g.board.ReduceSize()
		}
		minX, minY = g.board.GetWallWidth()
		width, height := g.board.GetSize()
		if appleCount%4 == 0 {
			minX += cellSize
			minY += cellSize
			width -= cellSize
			height -= cellSize
		}
		/////////////////////////////////////////
		g.apple.NewRandomCoordinates(minX, width, minY, height)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//TODO: I don't know why when I try to draw the apple/snake on the board that it doesn't clear the past images when I redraw the board
	g.board.Draw(screen)
	g.apple.Draw(screen)
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardSize + 2*cellSize, boardSize + 2*cellSize
}

func (g *Game) reset() {
	g.gameOver = false
	g.gameStart = false
	g.board.Reset()
	minX, minY := g.board.GetWallWidth()
	g.snake = NewSnake(minX+boardSize/2, minY+boardSize/2)
	g.apple = NewApple(minX, boardSize, minY, boardSize)
}
