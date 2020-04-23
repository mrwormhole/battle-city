package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	ownerEntity   *Entity
	renderer      *sdl.Renderer
	texture       *sdl.Texture
	width, height int32
}

func createSpriteComponent(entity *Entity, renderer *sdl.Renderer, filename string) *Sprite {
	texture := getTextureFromPNG(renderer, filename)
	_, _, width, height, err := texture.Query()
	checkError("Renderer Query Error! ", err)

	return &Sprite{ownerEntity: entity,
		renderer: renderer,
		texture:  texture,
		width:    width,
		height:   height}
}

func (sprite Sprite) onUpdate() error {
	err := sprite.renderer.CopyEx(sprite.texture,
		&sdl.Rect{X: 0, Y: 0, W: sprite.width, H: sprite.height},
		&sdl.Rect{X: int32(sprite.ownerEntity.position.x), Y: int32(sprite.ownerEntity.position.y), W: sprite.width, H: sprite.height},
		0,
		&sdl.Point{X: sprite.width / 2, Y: sprite.height / 2},
		sdl.FLIP_NONE)
	checkError("Renderer Copy Error! ", err)

	return nil
}
