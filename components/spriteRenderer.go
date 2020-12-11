package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
	_ "image/png"
	"log"
)

type SpriteRenderer struct {
	componentType enums.ComponentType
	componentAttributes core.ComponentAttributes
	ownerEntity *core.Entity
	sprite *ebiten.Image
}

func NewSpriteRenderer(ownerEntity *core.Entity,sourcePath string) *SpriteRenderer {
	sprite, _, err := ebitenutil.NewImageFromFile(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	return &SpriteRenderer{
		componentAttributes: core.NewComponentAttributes(true, true),
		componentType: enums.SpriteRenderer,
		ownerEntity: ownerEntity,
		sprite: sprite}
}

func (spriteRenderer *SpriteRenderer) OnUpdate() error {
	if spriteRenderer.componentAttributes.IsUpdatable() {

	}
	return nil
}

func (spriteRenderer *SpriteRenderer) OnDraw() error {
	if spriteRenderer.componentAttributes.IsDrawable() {

	}
	return nil
}