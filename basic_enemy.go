package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemyIdleDir    = "sprites/animations/basic_enemy/idle"
	basicEnemyDestroyDir = "sprites/animations/basic_enemy/destroy"
)

func NewBasicEnemy(renderer *sdl.Renderer, position Vector) *Element {
	var be Element
	be.position = position
	be.rotation = 180

	sequences := make(map[string]*Sequence, 2)

	sequences["idle"] = NewSequence(basicEnemyIdleDir, 5, true, renderer)
	sequences["destroy"] = NewSequence(basicEnemyDestroyDir, 15, false, renderer)

	animator := NewAnimator(&be, sequences, "idle")
	be.AddComponent(animator)

	vtb := NewVulnerableToBullets(&be)
	be.AddComponent(vtb)

	be.active = true

	col := Circle{}
	col.center = be.position
	col.radius = 38

	be.collisions = append(be.collisions, col)

	return &be
}
