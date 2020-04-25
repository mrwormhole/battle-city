package main

import "testing"

func TestAddEntity(t *testing.T) {
	test := createEntity("test entity")
	inputComponent := createInputComponent(test, 5)
	test.addComponent(inputComponent)

	if !test.hasComponent(&Input{}) {
		t.Log("Component couldn't get added to entity")
		t.Fail()
	}
}

func TestRemoveEntity(t *testing.T) {
	test := createEntity("test entity")
	inputComponent := createInputComponent(test, 5)
	test.addComponent(inputComponent)
	test.removeComponent(&Input{})

	if test.hasComponent(&Input{}) {
		t.Log("Component couldn't be removed from entity")
		t.Fail()
	}
}
