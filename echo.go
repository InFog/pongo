package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Echo struct {
	x, y   float32
	radius float32
}

func (e *Echo) Draw(s *ebiten.Image) {
	c := color.White

	if e.radius > 20 {
		c = color.Gray16{0xaaaf}
	}
	if e.radius > 30 {
		c = color.Gray16{0x777f}
	}

	vector.StrokeCircle(s, e.x+e.radius, e.y+e.radius, e.radius, 1, c, true)
}

func (e *Echo) Update() {
	e.radius++
	e.x--
	e.y--
}
