package main

import (
	"image/color"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (p *Paddle) Draw(s *ebiten.Image) {
	vector.DrawFilledRect(s, p.x, p.y, p.width, p.height, color.White, true)
}
