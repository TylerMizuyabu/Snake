package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	head    *Body
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
		next:   nil,
	}
}

func NewSnake(x, y float64) *Snake {
	return &Snake{
		head:    NewBody(x, y),
		heading: up,
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
func (s *Snake) Draw(dst *ebiten.Image) {
	body := s.head
	for body != nil {
		body.Draw(dst)
		body = body.next
	}
}

func (s *Snake) Move() {
	nextX, nextY := s.head.x, s.head.y
	switch s.heading {
	case up:
		nextY -= 0.5
	case right:
		nextX += 0.5
	case down:
		nextY += 0.5
	case left:
		nextX -= 0.5
	default:
		fmt.Println("Something is wrong")
	}
	body := s.head
	for body != nil {
		tempX, tempY := body.x, body.y
		body.x = nextX
		body.y = nextY
		nextX, nextY = tempX, tempY
		body = body.next
	}
}

/*
Updates the snakes direction. For now it's working off the knowledge that left is the largest
direction by numerical value and up is the smallest
*/
func (s *Snake) ChangeDirection(d direction) {
	if d >= up && d <= left {
		s.heading = d
	}
}

func (s *Snake) HitWall() bool {
	return s.head.x < 0 || s.head.x > boardSize || s.head.y < 0 || s.head.y > boardSize
}
