package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	basicEnemySize = 105
)

type BasicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func NewBasicEnemy(renderer *sdl.Renderer, x, y float64) *BasicEnemy {
	var be BasicEnemy
	be.tex = GetTextureFromBMP(renderer, "sprites/basic_enemy.bmp")
	be.x = x
	be.y = y
	return &be
}

func (be *BasicEnemy) Draw(renderer *sdl.Renderer) error {
	x := be.x - float64(basicEnemySize/2)
	y := be.y - float64(basicEnemySize/2)

	src := &sdl.Rect{
		X: 0,
		Y: 0,
		W: basicEnemySize,
		H: basicEnemySize,
	}
	dst := &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: basicEnemySize,
		H: basicEnemySize,
	}

	point := &sdl.Point{
		X: basicEnemySize / 2,
		Y: basicEnemySize / 2,
	}

	return renderer.CopyEx(
		be.tex,
		src,
		dst,
		180,
		point,
		sdl.FLIP_NONE,
	)
}

func (be *BasicEnemy) Destroy() {
	err := be.tex.Destroy()
	if err != nil {
		log.Panicln(err)
	}
}
