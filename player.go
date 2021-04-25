package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math"
	"time"
)

const (
	playerSize         = 105
	playerSpeed        = 0.07
	playerShotCooldown = time.Millisecond * 250
)

type Player struct {
	tex  *sdl.Texture
	x, y float64

	lastShot time.Time
}

func NewPlayer(renderer *sdl.Renderer) *Player {
	var p Player
	p.tex = GetTextureFromBMP(renderer, "sprites/player.bmp")
	p.x = screenWidth / 2
	p.y = screenHeight
	return &p
}

func (p *Player) Draw(renderer *sdl.Renderer) error {
	x := p.x - float64(playerSize/2)
	y := p.y - float64(playerSize)

	src := &sdl.Rect{
		X: 0,
		Y: 0,
		W: playerSize,
		H: playerSize,
	}
	dst := &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: playerSize,
		H: playerSize,
	}

	return renderer.Copy(p.tex, src, dst)
}

func (p *Player) Update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x - (playerSize / 2) > 0 {
			p.x -= playerSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x + (playerSize / 2) < screenWidth {
			p.x += playerSpeed
		}
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown {
			p.Shoot(p.x+25, p.y-20)
			p.Shoot(p.x-25, p.y-20)
			p.lastShot = time.Now()
		}
	}

	if keys[sdl.SCANCODE_R] == 1 {
		ReloadBullets()
	}
}

func (p *Player) Shoot(x, y float64) {
	if bullet, exists := GetBullet(); exists {
		bullet.active = true
		bullet.x = x
		bullet.y = y
		bullet.angle = 270 * (math.Pi / 180)
	}
}

func (p *Player) Destroy() {
	err := p.tex.Destroy()
	if err != nil {
		log.Panicln(err)
	}
}
