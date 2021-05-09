package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"path/filepath"
)

func DrawTexture(tex *sdl.Texture, position Vector, rotation float64, renderer *sdl.Renderer) error {
	_, _, width, height, err := tex.Query()
	if err != nil {
		log.Panicln(err.Error())
	}

	x := position.x - float64(width)/2.0
	y := position.y - float64(height)/2.0

	src := &sdl.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	}
	dst := &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: width,
		H: height,
	}

	point := &sdl.Point{
		X: width / 2,
		Y: height / 2,
	}

	return renderer.CopyEx(
		tex,
		src,
		dst,
		rotation,
		point,
		sdl.FLIP_NONE,
	)
}

func LoadTextures(renderer *sdl.Renderer, folderPath string) []*sdl.Texture {
	var textures []*sdl.Texture

	filenames := ListFiles(folderPath)

	for _, filename := range filenames {
		tex := GetTextureFromBMP(renderer, filename)
		textures = append(textures, tex)
	}

	return textures
}

func ListFiles(folderPath string) []string {
	var files []string
	err := filepath.Walk(folderPath, func(path string, _ os.FileInfo, err error) error {
		files = append(files, path)
		return err
	})
	if err != nil {
		panic(err)
	}
	return files
}
