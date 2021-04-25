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

	player := NewPlayer(renderer)
	destructor = append(destructor, player)

	var enemies []BasicEnemy
	BasicEnemySize := 105

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + float64(BasicEnemySize/2)
			y := float64(j*BasicEnemySize) + float64(BasicEnemySize/2)

			enemy := NewBasicEnemy(renderer, x, y)
			enemies = append(enemies, *enemy)
			destructor = append(destructor, enemy)
		}
	}

	initBullets(renderer)

GameLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				DestroyAll(destructor)
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

		err = player.Draw(renderer)
		if err != nil {
			log.Panicln(err)
		}
		player.Update()

		for _, enemy := range enemies {
			err = enemy.Draw(renderer)
			if err != nil {
				log.Panicln(err)
			}
		}

		for _, bullet := range Bullets {
			err := bullet.Draw(renderer)
			if err != nil {
				log.Panicln(err)
			}
			bullet.Update()
		}

		renderer.Present()
	}
}

type Destructor interface {
	Destroy()
}

var destructor []Destructor

func DestroyAll(d []Destructor) {
	for index := range d {
		d[index].Destroy()
		fmt.Println("object destroyed successfully")
	}
}

func GetTextureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		log.Panicln(err)
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		log.Panicln(err)
	}
	return tex
}