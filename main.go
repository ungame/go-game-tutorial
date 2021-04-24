package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Panicln(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Game Tutorial",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		log.Panicln(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Panicln(err)
	}
	defer renderer.Destroy()

	player, err := NewPlayer(renderer)
	if err != nil {
		log.Panicln(err)
	}
	defer player.Destroy()

GameLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit...")
				break GameLoop
			}
		}

		err = renderer.SetDrawColor(0, 0, 0, 0)
		if err != nil {
			log.Panicln(err)
		}
		err = renderer.Clear()
		if err != nil {
			log.Panicln(err)
		}

		err  = player.Draw(renderer)
		if err != nil {
			log.Panicln(err)
		}

		renderer.Present()
	}
}
