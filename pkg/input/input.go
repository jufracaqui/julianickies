package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Dir represents a direction.
type Dir int

const (
	DirRight Dir = iota
	DirLeft
	Rotate
)

// String returns a string representing the direction.
func (d Dir) String() string {
	switch d {
	case DirRight:
		return "Right"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

// Input represents the current key states.
type Input struct {
}

// NewInput generates a new Input object.
func NewInput() *Input {
	return &Input{}
}

// Update updates the current input states.
func (i *Input) Update() {

}

// Dir returns a currently pressed direction.
// Dir returns false if no direction key is pressed.
func (i *Input) Dir() (Dir, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return DirLeft, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return DirRight, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return Rotate, true
	}
	return 0, false
}
