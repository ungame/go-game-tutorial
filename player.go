package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Player struct {
	tex  *sdl.Texture
	w, h int32
}

func NewPlayer(renderer *sdl.Renderer) (*Player, error) {
	var p Player
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		log.Panicln(err)
	}
	defer img.Free()
	p.w = img.W
	p.h = img.H

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		log.Panicln(err)
	}
	return &p, nil
}

func (p *Player) Draw(renderer *sdl.Renderer) error {
	src := &sdl.Rect{
		X: 0,
		Y: 0,
		W: p.w,
		H: p.h,
	}
	dst := &sdl.Rect{
		X: screenWidth/2 - p.w/2,
		Y: screenHeight/2 - p.h/2,
		W: p.w,
		H: p.h,
	}

	return renderer.Copy(p.tex, src, dst)
}

func (p *Player) Destroy() {
	err := p.tex.Destroy()
	if err != nil {
		log.Panicln(err)
	}
}
