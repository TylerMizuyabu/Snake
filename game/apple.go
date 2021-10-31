package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Apple struct {
	*Entity
}

func NewApple(xMin, width, yMin, height int) *Apple {
	image := ebiten.NewImage(cellSize, cellSize)
	// image.Fill(color.RGBA{R: 255, G: 0, B:0, A:1})
	image.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	apple := &Apple{
		NewEntity(0, 0, cellSize, image),
	}
	apple.NewRandomCoordinates(xMin, width, yMin, height)
	return apple
}

func (a *Apple) NewRandomCoordinates(xMin, width, yMin, height int) {
	a.x = xMin + rand.Intn(width-cellSize)
	a.y = yMin + rand.Intn(height-cellSize)
}
