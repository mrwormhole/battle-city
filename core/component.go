package core

import "github.com/mrwormhole/battle-city/enums"

type Component interface {
	OnUpdate() error
	OnDraw() error
	IsUpdatable() bool
	IsDrawable() bool
	ComponentType() enums.ComponentType
}

type ComponentFeatures struct {
	componentType enums.ComponentType
	updatable bool
	drawable  bool
}

func NewComponentFeatures(componentType enums.ComponentType, updatable bool, drawable bool) *ComponentFeatures {
	return &ComponentFeatures{
		componentType: componentType,
		updatable:     updatable,
		drawable:      drawable,
	}
}

