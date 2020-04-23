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

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Present()
	}
}
