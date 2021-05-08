package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
	_ "image/png"
	"log"
)

type SpriteAnimator struct {
	componentType       enums.ComponentType
	componentAttributes core.ComponentAttributes
	ownerEntity         *core.Entity
	options             *ebiten.DrawImageOptions
	sprites             []*ebiten.Image
	counter             byte
}

func NewSpriteAnimator(ownerEntity *core.Entity, sourcePaths []string) *SpriteAnimator {
	var sprites []*ebiten.Image
	for _, sourcePath := range sourcePaths {
		sprite, _, err := ebitenutil.NewImageFromFile(sourcePath)
		if err != nil {
			log.Fatal(err)
		}
		sprites = append(sprites, sprite)
	}
	return &SpriteAnimator{
		componentType:       enums.SpriteAnimator,
		componentAttributes: core.NewComponentAttributes(true, true),
		ownerEntity:         ownerEntity,
		options:             &ebiten.DrawImageOptions{},
		sprites:             sprites,
	}
}

func (spriteAnimator *SpriteAnimator) OnUpdate() error {
	if !spriteAnimator.ownerEntity.HasComponent(enums.SpriteRenderer) {
		spriteAnimator.counter++
	}
	return nil
}

func (spriteAnimator *SpriteAnimator) OnDraw(screen *ebiten.Image) error {
	if spriteAnimator.componentAttributes.IsDrawable() && !spriteAnimator.ownerEntity.HasComponent(enums.SpriteRenderer) {
		x := spriteAnimator.ownerEntity.Velocity.X()
		y := spriteAnimator.ownerEntity.Velocity.GetY()
		spriteAnimator.options.GeoM.Translate(x, y)
		spriteAnimator.ownerEntity.Position.SetX(spriteAnimator.options.GeoM.Element(0, 2))
		spriteAnimator.ownerEntity.Position.SetY(spriteAnimator.options.GeoM.Element(1, 2))

		selectedIndex := (spriteAnimator.counter / 10) % byte(len(spriteAnimator.sprites))

		screen.DrawImage(spriteAnimator.sprites[selectedIndex], spriteAnimator.options)
	}
	return nil
}

func (spriteAnimator *SpriteAnimator) ComponentType() enums.ComponentType {
	return spriteAnimator.componentType
}

func (spriteAnimator *SpriteAnimator) ComponentAttributes() core.ComponentAttributes {
	return spriteAnimator.componentAttributes
}
