package game

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	x     int
	y     int
	size  int
	image *ebiten.Image
}

func NewEntity(x, y, size int, image *ebiten.Image) *Entity {
	return &Entity{
		x, y, size, image,
	}
}

func (e *Entity) Draw(dst *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(float64(e.x), float64(e.y))
	dst.DrawImage(e.image, &ebiten.DrawImageOptions{
		GeoM:          geom,
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	})
}

func (e *Entity) HasCollided(e2 *Entity) bool {
	return e.x < e2.x+e2.size &&
		e.x+e.size > e2.x &&
		e.y < e2.y+e2.size &&
		e.size+e.y > e2.y
}
