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

	vtb := NewVulnerableToBullets(&be)
	be.AddComponent(vtb)

	be.active = true

	col := Circle{}
	col.center = be.position
	col.radius = 38

	be.collisions = append(be.collisions, col)

	return &be
}
