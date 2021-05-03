package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

type KeyboardMover struct {
	container *Element
	speed     float64
	sr        *SpriteRenderer
}

func NewKeyboardMover(container *Element, speed float64) *KeyboardMover {
	var km KeyboardMover
	km.container = container
	km.speed = speed
	km.sr = container.GetComponent(&SpriteRenderer{}).(*SpriteRenderer)
	return &km
}

func (km *KeyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	container := km.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if container.position.x-(km.sr.width/2) > 0 {
			container.position.x -= km.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if container.position.x+(km.sr.width/2) < screenWidth {
			container.position.x += km.speed
		}
	}
	return nil
}

func (km *KeyboardMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

type KeyboardShooter struct {
	container *Element
	cooldown  time.Duration
	lastShot  time.Time
}

func NewKeyboardShooter(container *Element, cooldown time.Duration) *KeyboardShooter {
	return &KeyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (ks *KeyboardShooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	container := ks.container

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(ks.lastShot) >= ks.cooldown {
			ks.Shoot(container.position.x+25, container.position.y-20)
			ks.Shoot(container.position.x-25, container.position.y-20)
			ks.lastShot = time.Now()
		}
	}

	if keys[sdl.SCANCODE_R] == 1 {
		ReloadBullets()
	}

	return nil
}

func (ks *KeyboardShooter) Shoot(x, y float64) {
	if bullet, exists := GetBullet(); exists {
		bullet.active = true
		bullet.position.x = x
		bullet.position.y = y
		bullet.rotation = 270 * (math.Pi / 180)
	}
}

func (ks *KeyboardShooter) OnDraw(renderer *sdl.Renderer) error {
	return nil
}
