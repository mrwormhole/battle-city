package main

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type Animator struct {
	ownerEntity *Entity
	renderer    *sdl.Renderer
	textures    map[string][]*sdl.Texture
	sprite      *Sprite
	features    Features

	mux                  sync.Mutex
	currentAnimationName string
	count                int
}

func createAnimatorComponent(entity *Entity, renderer *sdl.Renderer) *Animator {
	sprite := entity.getComponent(&Sprite{}).(*Sprite)
	sprite.hasAnimator = true

	return &Animator{ownerEntity: entity,
		renderer: renderer,
		sprite:   sprite}
}

func (animator *Animator) loadTextures(animationName string, filenames []string) {
	var textures map[string][]*sdl.Texture
	textures = make(map[string][]*sdl.Texture)

	// temproary
	animator.currentAnimationName = "walk_down"
	animator.count = 0

	for _, filename := range filenames {
		texture := getTextureFromPNG(animator.renderer, filename)
		textures[animationName] = append(textures[animationName], texture)
	}
	animator.textures = textures
}

func (animator *Animator) onUpdate() error {
	if animator.sprite.hasAnimator {

		animator.count += 1
		err := animator.renderer.CopyEx(animator.textures[animator.currentAnimationName][animator.count%13],
			&sdl.Rect{X: 0, Y: 0, W: animator.sprite.width, H: animator.sprite.height},
			&sdl.Rect{X: int32(animator.sprite.ownerEntity.position.x),
				Y: int32(animator.sprite.ownerEntity.position.y),
				W: animator.sprite.width * 2,
				H: animator.sprite.height * 2},
			0,
			&sdl.Point{X: animator.sprite.width, Y: animator.sprite.height},
			sdl.FLIP_NONE)
		checkError("[Animator]Renderer Copy Error! ", err)

	}

	return nil
}

func (animator *Animator) onDraw() error {
	return nil
}
