package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	startSpeed = time.Millisecond * 100
	minSpeed   = time.Millisecond * 30
)

type Snake struct {
	head       *Body
	tail       *Body
	heading    direction
	stopMoving chan bool
	speed      time.Duration
}

type Body struct {
	*Entity
	next *Body
	prev *Body
}

func NewBody(x, y int, prev *Body) *Body {
	image := ebiten.NewImage(cellSize, cellSize)
	image.Fill(color.White)
	return &Body{
		Entity: NewEntity(x, y, cellSize, image),
		next:   nil,
		prev:   prev,
	}
}

func NewSnake(x, y int) *Snake {
	head := NewBody(x, y, nil)
	return &Snake{
		head:       head,
		tail:       head,
		heading:    up,
		stopMoving: make(chan bool),
		speed:      startSpeed,
	}
}

func (s *Snake) AddBody() {
	dirToAdd := oppositeDirection(s.heading)
	if s.tail.prev != nil {
		xDiff, yDiff := s.tail.prev.x-s.tail.x, s.tail.prev.y-s.tail.y
		if xDiff > 0 {
			dirToAdd = left
		} else if xDiff < 0 {
			dirToAdd = right
		} else if yDiff > 0 {
			dirToAdd = up
		} else if yDiff < 0 {
			dirToAdd = down
		}
	}

	if dirToAdd == up {
		s.tail.next = NewBody(s.tail.x, s.tail.y-cellSize, s.tail)
	} else if dirToAdd == right {
		s.tail.next = NewBody(s.tail.x+cellSize, s.tail.y, s.tail)
	} else if dirToAdd == down {
		s.tail.next = NewBody(s.tail.x, s.tail.y+cellSize, s.tail)
	} else if dirToAdd == left {
		s.tail.next = NewBody(s.tail.x-cellSize, s.tail.y, s.tail)
	}

	s.tail = s.tail.next

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

func (s *Snake) StartMove() {
	go func() {
		for {
			select {
			case <-s.stopMoving:
				return
			default:
				s.move()
			}
			time.Sleep(s.speed)
		}
	}()
}

func (s *Snake) StopMove() {
	s.stopMoving <- true
}

func (s *Snake) SpeedUp() {
	if s.speed > minSpeed {
		s.speed -= 10 * time.Millisecond
	}
}

/*
Updates the snakes direction. For now it's working off the knowledge that left is the largest
direction by numerical value and up is the smallest
*/
func (s *Snake) ChangeDirection(d direction) {
	if d >= up && d <= left {
		diff := s.heading - d
		if diff != -2 && diff != 2 {
			s.heading = d
		}
	}
}

func (s *Snake) HitWall(paddingX, paddingY int) bool {
	return s.head.x < paddingX || s.head.x >= boardSize+paddingX || s.head.y < paddingY || s.head.y >= paddingY+boardSize
}

func (s *Snake) HitItself() bool {
	body := s.head.next
	for body != nil {
		if s.head.HasCollided(body.Entity) {
			return true
		}
		body = body.next
	}
	return false
}

func (s *Snake) move() {
	nextX, nextY := s.head.x, s.head.y
	switch s.heading {
	case up:
		s.head.y -= cellSize
	case right:
		s.head.x += cellSize
	case down:
		s.head.y += cellSize
	case left:
		s.head.x -= cellSize
	default:
		fmt.Println("Something is wrong")
	}
	body := s.head.next
	for body != nil {
		tempX, tempY := body.x, body.y
		body.x = nextX
		body.y = nextY
		nextX, nextY = tempX, tempY
		body = body.next
	}
}
