package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	canvasWidth, canvasHeight float32 = 320, 320
	windowAspectRatio                 = 2
)

type Game struct {
	ball   Ball
	paddle Paddle
	keys   []ebiten.Key
}

// This is more useful when the window is resizeable.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(canvasWidth), int(canvasHeight)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.ball.Move()
	g.paddle.Move(g.keys)

	g.ball.CheckPaddleCollision(g.paddle)

	if g.ball.CheckOutOfBounds(canvasWidth) {
		g.ball = NewBall()
	}

	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	ebitenutil.DebugPrint(s, "Pon-Go!")

	g.paddle.Draw(s)
	g.ball.Draw(s)
}

func main() {
	ebiten.SetWindowSize(int(canvasWidth*windowAspectRatio), int(canvasHeight*windowAspectRatio))
	ebiten.SetWindowTitle("Pon-Go!")

	g := Game{
		ball: NewBall(),
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
