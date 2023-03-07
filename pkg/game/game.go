package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/board"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/colors"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/config"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/input"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Game represents a game state.
type Game struct {
	input      *input.Input
	board      *board.Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{
		input: input.NewInput(),
	}
	var err error
	g.board, err = board.NewBoard(config.BoardSizeX, config.BoardSizeY)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

var updateCount = 0

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()

	hasAlreadyUpdated := false
	_, hasInput := g.input.Dir()
	if hasInput {
		g.board.Update(g.input, true)
		hasAlreadyUpdated = true
	}

	if updateCount != 60 {
		updateCount += 1
		return nil
	}

	if !hasAlreadyUpdated {
		g.board.Update(g.input, false)
	}

	updateCount = 0

	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(colors.BackgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
