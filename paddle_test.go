package main

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestMovePaddleUp(t *testing.T) {
	p := NewPaddle(320, 320)

	if p.y != 120 {
		t.Error("Expected Paddle to be at Y position (120), got", p.y)
	}

	keys := []ebiten.Key{ebiten.KeyArrowUp}

	p.Move(320, keys)

	if p.y != 115 {
		t.Error("Expected Paddle to be at Y position (120), got", p.y)
	}
}

func TestMovePaddleDown(t *testing.T) {
	p := NewPaddle(320, 320)

	if p.y != 120 {
		t.Error("Expected Paddle to be at Y position (120), got", p.y)
	}

	keys := []ebiten.Key{ebiten.KeyArrowDown}

	p.Move(320, keys)

	if p.y != 125 {
		t.Error("Expected Paddle to be at Y position (120), got", p.y)
	}
}
