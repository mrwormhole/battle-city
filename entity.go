package main

import (
	"reflect"
)

var entities []*Entity

type Direction byte

const (
	north Direction = iota
	east  Direction = iota
	south Direction = iota
	west  Direction = iota
)

type Vector2 struct {
	x, y float64
}

type Features struct {
	updateble bool
	drawable  bool
}

type Component interface {
	onUpdate() error
	onDraw() error
	getFeatures() Features
}

type Entity struct {
	position   Vector2
	rotation   Direction
	active     bool
	tag        string
	components []Component
}

func createEntity(tag string) *Entity {
	return &Entity{active: true, tag: tag}
}

func (entity *Entity) setEntityPosition(position Vector2) {
	entity.position = position
}

func (entity *Entity) setEntityRotation(rotation Direction) {
	entity.rotation = rotation
}

func (entity *Entity) setActive(active bool) {
	entity.active = active
}

func (entity *Entity) setTag(tag string) {
	entity.tag = tag
}

func (entity *Entity) getComponent(component Component) Component {
	searchedType := reflect.TypeOf(component)
	for _, existingComponent := range entity.components {
		existingType := reflect.TypeOf(existingComponent)
		if existingType == searchedType {
			return existingComponent
		}
	}

	return nil
}

func (entity *Entity) getComponentIndex(component Component) int {
	searchedType := reflect.TypeOf(component)
	for index, existingComponent := range entity.components {
		existingType := reflect.TypeOf(existingComponent)
		if existingType == searchedType {
			return index
		}
	}

	return -1
}

func (entity *Entity) hasComponent(component Component) bool {
	if entity.getComponent(component) == nil {
		return false
	}
	return true
}

func (entity *Entity) addComponent(newComponent Component) {
	if entity.hasComponent(newComponent) {
		return
	}
	entity.components = append(entity.components, newComponent)
}

func remove(components []Component, index int) []Component {
	newSize := len(components) - 1
	components[index] = components[newSize]
	return components[:newSize]
}

func (entity *Entity) removeComponent(component Component) {
	if !entity.hasComponent(component) {
		return
	}
	entity.components = remove(entity.components, entity.getComponentIndex(component))
}

func (entity *Entity) update() error {
	for _, component := range entity.components {
		if component.getFeatures().updateble {
			err := component.onUpdate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (entity *Entity) draw() error {
	for _, component := range entity.components {
		if component.getFeatures().drawable {
			err := component.onDraw()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
