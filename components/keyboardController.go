package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
)

type KeyboardController struct {
	componentType       enums.ComponentType
	componentAttributes core.ComponentAttributes
	ownerEntity         *core.Entity
}

func NewKeyboardController(ownerEntity *core.Entity) *KeyboardController {
	return &KeyboardController{
		componentType:       enums.KeyboardController,
		componentAttributes: core.NewComponentAttributes(true, false),
		ownerEntity:         ownerEntity,
	}
}

func (keyboardController *KeyboardController) OnUpdate() error {
	if keyboardController.componentAttributes.IsUpdatable() {
		keyboardController.ownerEntity.Velocity.SetX(0)
		keyboardController.ownerEntity.Velocity.SetY(0)

		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			keyboardController.ownerEntity.Velocity.SetY(-1)
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			keyboardController.ownerEntity.Velocity.SetY(+1)
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			keyboardController.ownerEntity.Velocity.SetX(-1)
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			keyboardController.ownerEntity.Velocity.SetX(1)
		}
	}
	return nil
}

func (keyboardController *KeyboardController) OnDraw(screen *ebiten.Image) error {
	return nil
}

func (keyboardController *KeyboardController) ComponentType() enums.ComponentType {
	return keyboardController.componentType
}

func (keyboardController *KeyboardController) ComponentAttributes() core.ComponentAttributes {
	return keyboardController.componentAttributes
}
