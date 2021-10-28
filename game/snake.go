package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	head *Body
	heading direction
}

type Body struct {
	*Entity
	next *Body
}

func NewBody(x, y float64) *Body {
	image := ebiten.NewImage(cellSize, cellSize)
	image.Fill(color.White)
	return &Body{
		Entity: NewEntity(x, y, cellSize, image),
		next: nil,
	}
}

func NewSnake(x, y float64) *Snake {
	return &Snake{
		head: NewBody(x, y),
	}
}

/*
The idea behind this is that when the snake eats the food, instead of adding a body part
to the end of the snake we make one where the food use to be and set it as the snakes new head
*/
func (s *Snake) AddBody(x, y float64) {
	newHead := NewBody(x, y)
	newHead.next = s.head
	s.head = newHead
}

/*
Draws each body part of the snake starting with the head
*/
func (s *Snake) Draw() {
	body := s.head
	for body != nil {
		body.Draw()
		body = body.next
	}
}

// TODO
func (s *Snake) Move() {
	switch s.heading {
	case up:
	case right:
	case down:
	case left:
	default:
		fmt.Println("Something is wrong")
	}
}

/*
Updates the snakes direction. For now it's working off the knowledge that left is the largest
direction by numerical value and up is the smallest
*/
func (s *Snake) ChangeDirection(d direction) {
	if (d >= up && d <= left) {
		s.heading = d
	}
} 