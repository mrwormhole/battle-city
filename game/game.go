package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/components"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 320, 240

var player *core.Entity
var dummy *core.Entity
//var spriteRenderer *components.SpriteRenderer
var spriteAnimator *components.SpriteAnimator
var keyboardController *components.KeyboardController
var boxCollider *components.BoxCollider

type Game struct{}

func NewGame() *Game {
	player = core.NewEntity(core.NewVector2D(0, 0), core.NewVector2D(0, 0), true, "player")
	//spriteRenderer = components.NewSpriteRenderer(player, "./assets/sprites/tank_basic_up_c0_t1.png")
	spriteAnimator = components.NewSpriteAnimator(player, []string{"./assets/sprites/tank_basic_down_c0_t1.png",
		"./assets/sprites/tank_basic_right_c0_t1.png",
		"./assets/sprites/tank_basic_up_c0_t1.png",
		"./assets/sprites/tank_basic_left_c0_t1.png"})
	keyboardController = components.NewKeyboardController(player)
	boxCollider = components.NewBoxCollider(player, 32, 32)
	//err := player.AddComponent(spriteRenderer)
	//if err != nil {
	//	panic(err)
	//}
	// Note: spriteRenderer shouldn't exist in the same component with spriteAnimator
	err := player.AddComponent(spriteAnimator)
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

	dummy = core.NewEntity(core.NewVector2D(SCREEN_WIDTH/2, SCREEN_HEIGHT/2), core.NewVector2D(0, 0), true, "dummy")
	lineCollider := components.NewLineCollider(dummy, dummy.Position, core.NewVector2D(dummy.Position.X()+50, dummy.Position.Y()+50))
	err = dummy.AddComponent(lineCollider)
	if err != nil {
		panic(err)
	}

	playerBoxCollider := player.GetComponent(enums.BoxCollider).(*components.BoxCollider)
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
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
