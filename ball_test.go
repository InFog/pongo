package main

import (
	"testing"
)

func TestMoveBallForwards(t *testing.T) {
	b := NewBall(320)

	b.x = 0
	b.y = 0
	b.xspeed = 2
	b.yspeed = 3

	b.Move(320)

	if b.x != 2 || b.y != 3 {
		t.Error("Expected Ball to be in position (2 3) , got", b.x, b.y)
	}
}

func TestMoveBallBackwards(t *testing.T) {
	b := NewBall(320)

	b.x = 10
	b.y = 10

	b.xspeed = 2
	b.yspeed = 3

	b.dx = -1
	b.dy = -1

	b.Move(320)

	if b.x != 8 || b.y != 7 {
		t.Error("Expected Ball to be in position (8 7) , got", b.x, b.y)
	}
}
