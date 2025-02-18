package main

import (
	"image/color"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	canvasWidth, canvasHeight float32 = 320, 320
	windowAspectRatio                 = 2
)

type Ball struct {
	x, y           float32
	radius         float32
	dx, dy         float32 // The ball's direction. 1 for right/down, -1 for left/up.
	xspeed, yspeed float32
}

type Paddle struct {
	x, y          float32
	width, height float32
}

type Game struct {
	ball   Ball
	paddle Paddle
	keys   []ebiten.Key
}

// TODO: Move this to the Ball struct
func (g *Game) MoveBall() {
	g.ball.x += g.ball.dx * g.ball.xspeed
	g.ball.y += g.ball.dy * g.ball.yspeed

	ballRX := g.ball.x + (g.ball.radius * 2)

	if g.ball.x < 0 || ballRX > canvasWidth {
		g.ball.dx *= -1
		if ballRX > canvasWidth {
			g.ball.x = canvasWidth + (canvasWidth - ballRX)
		}
	}
	if g.ball.y < 0 || g.ball.y+(g.ball.radius*2) > canvasHeight {
		g.ball.dy *= -1
	}
}

func (g *Game) MovePaddle() {
	if slices.Contains(g.keys, ebiten.KeyArrowDown) {
		g.paddle.y += 5
	}
	if slices.Contains(g.keys, ebiten.KeyArrowUp) {
		g.paddle.y -= 5
	}
}

// This is more useful when the window is resizeable.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(canvasWidth), int(canvasHeight)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.MoveBall()
	g.MovePaddle()
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	ebitenutil.DebugPrint(s, "Pon-Go!")

	vector.DrawFilledCircle(s, g.ball.x+g.ball.radius, g.ball.y+g.ball.radius, g.ball.radius, color.White, true)
	vector.DrawFilledRect(s, g.paddle.x, g.paddle.y, g.paddle.width, g.paddle.height, color.White, true)
}

func main() {
	ebiten.SetWindowSize(int(canvasWidth*windowAspectRatio), int(canvasHeight*windowAspectRatio))
	ebiten.SetWindowTitle("Pon-Go!")

	g := Game{
		ball: Ball{
			x:      0,
			y:      0,
			radius: 5,
			dx:     1,
			dy:     1,
			xspeed: 6,
			yspeed: 3,
		},
		paddle: Paddle{
			x:      canvasWidth - 10,
			y:      (canvasHeight / 2) - 40,
			width:  10,
			height: 80,
		},
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
