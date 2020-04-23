package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Collider struct {
	ownerEntity   *Entity
	renderer      *sdl.Renderer
	position      Vector2
	width, height float64
}

func createColliderComponent(entity *Entity, renderer *sdl.Renderer, position Vector2, width float64, height float64) *Collider {
	return &Collider{ownerEntity: entity,
		renderer: renderer,
		position: position,
		width:    width,
		height:   height}
}

func (collider *Collider) isColliding(otherCollider *Collider) bool {
	if collider.position.x < otherCollider.position.x+otherCollider.width &&
		collider.position.x+collider.width > otherCollider.position.x &&
		collider.position.y < otherCollider.position.y+otherCollider.height &&
		collider.position.y+collider.height > otherCollider.position.y {
		return true
	}
	return false
}

func (collider *Collider) paintCollider() {
	collider.renderer.SetDrawColor(0, 255, 0, 255)
	collider.renderer.DrawRect(&sdl.Rect{X: int32(collider.position.x),
		Y: int32(collider.position.y),
		W: int32(collider.width),
		H: int32(collider.height)})
}

func (collider *Collider) onUpdate() error {
	collider.position.x = collider.ownerEntity.position.x
	collider.position.y = collider.ownerEntity.position.y
	collider.paintCollider()

	for _, otherEntity := range entities {
		if otherEntity == collider.ownerEntity {
			//dont collide with yourself
			continue
		}
		if otherEntity.active && otherEntity.hasComponent(&Collider{}) && otherEntity.tag == "dummy" {
			if collider.isColliding(otherEntity.getComponent(&Collider{}).(*Collider)) {
				fmt.Println("Colliding with dummy")
			}
		}

	}

	return nil
}
