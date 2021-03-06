package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const bulletSpeed = 15

func NewBullet(renderer *sdl.Renderer) *Element {
	var b Element
	b.tag = "bullet"
	b.position = Vector{}
	sr := NewSpriteRenderer(&b, renderer, "sprites/player_bullet.bmp")
	b.AddComponent(sr)

	mover := NewBulletMover(&b, bulletSpeed)
	b.AddComponent(mover)

	col := Circle{}
	col.center = b.position
	col.radius = 8

	b.collisions = append(b.collisions, col)

	return &b
}

var Bullets []*Element

func InitBullets(renderer *sdl.Renderer, max int) {
	for i := 0; i < max; i++ {
		bullet := NewBullet(renderer)
		Bullets = append(Bullets, bullet)
	}
}

func GetBullet() (*Element, bool) {
	for _, bullet := range Bullets {
		if !bullet.active {
			return bullet, true
		}
	}
	return nil, false
}

func ReloadBullets() {
	if _, exists := GetBullet(); !exists {
		for _, bullet := range Bullets {
			bullet.active = false
		}
	}
	fmt.Print("\rReload bullets")
}
