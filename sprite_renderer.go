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

	x := sr.container.position.x - sr.width/2.0
	y := sr.container.position.y - sr.height/2.0

	src := &sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(sr.width),
		H: int32(sr.height),
	}
	dst := &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: int32(sr.width),
		H: int32(sr.height),
	}

	point := &sdl.Point{
		X: int32(sr.width) / 2,
		Y: int32(sr.height) / 2,
	}

	return renderer.CopyEx(
		sr.tex,
		src,
		dst,
		sr.container.rotation,
		point,
		sdl.FLIP_NONE,
	)

}

func (sr *SpriteRenderer) OnUpdate() error {
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
