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
	speed         float32
}

func NewPaddle(canvasWidth float32, canvasHeight float32) Paddle {
	var w float32 = 10

	return Paddle{
		x:      canvasWidth - w,
		y:      (canvasHeight / 2) - 40,
		width:  w,
		height: 80,
		speed:  5,
	}
}

func (p *Paddle) Move(canvasHeight float32, keys []ebiten.Key) {
	if slices.Contains(keys, ebiten.KeyArrowDown) {
		if (p.y + p.height) < canvasHeight {
			p.y += p.speed
		}
	}
	if slices.Contains(keys, ebiten.KeyArrowUp) {
		if p.y > 0 {
			p.y -= p.speed
		}
	}
}

func (p *Paddle) Draw(s *ebiten.Image) {
	vector.DrawFilledRect(s, p.x, p.y, p.width, p.height, color.White, true)
}
