package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrwormhole/battle-city/enums"
)

type Entity struct {
	Position   Vector2D
	Velocity   Vector2D
	IsActive   bool
	Tag        string
	components []Component
}

func NewEntity(position Vector2D, velocity Vector2D, isActive bool, tag string) *Entity {
	return &Entity{
		Position: position,
		Velocity: velocity,
		IsActive: isActive,
		Tag:      tag,
	}
}

func (entity *Entity) GetComponent(componentType enums.ComponentType) Component {
	for _, existingComponent := range entity.components {
		if existingComponent.ComponentType() == componentType {
			return existingComponent
		}
	}
	return nil
}

func (entity *Entity) GetComponentIndex(componentType enums.ComponentType) int {
	for index, existingComponent := range entity.components {
		if existingComponent.ComponentType() == componentType {
			return index
		}
	}
	return -1
}

func (entity *Entity) HasComponent(componentType enums.ComponentType) bool {
	return entity.GetComponent(componentType) != nil
}

func (entity *Entity) AddComponent(newComponent Component) error {
	if entity.HasComponent(newComponent.ComponentType()) {
		return fmt.Errorf("The same component already exists in this entity!")
	}
	entity.components = append(entity.components, newComponent)
	return nil
}

func removeComponentByIndex(components []Component, index int) []Component {
	newSize := len(components) - 1
	components[index] = components[newSize]
	return components[:newSize]
}

func (entity *Entity) RemoveComponent(component Component) error {
	if !entity.HasComponent(component.ComponentType()) {
		return fmt.Errorf("The component that you are trying to remove from this entity doesn't exist!")
	}
	entity.components = removeComponentByIndex(entity.components, entity.GetComponentIndex(component.ComponentType()))
	return nil
}

func (entity *Entity) Update() error {
	for _, component := range entity.components {
		if component.ComponentAttributes().IsUpdatable() {
			err := component.OnUpdate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (entity *Entity) Draw(screen *ebiten.Image) error {
	for _, component := range entity.components {
		if component.ComponentAttributes().IsDrawable() {
			err := component.OnDraw(screen)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
