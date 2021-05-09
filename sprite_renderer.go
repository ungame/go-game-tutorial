package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type SpriteRenderer struct {
	container     *Element
	tex           *sdl.Texture
	width, height float64
}

func NewSpriteRenderer(element *Element, renderer *sdl.Renderer, filename string) *SpriteRenderer {
	tex := GetTextureFromBMP(renderer, filename)
	_, _, width, height, err := tex.Query()
	if err != nil {
		log.Panicln(err.Error())
	}
	var sr SpriteRenderer
	sr.container = element
	sr.tex = tex
	sr.width = float64(width)
	sr.height = float64(height)
	return &sr
}

func (sr *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	return DrawTexture(sr.tex, sr.container.position, sr.container.rotation, renderer)
}

func (sr *SpriteRenderer) OnUpdate() error {
	return nil
}

func (sr *SpriteRenderer) OnCollision(_ *Element) error {
	return nil
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
