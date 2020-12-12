package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrwormhole/battle-city/enums"
)

type Component interface {
	OnUpdate() error
	OnDraw(screen *ebiten.Image) error
	ComponentType() enums.ComponentType
	ComponentAttributes() ComponentAttributes
}

type ComponentAttributes interface {
	IsUpdatable() bool
	IsDrawable() bool
}

type componentAttributes struct {
	updatable bool
	drawable  bool
}

func NewComponentAttributes(updatable bool, drawable bool) ComponentAttributes {
	return &componentAttributes{
		updatable:     updatable,
		drawable:      drawable,
	}
}

func (c componentAttributes) IsUpdatable() bool {
	return c.updatable
}

func (c componentAttributes) IsDrawable() bool {
	return c.drawable
}