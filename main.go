package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 800, 600

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkError("SDL Initialization Error: ", err)

	window, err := sdl.CreateWindow("Legend Of Zelda", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
	checkError("Window Initialization Error: ", err)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	checkError("Renderer Initialization Error: ", err)
	defer renderer.Destroy()

	player := createEntity("player")
	player.setEntityPosition(Vector2{x: SCREEN_WIDTH / 2.0, y: SCREEN_HEIGHT / 2.0})
	player.addComponent(createSpriteComponent(player, renderer, "./assets/link_blue/walk_down/0.png"))
	player.addComponent(createInputComponent(player, 0.01))
	//input then
	entities = append(entities, player)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, entity := range entities {
			if entity.active {
				err = entity.update()
				checkError("Entity Updating Error! ", err)
			}
		}

		renderer.Present()
	}
}
