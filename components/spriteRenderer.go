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
	options *ebiten.DrawImageOptions
}

func NewSpriteRenderer(ownerEntity *core.Entity, sourcePath string) *SpriteRenderer {
	sprite, _, err := ebitenutil.NewImageFromFile(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	return &SpriteRenderer{
		componentType: enums.SpriteRenderer,
		componentAttributes: core.NewComponentAttributes(false, true),
		ownerEntity: ownerEntity,
		options: &ebiten.DrawImageOptions{},
		sprite: sprite}
}

func (spriteRenderer *SpriteRenderer) OnUpdate() error {
	return nil
}

func (spriteRenderer *SpriteRenderer) OnDraw(screen *ebiten.Image) error {
	if spriteRenderer.componentAttributes.IsDrawable() {
		x := spriteRenderer.ownerEntity.Velocity.GetX()
		y := spriteRenderer.ownerEntity.Velocity.GetY()
		spriteRenderer.options.GeoM.Translate(x,y)
		spriteRenderer.ownerEntity.Position.SetX(spriteRenderer.options.GeoM.Element(0,2))
		spriteRenderer.ownerEntity.Position.SetY(spriteRenderer.options.GeoM.Element(1,2))

		screen.DrawImage(spriteRenderer.sprite, spriteRenderer.options)
	}
	return nil
}

func (spriteRenderer *SpriteRenderer) ComponentType() enums.ComponentType {
	return spriteRenderer.componentType
}

func (spriteRenderer *SpriteRenderer) ComponentAttributes() core.ComponentAttributes {
	return spriteRenderer.componentAttributes
}