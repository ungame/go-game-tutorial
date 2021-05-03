package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	playerSpeed        = 0.07
	playerShotCooldown = time.Millisecond * 250
)

func NewPlayer(renderer *sdl.Renderer) *Element {
	var p Element
	p.position = Vector{}
	p.active = true
	sr := NewSpriteRenderer(&p, renderer, "sprites/player.bmp")
	p.AddComponent(sr)

	p.position.x = screenWidth / 2
	p.position.y = screenHeight - (sr.height / 2)

	mover := NewKeyboardMover(&p, playerSpeed)
	p.AddComponent(mover)

	shooter := NewKeyboardShooter(&p, playerShotCooldown)
	p.AddComponent(shooter)

	return &p
}
