package main

type Ball struct {
	x, y           float32
	radius         float32
	dx, dy         float32 // The ball's direction. 1 for right/down, -1 for left/up.
	xspeed, yspeed float32
}

func (b *Ball) Move() {
	b.x += b.dx * b.xspeed
	b.y += b.dy * b.yspeed

	ballRX := b.x + (b.radius * 2)

	if b.x < 0 || ballRX > canvasWidth {
		b.dx *= -1
		if ballRX > canvasWidth {
			b.x = canvasWidth + (canvasWidth - ballRX)
		}
	}
	if b.y < 0 || b.y+(b.radius*2) > canvasHeight {
		b.dy *= -1
	}
}
