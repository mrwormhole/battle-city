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
		uY := lineCollider.ownerEntity.Position.GetY() - lineCollider.startPosition.GetY()
		lineCollider.startPosition.SetX(lineCollider.ownerEntity.Position.GetX())
		lineCollider.startPosition.SetY(lineCollider.ownerEntity.Position.GetY())
		lineCollider.endPosition.SetX(lineCollider.endPosition.GetX() + uX)
		lineCollider.endPosition.SetY(lineCollider.endPosition.GetY() + uY)

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
			lineCollider.startPosition.GetX(),
			lineCollider.startPosition.GetY(),
			lineCollider.endPosition.GetX(),
			lineCollider.endPosition.GetY(),
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
	uA := ((otherLineCollider.endPosition.GetX()-otherLineCollider.startPosition.GetX()) *
		(lineCollider.startPosition.GetY() - otherLineCollider.startPosition.GetY()) -
		(otherLineCollider.endPosition.GetY() - otherLineCollider.startPosition.GetY()) *
		(lineCollider.startPosition.GetX() - otherLineCollider.startPosition.GetX())) /
		((otherLineCollider.endPosition.GetY() - otherLineCollider.startPosition.GetY()) *
		(lineCollider.endPosition.GetX() - lineCollider.startPosition.GetX()) -
		(otherLineCollider.endPosition.GetX() - otherLineCollider.startPosition.GetX()) *
		(lineCollider.endPosition.GetY()- lineCollider.startPosition.GetY()))

	uB := ((lineCollider.endPosition.GetX() - lineCollider.startPosition.GetX()) *
		(lineCollider.startPosition.GetY() - otherLineCollider.startPosition.GetY()) -
		(lineCollider.endPosition.GetY() - lineCollider.startPosition.GetY()) *
		(lineCollider.startPosition.GetX() - otherLineCollider.startPosition.GetX())) /
		((otherLineCollider.endPosition.GetY() - otherLineCollider.startPosition.GetY()) *
		(lineCollider.endPosition.GetX() - lineCollider.startPosition.GetX()) -
		(otherLineCollider.endPosition.GetX() - otherLineCollider.startPosition.GetX()) *
		(lineCollider.endPosition.GetY() - lineCollider.startPosition.GetY()))

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		return true
	}
	return false
}

func (lineCollider *LineCollider) CollidesWithBox(otherBoxCollider *BoxCollider) bool {
	left := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.GetX(), otherBoxCollider.position.GetY()),
		endPosition: core.NewVector2D(otherBoxCollider.position.GetX(), otherBoxCollider.position.GetY() + otherBoxCollider.height),
	})
	right := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.GetX() + otherBoxCollider.width, otherBoxCollider.position.GetY()),
		endPosition: core.NewVector2D(otherBoxCollider.position.GetX() + otherBoxCollider.width, otherBoxCollider.position.GetY() + otherBoxCollider.height),
	})
	top := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.GetX(), otherBoxCollider.position.GetY()),
		endPosition: core.NewVector2D(otherBoxCollider.position.GetX() + otherBoxCollider.width, otherBoxCollider.position.GetY()),
	})
	bottom := lineCollider.collidesWithLine(&LineCollider{
		startPosition: core.NewVector2D(otherBoxCollider.position.GetX(), otherBoxCollider.position.GetY() + otherBoxCollider.height),
		endPosition: core.NewVector2D(otherBoxCollider.position.GetX() + otherBoxCollider.width, otherBoxCollider.position.GetY() + otherBoxCollider.height),
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
