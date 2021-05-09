package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type VulnerableToBullets struct {
	container *Element
	animator  *Animator
}

func NewVulnerableToBullets(container *Element) *VulnerableToBullets {
	var vtb VulnerableToBullets
	vtb.container = container
	vtb.animator, _ = container.GetComponent(&Animator{}).(*Animator)
	return &vtb
}

func (vtb *VulnerableToBullets) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (vtb *VulnerableToBullets) OnUpdate() error {
	switch vtb.animator.current {
	case "destroy":
		if vtb.animator.HasFinished() {
			vtb.container.active = false
		}
	}
	return nil
}

func (vtb *VulnerableToBullets) OnCollision(other *Element) error {
	if other.tag == "bullet" {
		vtb.animator.SetSequence("destroy")
	}
	return nil
}
