package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	minWidth  = 30
	minHeight = 30
)

type Board struct {
	whiteSquare    *ebiten.Image
	blackSquare    *ebiten.Image
	originalWidth  int
	originalHeight int
	reduceCount    int
}

func NewBoard(width, height int) *Board {
	if width < minWidth {
		width = minWidth
	}
	if height < minHeight {
		height = minHeight
	}
	ws := ebiten.NewImage(int(width)+2*cellSize, int(height)+2*cellSize)
	ws.Fill(color.White)
	bs := ebiten.NewImage(int(width), int(height))
	bs.Fill(color.Black)
	return &Board{
		whiteSquare:    ws,
		blackSquare:    bs,
		originalWidth:  width,
		originalHeight: height,
		reduceCount:    0,
	}
}

func (b *Board) ReduceSize() {
	if b.originalWidth-b.reduceCount*cellSize <= minWidth || b.originalHeight-b.reduceCount*cellSize <= minHeight {
		return
	}
	b.reduceCount++
}

func (b *Board) GetSize() (int, int) {
	return b.originalWidth - 2*cellSize*b.reduceCount, b.originalHeight - 2*cellSize*b.reduceCount
}

func (b *Board) Reset() {
	b.reduceCount = 0
}

func (b *Board) Draw(dest *ebiten.Image) {
	blackSquareOp := &ebiten.DrawImageOptions{}
	width, height := b.GetSize()
	minX, minY := b.GetWallWidth()
	blackSquareOp.GeoM.Scale(float64(width)/float64(b.originalWidth), float64(height)/float64(b.originalHeight))
	blackSquareOp.GeoM.Translate(float64(minX), float64(minY))
	dest.DrawImage(b.whiteSquare, nil)
	dest.DrawImage(b.blackSquare, blackSquareOp)
}

func (b *Board) GetWallWidth() (int, int) {
	return cellSize * (1 + b.reduceCount), cellSize * (1 + b.reduceCount)
}
