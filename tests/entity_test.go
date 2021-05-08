package tests

import (
	"github.com/mrwormhole/battle-city/components"
	"github.com/mrwormhole/battle-city/core"
	"github.com/mrwormhole/battle-city/enums"
	"testing"
)

//TODO: more tests need to be added for code cov

func TestAddToEntity(t *testing.T) {
	test := core.NewEntity(core.NewVector2D(0,0), core.NewVector2D(0,0), true, "test entity")
	spriteRenderer := components.NewSpriteRenderer(test, "../assets/sprites/tank_basic_up_c0_t1.png")
	err := test.AddComponent(spriteRenderer)
	if err != nil {
		t.Log("Component couldn't get added to the entity")
		t.Fail()
	}

	if !test.HasComponent(enums.SpriteRenderer) {
		t.Log("Sprite renderer doesn't exist in the entity")
		t.Fail()
	}
}

func TestRemoveFromEntity(t *testing.T) {
	test := core.NewEntity(core.NewVector2D(0,0), core.NewVector2D(0,0), true, "test entity")
	spriteRenderer := components.NewSpriteRenderer(test, "../assets/sprites/tank_basic_up_c0_t1.png")
	err := test.AddComponent(spriteRenderer)
	if err != nil {
		t.Log("Component couldn't get added to the entity")
		t.Fail()
	}
	err = test.RemoveComponent(spriteRenderer)
	if err != nil {
		t.Log("Component couldn't get deleted from the entity")
		t.Fail()
	}

	if test.HasComponent(enums.SpriteRenderer) {
		t.Log("Sprite renderer exists in the entity")
		t.Fail()
	}
}
