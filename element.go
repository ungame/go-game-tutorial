package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"reflect"
)

type Vector struct {
	x, y float64
}

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
}

type Element struct {
	position   Vector
	rotation   float64
	active     bool
	tag        string
	collisions []Circle
	components []Component
}

func (e *Element) Draw(renderer *sdl.Renderer) {
	for _, component := range e.components {
		err := component.OnDraw(renderer)
		if err != nil {
			log.Panicln(err)
		}
	}
}

func (e *Element) Update() {
	for _, component := range e.components {
		err := component.OnUpdate()
		if err != nil {
			log.Panicln(err)
		}
	}
}

func (e *Element) Collision(other *Element) error {
	for _, component := range e.components {
		err := component.OnCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) AddComponent(c Component) {
	if e.hasComponent(c) {
		log.Panicln("component has already been added: ", reflect.TypeOf(c))
	}

	e.components = append(e.components, c)
}

func (e *Element) hasComponent(c Component) bool {
	return e.GetComponent(c) != nil
}

func (e *Element) GetComponent(cType Component) Component {
	for _, component := range e.components {
		if reflect.TypeOf(cType) == reflect.TypeOf(component) {
			return component
		}
	}
	return nil
}

func (e *Element) GetWidth() float64 {
	sr := e.GetComponent(&SpriteRenderer{}).(*SpriteRenderer)
	if sr != nil {
		return sr.width
	}
	return 0
}

func (e *Element) GetHeight() float64 {
	sr := e.GetComponent(&SpriteRenderer{}).(*SpriteRenderer)
	if sr != nil {
		return sr.height
	}
	return 0
}

func (e *Element) Destroy() {
	sr := e.GetComponent(&SpriteRenderer{}).(*SpriteRenderer)
	if sr != nil {
		if err := sr.tex.Destroy(); err != nil {
			log.Panicln(err.Error())
		}
	}
}
