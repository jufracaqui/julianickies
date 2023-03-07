package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/config"
	"gitlab.bitban.com/jfcatalina/julianickies/pkg/game"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Julianickies")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
