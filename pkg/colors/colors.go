package colors

import (
	"image/color"
)

var (
	BackgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	FrameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

const (
	TileTypeEmpty int = iota
	TileTypeFirstRow
	TileTypeRed
	TileTypeBlue
	TileTypeGreen
	TileTypeYellow
	TileTypeRedStar
	TileTypeBlueStar
	TileTypeGreenStar
	TileTypeYellowStar
)

func TileColor(value int) color.Color {
	return color.RGBA{0x77, 0x6e, 0x65, 0xff}
}

func TileBackgroundColor(value int) color.Color {
	switch value {
	case TileTypeFirstRow:
		return color.RGBA{0x00, 0x00, 0x00, 0xff}
	case TileTypeRed:
		return color.RGBA{0xff, 0x00, 0x00, 0xaa}
	case TileTypeBlue:
		return color.RGBA{0x00, 0x00, 0xff, 0xaa}
	case TileTypeGreen:
		return color.RGBA{0x00, 0xff, 0x00, 0xaa}
	case TileTypeYellow:
		return color.RGBA{0xff, 0xce, 0x36, 0xaa}
	case TileTypeRedStar:
		return color.RGBA{0xff, 0x00, 0x00, 0xff}
	case TileTypeBlueStar:
		return color.RGBA{0x00, 0x00, 0xff, 0xff}
	case TileTypeGreenStar:
		return color.RGBA{0x00, 0xff, 0x00, 0xff}
	case TileTypeYellowStar:
		return color.RGBA{0xff, 0xce, 0x36, 0xff}
	}

	return color.NRGBA{0xff, 0xff, 0xff, 0xff}
}
