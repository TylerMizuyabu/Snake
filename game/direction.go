package game

type direction int

const (
	up direction = iota
	right
	down
	left
)

func oppositeDirection(d direction) direction {
	switch d {
	case up:
		return down
	case right:
		return left
	case down:
		return up
	case left:
		return right
	}
	return 0
}
