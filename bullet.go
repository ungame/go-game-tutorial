package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math"
)

const (
	bulletSize  = 32
	bulletSpeed = 0.2
)

type Bullet struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	active bool
}

func NewBullet(renderer *sdl.Renderer) *Bullet {
	var b Bullet
	b.tex = GetTextureFromBMP(renderer, "sprites/player_bullet.bmp")
	return &b
}

func (b *Bullet) Draw(renderer *sdl.Renderer) error {
	if !b.active {
		return nil
	}

	x := b.x - bulletSize/2.0
	y := b.y - bulletSize/2.0

	src := &sdl.Rect{
		X: 0,
		Y: 0,
		W: bulletSize,
		H: bulletSize,
	}
	dst := &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: bulletSize,
		H: bulletSize,
	}

	return renderer.Copy(b.tex, src, dst)
}

func (b *Bullet) Update() {
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)

	//if b.x > screenWidth || b.x < 0 || b.y > screenHeight || b.y < 0{
	//	b.active = false
	//}
}

func (b *Bullet) Destroy() {
	err := b.tex.Destroy()
	if err != nil {
		log.Panicln(err)
	}
}

var Bullets []*Bullet

func initBullets(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bullet := NewBullet(renderer)
		Bullets = append(Bullets, bullet)
		destructor = append(destructor, bullet)
	}
}

func GetBullet() (*Bullet, bool) {
	for index, bullet := range Bullets {

		fmt.Printf("\rBullets: %v/%v", index+1, len(Bullets))

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
}
