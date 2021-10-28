package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Apple struct {
	*Entity
}

func NewApple(x, y float64) *Apple {
	image := ebiten.NewImage(cellSize, cellSize)
	// image.Fill(color.RGBA{R: 255, G: 0, B:0, A:1})
	image.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	return &Apple{
		NewEntity(x, y, cellSize, image),
	}
}

func (a *Apple) SetCoordinates(x, y float64) {
	a.x = x
	a.y = y
}
