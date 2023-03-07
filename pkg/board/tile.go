package board

import (
	"fmt"
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/colors"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/fonts"
)

var (
	mplusFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Tile represents a tile information including TileData and animation states.
type Tile struct {
	value int
	x     int
	y     int
}

// Pos returns the tile's current position.
// Pos is used only at testing so far.
func (t *Tile) Pos() (int, int) {
	return t.x, t.y
}

// Value returns the tile's current value.
// Value is used only at testing so far.
func (t *Tile) Value() int {
	return t.value
}

// NewTile creates a new Tile object.
func NewTile(value int, x, y int) *Tile {
	return &Tile{
		value: value,
		x:     x,
		y:     y,
	}
}

const (
	tileSize   = 50
	tileMargin = 1
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.White)
}

// Draw draws the current tile to the given boardImage.
func (t *Tile) Draw(boardImage *ebiten.Image) {
	i, j := t.x, t.y
	v := fmt.Sprintf("%d-%d", t.x, t.y)

	if t.value == colors.TileTypeBlueStar || t.value == colors.TileTypeGreenStar || t.value == colors.TileTypeRedStar || t.value == colors.TileTypeYellowStar {
		v = "*"
	}

	op := &ebiten.DrawImageOptions{}
	x := i*tileSize + (i+1)*tileMargin
	y := j*tileSize + (j+1)*tileMargin

	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.ScaleWithColor(colors.TileBackgroundColor(t.value))

	boardImage.DrawImage(tileImage, op)

	bound, _ := font.BoundString(mplusFont, v)
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	x = x + (tileSize-w)/2
	y = y + (tileSize-h)/2 + h
	text.Draw(boardImage, v, mplusFont, x, y, colors.TileColor(t.value))
}
