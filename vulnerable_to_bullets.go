package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type VulnerableToBullets struct {
	container *Element
}

func NewVulnerableToBullets(container *Element) *VulnerableToBullets {
	return &VulnerableToBullets{container}
}


func (bm *VulnerableToBullets) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (bm *VulnerableToBullets) OnUpdate() error {

	return nil
}

func (bm *VulnerableToBullets) OnCollision(other *Element) error {
	if other.tag == "bullet" {
		bm.container.active = false
	}
	return nil
}