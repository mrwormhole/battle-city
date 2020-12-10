package core

import (
	"fmt"
)

type Entity struct {
	Position   Vector2D
	isActive   bool
	tag        string
	components []Component
}

func NewEntity(Position Vector2D, isActive bool, tag string) *Entity {
	return &Entity{
		Position: Position,
		isActive: isActive,
		tag: tag,
	}
}

func (entity *Entity) GetComponent(component Component) Component {
	searchedType := component.ComponentType()
	for _, existingComponent := range entity.components {
		existingType := existingComponent.ComponentType()
		if existingType == searchedType {
			return existingComponent
		}
	}

	return nil
}

func (entity *Entity) GetComponentIndex(component Component) int {
	searchedType := component.ComponentType()
	for index, existingComponent := range entity.components {
		existingType := existingComponent.ComponentType()
		if existingType == searchedType {
			return index
		}
	}

	return -1
}

func (entity *Entity) HasComponent(component Component) bool {
	if entity.GetComponent(component) == nil {
		return false
	}
	return true
}

func (entity *Entity) AddComponent(newComponent Component) error {
	if entity.HasComponent(newComponent) {
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

func (entity *Entity) DeleteComponent(component Component) error {
	if !entity.HasComponent(component) {
		return fmt.Errorf("The component that you are trying to remove from this entity doesn't exist!")
	}

	entity.components = removeComponentByIndex(entity.components, entity.GetComponentIndex(component))
	return nil
}

func (entity *Entity) Update() error {
	for _, component := range entity.components {
		if component.IsUpdatable() {
			err := component.OnDraw()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (entity *Entity) Draw() error {
	for _, component := range entity.components {
		if component.IsDrawable() {
			err := component.OnDraw()
			if err != nil {
				return err
			}
		}
	}

	return nil
}