package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type BulletMover struct {
	container *Element
	speed     float64
}

func NewBulletMover(container *Element, speed float64) *BulletMover {
	var bm BulletMover
	bm.container = container
	bm.speed = speed
	return &bm
}

func (bm *BulletMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (bm *BulletMover) OnUpdate() error {
	bm.container.position.x += bulletSpeed * math.Cos(bm.container.rotation)
	bm.container.position.y += bulletSpeed * math.Sin(bm.container.rotation)

	for index := range bm.container.collisions {
		bm.container.collisions[index].center = bm.container.position
	}

	//if bm.container.position.x > screenWidth ||
	//	bm.container.position.x < 0 ||
	//	bm.container.position.y > screenHeight ||
	//	bm.container.position.y < 0 {
	//
	//	bm.container.active = false
	//}

	return nil
}

func (bm *BulletMover) OnCollision(_ *Element) error {
	bm.container.active = false
	return nil
}
