package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
	"path"
	"time"
)

type Animator struct {
	container       *Element
	sequences       map[string]*Sequence
	current         string
	lastFrameChange time.Time
	finished        bool
}

func NewAnimator(container *Element, sequences map[string]*Sequence, defaultSequence string) *Animator {
	var a Animator
	a.container = container
	a.sequences = sequences
	a.current = defaultSequence
	a.lastFrameChange = time.Now()
	return &a
}

func (a *Animator) OnDraw(renderer *sdl.Renderer) error {
	tex := a.sequences[a.current].GetTexture()

	return DrawTexture(tex, a.container.position, a.container.rotation, renderer)
}

func (a *Animator) OnUpdate() error {
	sequence := a.sequences[a.current]

	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(a.lastFrameChange) >= time.Duration(frameInterval) {
		a.finished = !sequence.Next()
		a.lastFrameChange = time.Now()
	}

	return nil
}

func (a *Animator) OnCollision(_ *Element) error {
	return nil
}

func (a *Animator) SetSequence(name string) {
	a.current = name
	a.lastFrameChange = time.Now()
}

func (a *Animator) HasFinished() bool {
	return a.finished
}

type Sequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64
	loop       bool
}

func NewSequence(dirPath string, sampleRate float64, loop bool, renderer *sdl.Renderer) *Sequence {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Panicln(err)
	}
	var s Sequence

	for _, file := range files {
		filepath := path.Join(dirPath, file.Name())
		tex := GetTextureFromBMP(renderer, filepath)
		s.textures = append(s.textures, tex)
	}

	s.sampleRate = sampleRate
	s.loop = loop

	return &s
}

func (s *Sequence) GetTexture() *sdl.Texture {
	return s.textures[s.frame]
}

func (s *Sequence) Next() bool {
	if s.IsLastFrame() {
		if s.loop {
			s.frame = 0
		} else {
			return false
		}
	} else {
		s.frame++
	}
	return true
}

func (s *Sequence) IsLastFrame() bool {
	return s.frame == len(s.textures)-1
}
