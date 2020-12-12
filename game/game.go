package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/components"
	"github.com/mrwormhole/battle-city/core"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 320, 240

var entity *core.Entity
var spriteRenderer *components.SpriteRenderer
var keyboardController *components.KeyboardController

type Game struct {}

func NewGame() *Game {
	entity = core.NewEntity(core.NewVector2D(0,0), core.NewVector2D(0,0), true, "player")
	spriteRenderer = components.NewSpriteRenderer(entity, "./assets/sprites/tank_basic_up_c0_t1.png")
	keyboardController = components.NewKeyboardController(entity)
	err := entity.AddComponent(spriteRenderer)
	if err != nil {
		panic(err)
	}
	err = entity.AddComponent(keyboardController)
	if err != nil {
		panic(err)
	}
	return &Game{}
}

func (g *Game) Update() error {
	err := entity.Update()
	if err != nil {
		panic(err)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("FPS: %f \nTPS: %f",
			ebiten.CurrentFPS(),
			ebiten.CurrentTPS()))

	err := entity.Draw(screen)
	if err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH,SCREEN_HEIGHT
}
