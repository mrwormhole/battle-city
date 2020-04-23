package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 800, 600
const TARGET_FPS = 60
const FRAME_DELAY = 1000 / TARGET_FPS

var timer *Timer

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
	player_position := Vector2{x: SCREEN_WIDTH / 2.0, y: SCREEN_HEIGHT / 2.0}
	player.setEntityPosition(player_position)
	player.addComponent(createSpriteComponent(player, renderer, "./assets/link_blue/walk_down/0.png"))
	player.addComponent(createInputComponent(player, 5))
	animatorComponent := createAnimatorComponent(player, renderer)
	animatorComponent.loadTextures("walk_down", []string{
		"./assets/link_blue/walk_down/0.png",
		"./assets/link_blue/walk_down/1.png",
		"./assets/link_blue/walk_down/2.png",
		"./assets/link_blue/walk_down/3.png",
		"./assets/link_blue/walk_down/4.png",
		"./assets/link_blue/walk_down/5.png",
		"./assets/link_blue/walk_down/6.png",
		"./assets/link_blue/walk_down/7.png",
		"./assets/link_blue/walk_down/8.png",
		"./assets/link_blue/walk_down/9.png",
		"./assets/link_blue/walk_down/10.png",
		"./assets/link_blue/walk_down/11.png",
		"./assets/link_blue/walk_down/12.png"})
	player.addComponent(animatorComponent)
	player.addComponent(createColliderComponent(player, renderer, player_position, 36, 44))

	entities = append(entities, player)

	dummy := createEntity("dummy")
	dummy_position := Vector2{x: 300, y: 300}
	dummy.setEntityPosition(dummy_position)
	dummy.addComponent(createColliderComponent(dummy, renderer, dummy_position, 100, 100))

	entities = append(entities, dummy)

	timer = createTimer(FRAME_DELAY)

	for {
		timer.start()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}

		}
		if timer.tick() {
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
}
