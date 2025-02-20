package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y           float32
	radius         float32
	dx, dy         float32 // The ball's direction. 1 for right/down, -1 for left/up.
	xspeed, yspeed float32
}

func NewBall(canvasHeight float32) Ball {
	return Ball{
		x:      0,
		y:      float32(rand.Intn(int(canvasHeight))),
		radius: 5,
		dx:     1,
		dy:     1,
		xspeed: 6,
		yspeed: 3,
	}
}

func (b *Ball) Move(canvasHeight float32) {
	b.x += b.dx * b.xspeed
	b.y += b.dy * b.yspeed

	ballRY := b.y + (b.radius * 2)

	if b.x <= 0 {
		b.dx *= -1
	}
	if b.y <= 0 || ballRY > canvasHeight {
		b.dy *= -1
	}
}

func (b *Ball) Draw(s *ebiten.Image) {
	vector.DrawFilledCircle(s, b.x+b.radius, b.y+b.radius, b.radius, color.White, true)
}

func (b *Ball) CheckOutOfBounds(sw float32) bool {
	if b.x+(b.radius*2) > sw {
		return true
	}

	return false
}

func (b *Ball) CheckPaddleCollision(p Paddle) bool {
	bw := b.x + (b.radius * 2)

	if bw >= p.x {
		bh := b.y + (b.radius * 2)

		// Hitting the paddle in the top
		if bh > p.y && b.y < p.y {
			b.dx *= -1
			b.dy *= -1
			return true
		}

		ph := p.y + p.height

		// Hitting the paddle in the bottom
		if b.y < (ph) && bh > ph {
			b.dx *= -1
			b.dy *= -1
			return true
		}

		if bh >= p.y && bh <= p.y+p.height {
			b.dx *= -1
			return true
		}
	}

	return false
}
