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

var Elements []*Element

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

	player := NewPlayer(renderer)
	Elements = append(Elements, player)
	destructor = append(destructor, player)

	BasicEnemySize := 105

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {

			x := (float64(i)/5)*screenWidth + float64(BasicEnemySize/2)
			y := float64(j*BasicEnemySize) + float64(BasicEnemySize/2)

			enemy := NewBasicEnemy(renderer, Vector{x, y})
			Elements = append(Elements, enemy)
			destructor = append(destructor, enemy)
		}
	}

	InitBullets(renderer, 30)
	Elements = append(Elements, Bullets...)

GameLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				DestroyAll(Elements)
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

		for _, element := range Elements {
			if element.active {
				element.Draw(renderer)
				element.Update()
			}
		}

		err = CheckCollisions(Elements)
		if err != nil {
			log.Panicln(err)
		}

		renderer.Present()
	}
}

type Destructor interface {
	Destroy()
}

var destructor []Destructor

func DestroyAll(d []*Element) {
	for index := range d {
		d[index].Destroy()
	}
	fmt.Println("objects destroyed successfully:", len(d))
}
