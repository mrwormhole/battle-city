package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrwormhole/battle-city/game"
)

func main() {
	ebiten.SetWindowSize(game.SCREEN_WIDTH*2, game.SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Battle City")
	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
