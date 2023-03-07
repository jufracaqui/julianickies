package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/colors"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/input"
)

// Board represents the game board.
type Board struct {
	sizeX  int
	sizeY  int
	tiles  map[*Tile]struct{}
	blocks []*Block
}

// NewBoard generates a new Board with giving a size.
func NewBoard(sizeX, sizeY int) (*Board, error) {
	b := &Board{
		sizeX:  sizeX,
		sizeY:  sizeY,
		tiles:  map[*Tile]struct{}{},
		blocks: []*Block{NewBlock()},
	}
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			v := colors.TileTypeEmpty

			if j == 0 && i != 2 {
				v = colors.TileTypeFirstRow
			}

			b.tiles[NewTile(v, i, j)] = struct{}{}
		}
	}

	return b, nil
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := b.sizeX*tileSize + (b.sizeX+1)*tileMargin
	y := b.sizeY*tileSize + (b.sizeY+1)*tileMargin
	return x, y
}

func (b Board) MovingBlock() *Block {
	for _, block := range b.blocks {
		if !block.hasReachedBottom {
			return block
		}
	}
	return nil
}

func (b *Board) Update(i *input.Input, ignoreVerticalUpdate bool) {
	movingBlock := b.MovingBlock()
	if movingBlock != nil {
		action, hasAction := i.Dir()
		if hasAction {
			if action == input.DirRight {
				movingBlock.MoveRight()
			} else if action == input.DirLeft {
				movingBlock.MoveLeft()
			} else if action == input.Rotate {
				movingBlock.Rotate()
			}
		}
		if !ignoreVerticalUpdate {
			movingBlock.Update()
		}
	} else {
		b.blocks = append(b.blocks, NewBlock())
	}

	for tile := range b.tiles {
		if tile.y == 0 && tile.x != 2 {
			tile.value = colors.TileTypeFirstRow
		} else {
			tile.value = colors.TileTypeEmpty
		}

		for _, block := range b.blocks {
			if block.x1 == tile.x && block.y1 == tile.y {
				tile.value = block.value1
			} else if block.x2 == tile.x && block.y2 == tile.y {
				tile.value = block.value2
			}
		}
	}
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(colors.FrameColor)
	for j := 0; j < b.sizeY; j++ {
		for i := 0; i < b.sizeX; i++ {
			v := 0
			op := &ebiten.DrawImageOptions{}
			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorM.ScaleWithColor(colors.TileBackgroundColor(v))
			boardImage.DrawImage(tileImage, op)
		}
	}
	for t := range b.tiles {
		t.Draw(boardImage)
	}
}
