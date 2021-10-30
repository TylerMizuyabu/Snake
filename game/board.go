package game

const (
	minWidth = 30
	minHeight = 30
)

type Board struct {
	originalWidth float64
	originalHeight float64
	width float64
	height float64
}

func NewBoard(width, height float64) *Board {
	if width < minWidth {
		width = minWidth
	}
	if height < minHeight {
		height = minHeight
	}
	return &Board {
		originalWidth: width,
		originalHeight: height,
		width: width,
		height: height,
	}
}

func (b *Board) ReduceSize() {
	b.width -= 1
	b.height -=1
}

func (b *Board) GetSize() (float64, float64) {
	return b.width, b.height
}

func (b *Board) Reset() {
	b.width = b.originalWidth
	b.height = b.originalHeight
}
