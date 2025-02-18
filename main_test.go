package main

import (
	"testing"
)

func createGame() Game {
	return Game{
		ball: Ball{
			x:      0,
			y:      0,
			radius: 5,
			dx:     1,
			dy:     1,
			xspeed: 0,
			yspeed: 0,
		},
	}
}

func TestMoveBallForwards(t *testing.T) {
	g := createGame()

	g.ball.xspeed = 2
	g.ball.yspeed = 3

	g.MoveBall()

	if g.ball.x != 2 || g.ball.y != 3 {
		t.Error("Expected Ball to be in position (2 3) , got", g.ball.x, g.ball.y)
	}
}

func TestMoveBallBackwards(t *testing.T) {
	g := createGame()

	g.ball.x = 10
	g.ball.y = 10

	g.ball.xspeed = 2
	g.ball.yspeed = 3

	g.ball.dx = -1
	g.ball.dy = -1

	g.MoveBall()

	if g.ball.x != 8 || g.ball.y != 7 {
		t.Error("Expected Ball to be in position (8 7) , got", g.ball.x, g.ball.y)
	}
}

func TestMoveBallHitWall(t *testing.T) {
	g := createGame()

	g.ball.xspeed = 10
	g.ball.yspeed = 10

	g.ball.x = canvasWidth - g.ball.radius - 5 // 310
	g.ball.y = 0

	g.MoveBall()

	if g.ball.x != 300 || g.ball.y != 10 {
		t.Error("Expected Ball to be in position (300 10) , got", g.ball.x, g.ball.y)
	}
}
