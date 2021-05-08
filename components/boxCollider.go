package components

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
	"image/color"
)

type BoxCollider struct {
	componentType       enums.ComponentType
	componentAttributes core.ComponentAttributes
	ownerEntity         *core.Entity
	entityCollisionPool []*core.Entity
	position            core.Vector2D
	width               float64
	height              float64
}

func NewBoxCollider(ownerEntity *core.Entity, width float64, height float64) *BoxCollider {
	return &BoxCollider{
		componentType:       enums.BoxCollider,
		componentAttributes: core.NewComponentAttributes(true, true),
		ownerEntity:         ownerEntity,
		position:            ownerEntity.Position,
		width:               width,
		height:              height,
	}
}

func (boxCollider *BoxCollider) OnUpdate() error {
	if boxCollider.componentAttributes.IsUpdatable() {
		boxCollider.position.SetX(boxCollider.ownerEntity.Position.X())
		boxCollider.position.SetY(boxCollider.ownerEntity.Position.Y())

		for _, otherEntity := range boxCollider.entityCollisionPool {
			if otherEntity.IsActive {
				if otherEntity.HasComponent(enums.BoxCollider) && otherEntity.Tag == "dummy" {
					if boxCollider.collidesWithBox(otherEntity.GetComponent(enums.BoxCollider).(*BoxCollider)) {
						fmt.Println("[BOX-BOX]TEST COLLISION FROM BOX COLLIDER COMPONENT")
					}
				}
				if otherEntity.HasComponent(enums.LineCollider) && otherEntity.Tag == "dummy" {
					if otherEntity.GetComponent(enums.LineCollider).(*LineCollider).CollidesWithBox(boxCollider) {
						fmt.Println("[BOX-LINE]TEST COLLISION FROM BOX COLLIDER COMPONENT")
					}
				}
			}
		}
	}
	return nil
}

func (boxCollider *BoxCollider) OnDraw(screen *ebiten.Image) error {
	if boxCollider.componentAttributes.IsDrawable() {
		ebitenutil.DrawRect(screen,
			boxCollider.position.X(),
			boxCollider.position.Y(),
			boxCollider.width,
			boxCollider.height,
			color.RGBA{
				R: 255,
				G: 0,
				B: 0,
				A: 155,
			})
	}
	return nil
}

func (boxCollider *BoxCollider) ComponentType() enums.ComponentType {
	return boxCollider.componentType
}

func (boxCollider *BoxCollider) ComponentAttributes() core.ComponentAttributes {
	return boxCollider.componentAttributes
}

func (boxCollider *BoxCollider) collidesWithBox(otherBoxCollider *BoxCollider) bool {
	if boxCollider.position.X() < otherBoxCollider.position.X()+otherBoxCollider.width &&
		boxCollider.position.X()+boxCollider.width > otherBoxCollider.position.X() &&
		boxCollider.position.Y() < otherBoxCollider.position.Y()+otherBoxCollider.height &&
		boxCollider.position.Y()+boxCollider.height > otherBoxCollider.position.Y() {
		return true
	}
	return false
}

func (boxCollider *BoxCollider) AddEntityToCollisionPool(entity *core.Entity) {
	if entity == boxCollider.ownerEntity {
		return
	}
	boxCollider.entityCollisionPool = append(boxCollider.entityCollisionPool, entity)
}

func (boxCollider *BoxCollider) RemoveEntityFromCollisionPool(entity *core.Entity) {
	i := 0
	if len(boxCollider.entityCollisionPool) == 0 {
		return
	}
	for _, otherEntity := range boxCollider.entityCollisionPool {
		if otherEntity == entity {
			break
		}
		i++
	}
	j := len(boxCollider.entityCollisionPool) - 1
	boxCollider.entityCollisionPool[i] = boxCollider.entityCollisionPool[j]
	boxCollider.entityCollisionPool[j] = nil
	boxCollider.entityCollisionPool = boxCollider.entityCollisionPool[:j]
}
