package main

import (
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type direction byte

const (
	North direction = iota
	East  direction = iota
	South direction = iota
	West  direction = iota
)

type vector2 struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type entity struct {
	position   vector2
	rotation   direction
	active     bool
	tag        string
	components []component
}

func createEntity(tag string) *entity {
	return &entity{tag: tag}
}

func (entity *entity) setEntityPosition(position vector2) {
	entity.position = position
}

func (entity *entity) setEntityRotation(rotation direction) {
	entity.rotation = rotation
}

func (entity *entity) setActive(active bool) {
	entity.active = active
}

func (entity *entity) setTag(tag string) {
	entity.tag = tag
}

func (entity *entity) getComponent(component component) component {
	searchedType := reflect.TypeOf(component)
	for _, existingComponent := range entity.components {
		existingType := reflect.TypeOf(existingComponent)
		if existingType == searchedType {
			return existingComponent
		}
	}

	return nil
}

func (entity *entity) hasComponent(component component) bool {
	if entity.getComponent(component) == nil {
		return false
	}
	return true
}

func (entity *entity) addComponent(newComponent component) {
	if entity.hasComponent(newComponent) {
		return
	}
	entity.components = append(entity.components, newComponent)
}
