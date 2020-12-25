package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/components"
	"github.com/mrwormhole/battle-city/core"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 320, 240

var player *core.Entity
var dummy *core.Entity
var spriteRenderer *components.SpriteRenderer
var keyboardController *components.KeyboardController
var boxCollider *components.BoxCollider

type Game struct {}

func NewGame() *Game {
	player = core.NewEntity(core.NewVector2D(0,0), core.NewVector2D(0,0), true, "player")
	spriteRenderer = components.NewSpriteRenderer(player, "./assets/sprites/tank_basic_up_c0_t1.png")
	keyboardController = components.NewKeyboardController(player)
	boxCollider = components.NewBoxCollider(player, 32,32)
	err := player.AddComponent(spriteRenderer)
	if err != nil {
		panic(err)
	}
	err = player.AddComponent(keyboardController)
	if err != nil {
		panic(err)
	}
	err = player.AddComponent(boxCollider)
	if err != nil {
		panic(err)
	}

	dummy = core.NewEntity(core.NewVector2D(SCREEN_WIDTH/2, SCREEN_HEIGHT/2), core.NewVector2D(0,0), true, "dummy")
	boxCollider2 := components.NewBoxCollider(dummy, 32, 32)
	err = dummy.AddComponent(boxCollider2)
	if err != nil {
		panic(err)
	}

	playerBoxCollider := player.GetComponent(boxCollider).(*components.BoxCollider)
	playerBoxCollider.AddEntityToCollisionPool(dummy)

	return &Game{}
}

func (g *Game) Update() error {
	err := player.Update()
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

	err := player.Draw(screen)
	if err != nil {
		panic(err)
	}

	err = dummy.Draw(screen)
	if err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH,SCREEN_HEIGHT
}
