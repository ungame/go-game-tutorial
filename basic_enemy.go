package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func NewBasicEnemy(renderer *sdl.Renderer, position Vector) *Element {
	var be Element
	be.position = position
	be.rotation = 180

	sr := NewSpriteRenderer(&be, renderer, "sprites/basic_enemy.bmp")
	be.AddComponent(sr)

	be.active = true

	return &be
}
