package components

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
	"image/color"
)

type LineCollider struct {
	componentType enums.ComponentType
	componentAttributes core.ComponentAttributes
	ownerEntity *core.Entity
	entityCollisionPool []*core.Entity
	startPosition core.Vector2D
	endPosition core.Vector2D
}

func NewLineCollider(ownerEntity *core.Entity, startPosition core.Vector2D, endPosition core.Vector2D) *LineCollider {
	return &LineCollider{
		componentType:       enums.LineCollider,
		componentAttributes: core.NewComponentAttributes(true, true),
		ownerEntity:         ownerEntity,
		startPosition: startPosition,
		endPosition: endPosition,
	}
}

func (lineCollider *LineCollider) OnUpdate() error {
	if lineCollider.componentAttributes.IsUpdatable() {
		uX := lineCollider.ownerEntity.Position.X() - lineCollider.startPosition.X()
		uY := lineCollider.ownerEntity.Position.Y() - lineCollider.startPosition.Y()
		lineCollider.startPosition.SetX(lineCollider.ownerEntity.Position.X())
		lineCollider.startPosition.SetY(lineCollider.ownerEntity.Position.Y())
		lineCollider.endPosition.SetX(lineCollider.endPosition.X() + uX)
		lineCollider.endPosition.SetY(lineCollider.endPosition.Y() + uY)

		for _, otherEntity := range lineCollider.entityCollisionPool {
			if otherEntity.IsActive {
				if otherEntity.HasComponent(enums.LineCollider) && otherEntity.Tag == "dummy" {
					if lineCollider.collidesWithLine(otherEntity.GetComponent(enums.LineCollider).(*LineCollider)) {
						fmt.Println("[LINE-LINE]TEST COLLISION FROM LINE COLLIDER COMPONENT")
					}
				}
				if otherEntity.HasComponent(enums.BoxCollider) && otherEntity.Tag == "dummy" {
					if lineCollider.CollidesWithBox(otherEntity.GetComponent(enums.BoxCollider).(*BoxCollider)) {
						fmt.Println("[LINE-BOX]TEST COLLISION FROM LINE COLLIDER COMPONENT")
					}
				}
			}
		}
	}
	return nil
}

func (lineCollider *LineCollider) OnDraw(screen *ebiten.Image) error {
	if lineCollider.componentAttributes.IsDrawable() {
		ebitenutil.DrawLine(screen,
			lineCollider.startPosition.X(),
			lineCollider.startPosition.Y(),
			lineCollider.endPosition.X(),
			lineCollider.endPosition.Y(),
			color.RGBA{
				R: 255,
				G: 0,
				B: 0,
				A: 155,
			})
	}
	return nil
}

func (lineCollider *LineCollider) ComponentType() enums.ComponentType {
	return lineCollider.componentType
}

func (lineCollider *LineCollider) ComponentAttributes() core.ComponentAttributes {
	return lineCollider.componentAttributes
}

func (lineCollider *LineCollider) collidesWithLine(otherLineCollider *LineCollider) bool {
	uA := ((otherLineCollider.endPosition.X()-otherLineCollider.startPosition.X()) *
		(lineCollider.startPosition.Y() - otherLineCollider.startPosition.Y()) -
		(otherLineCollider.endPosition.Y() - otherLineCollider.startPosition.Y()) *
		(lineCollider.startPosition.X() - otherLineCollider.startPosition.X())) /
		((otherLineCollider.endPosition.Y() - otherLineCollider.startPosition.Y()) *
		(lineCollider.endPosition.X() - lineCollider.startPosition.X()) -
		(otherLineCollider.endPosition.X() - otherLineCollider.startPosition.X()) *
		(lineCollider.endPosition.Y()- lineCollider.startPosition.Y()))

	uB := ((lineCollider.endPosition.X() - lineCollider.startPosition.X()) *
		(lineCollider.startPosition.Y() - otherLineCollider.startPosition.Y()) -
		(lineCollider.endPosition.Y() - lineCollider.startPosition.Y()) *
		(lineCollider.startPosition.X() - otherLineCollider.startPosition.X())) /
		((otherLineCollider.endPosition.Y() - otherLineCollider.startPosition.Y()) *
		(lineCollider.endPosition.X() - lineCollider.startPosition.X()) -
		(otherLineCollider.endPosition.X() - otherLineCollider.startPosition.X()) *
		(lineCollider.endPosition.Y() - lineCollider.startPosition.Y()))

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		return true
	}
	return false
}

func (lineCollider *LineCollider) CollidesWithBox(otherBoxCollider *BoxCollider) bool {
	left := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.X(), otherBoxCollider.position.Y()),
		endPosition: core.NewVector2D(otherBoxCollider.position.X(), otherBoxCollider.position.Y() + otherBoxCollider.height),
	})
	right := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.X() + otherBoxCollider.width, otherBoxCollider.position.Y()),
		endPosition: core.NewVector2D(otherBoxCollider.position.X() + otherBoxCollider.width, otherBoxCollider.position.Y() + otherBoxCollider.height),
	})
	top := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.X(), otherBoxCollider.position.Y()),
		endPosition: core.NewVector2D(otherBoxCollider.position.X() + otherBoxCollider.width, otherBoxCollider.position.Y()),
	})
	bottom := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.X(), otherBoxCollider.position.Y() + otherBoxCollider.height),
		endPosition: core.NewVector2D(otherBoxCollider.position.X() + otherBoxCollider.width, otherBoxCollider.position.Y() + otherBoxCollider.height),
	})

	if left || right || top || bottom {
		return true
	}
	return false
}

func (lineCollider *LineCollider) AddEntityToCollisionPool(entity *core.Entity) {
	if entity == lineCollider.ownerEntity {
		return
	}
	lineCollider.entityCollisionPool = append(lineCollider.entityCollisionPool, entity)
}

func (lineCollider *LineCollider) RemoveEntityFromCollisionPool(entity *core.Entity) {
	if len(lineCollider.entityCollisionPool) == 0 {
		return
	}

	i := 0
	var otherEntity *core.Entity
	for i, otherEntity = range lineCollider.entityCollisionPool {
		if otherEntity == entity {
			break
		}
	}
	j := len(lineCollider.entityCollisionPool) - 1
	lineCollider.entityCollisionPool[i] = lineCollider.entityCollisionPool[j]
	lineCollider.entityCollisionPool[j] = nil
	lineCollider.entityCollisionPool = lineCollider.entityCollisionPool[:j]
}
