package game

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	x float64
	y float64
	size float64
	image *ebiten.Image
}

func NewEntity(x, y, size float64, image *ebiten.Image) *Entity {
	return &Entity{
		x, y, size, image,
	}
}

func (e *Entity) Draw() {
	geom := ebiten.GeoM{}
	geom.Translate(e.x, e.y)
	e.image.DrawImage(e.image, &ebiten.DrawImageOptions{
		GeoM:          geom,
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	})
}