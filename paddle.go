package main

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	x, y          float32
	width, height float32
}

func (p *Paddle) Move(keys []ebiten.Key) {
	if slices.Contains(keys, ebiten.KeyArrowDown) {
		p.y += 5
	}
	if slices.Contains(keys, ebiten.KeyArrowUp) {
		p.y -= 5
	}
}
