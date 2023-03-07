package board

import (
	"math/rand"

	"gitlab.bitban.com/jfcatalina/julianickies/pkg/colors"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/config"
)

type Block struct {
	x1               int
	y1               int
	value1           int
	x2               int
	y2               int
	value2           int
	hasReachedBottom bool
}

func NewBlock() *Block {
	value1Values := []int{
		colors.TileTypeRedStar,
		colors.TileTypeBlueStar,
		colors.TileTypeGreenStar,
		colors.TileTypeYellowStar,
	}
	value2Values := []int{
		colors.TileTypeRed,
		colors.TileTypeBlue,
		colors.TileTypeGreen,
		colors.TileTypeYellow,
	}

	return &Block{
		x1:               2,
		y1:               -1,
		value1:           value1Values[rand.Intn(len(value1Values))],
		x2:               2,
		y2:               0,
		value2:           value2Values[rand.Intn(len(value2Values))],
		hasReachedBottom: false,
	}
}

func (b Block) Pos1() (int, int) {
	return b.x1, b.y1
}

func (b Block) Val1() int {
	return b.value1
}

func (b Block) Pos2() (int, int) {
	return b.x2, b.y2
}

func (b Block) Val2() int {
	return b.value2
}

func (b *Block) MoveLeft() {
	if b.y1 <= 0 {
		return
	}
	if b.x1 == 0 || b.x2 == 0 {
		return
	}

	b.x1 -= 1
	b.x2 -= 1
}

func (b *Block) MoveRight() {
	if b.y1 <= 0 {
		return
	}
	if b.x1 == config.BoardSizeX-1 || b.x2 == config.BoardSizeX-1 {
		return
	}

	b.x1 += 1
	b.x2 += 1
}

func (b *Block) Rotate() {
	x1, y1 := b.Pos1()
	// Block has just spawned
	if y1 <= 0 {
		return
	}

	x2, y2 := b.Pos2()

	// Block at the top already rotated
	if y1 == 1 && y2 == 1 {
		return
	}

	// Block is completly to the left
	if x1 == 0 && x2 == 0 {
		return
	}

	// Block is completly to the rigt
	if x1 == config.BoardSizeX-1 && x2 == config.BoardSizeX-1 {
		return
	}

	// x1, y1 are static when rotating

	if x1 > x2 {
		b.x2 = x1
		b.y2 = y1 - 1
	} else if x1 < x2 {
		b.x2 = x1
		b.y2 = y1 + 1
	} else if y1 < y2 {
		b.x2 = x1 - 1
		b.y2 = y1
	} else if y1 > y2 {
		b.x2 = x1 + 1
		b.y2 = y1
	}
}

func (b *Block) Update() {
	if b.y1 == config.BoardSizeY-1 || b.y2 == config.BoardSizeY-1 {
		b.hasReachedBottom = true
		return
	}

	// todo: check we're already at the bottom if there's other block bellow

	b.y1 += 1
	b.y2 += 1
}
