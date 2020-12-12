package tests

import (
	"github.com/mrwormhole/battle-city/components"
	"github.com/mrwormhole/battle-city/core"
	"testing"
)

func TestAddEntity(t *testing.T) {
	test := core.NewEntity(core.NewVector2D(0,0), core.NewVector2D(0,0), true, "test entity")
	spriteRenderer := components.NewSpriteRenderer(test, "../assets/sprites/tank_basic_up_c0_t1.png")
	_ = test.AddComponent(spriteRenderer)

	if !test.HasComponent(spriteRenderer) {
		t.Log("Component couldn't get added to the entity")
		t.Fail()
	}
}

func TestRemoveEntity(t *testing.T) {
	test := core.NewEntity(core.NewVector2D(0,0), core.NewVector2D(0,0), true, "test entity")
	spriteRenderer := components.NewSpriteRenderer(test, "../assets/sprites/tank_basic_up_c0_t1.png")
	_ = test.AddComponent(spriteRenderer)
	_ = test.DeleteComponent(spriteRenderer)

	if test.HasComponent(spriteRenderer) {
		t.Log("Component couldn't be removed from the entity")
		t.Fail()
	}
}
