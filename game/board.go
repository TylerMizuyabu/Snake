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
	originalWidth  float64
	originalHeight float64
	width          float64
	height         float64
}

func NewBoard(width, height float64) *Board {
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
		width:          width,
		height:         height,
	}
}

func (b *Board) ReduceSize() {
	b.width -= 1
	b.height -= 1
}

func (b *Board) GetSize() (float64, float64) {
	return b.width, b.height
}

func (b *Board) Reset() {
	b.width = b.originalWidth
	b.height = b.originalHeight
}

func (b *Board) Draw(dest *ebiten.Image) {
	blackSquareOp := &ebiten.DrawImageOptions{}
	blackSquareOp.GeoM.Scale(b.width/b.originalWidth, b.height/b.originalHeight)
	blackSquareOp.GeoM.Translate(cellSize*(1+b.originalWidth-b.width), cellSize*(1+b.originalHeight-b.height))

	dest.DrawImage(b.whiteSquare, nil)
	b.whiteSquare.DrawImage(b.blackSquare, blackSquareOp)
}