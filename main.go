package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	canvasWidth, canvasHeight float32 = 320, 320
	windowAspectRatio                 = 2
)

type Game struct {
	ball      Ball
	echoes    []Echo
	paddle    Paddle
	keys      []ebiten.Key
	score     int
	highScore int
}

type Echo struct {
	x, y   float32
	radius float32
}

func (e *Echo) Draw(s *ebiten.Image) {
	vector.StrokeCircle(s, e.x+e.radius, e.y+e.radius, e.radius, 1, color.White, true)
}

func (e *Echo) Update() {
	e.radius++
	e.x--
	e.y--
}

// This is more useful when the window is resizeable.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(canvasWidth), int(canvasHeight)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.ball.Move(canvasHeight)
	var newEchoes []Echo = []Echo{}

	for i := len(g.echoes) - 1; i >= 0; i-- {
		g.echoes[i].Update()

		if g.echoes[i].radius < 40 {
			newEchoes = append(newEchoes, g.echoes[i])
		}
	}

	g.echoes = newEchoes

	if g.ball.CheckWallCollision() {
		g.echoes = append(g.echoes, Echo{
			g.ball.x, g.ball.y, g.ball.radius,
		})
	}

	g.paddle.Move(canvasHeight, g.keys)

	if g.ball.CheckPaddleCollision(g.paddle) {
		g.score++

		if g.score > g.highScore {
			g.highScore = g.score
		}

		if int(g.ball.xspeed)%3 == 0 {
			g.ball.xspeed++
		}

		g.echoes = append(g.echoes, Echo{
			g.ball.x, g.ball.y, g.ball.radius,
		})
	}

	if g.ball.CheckOutOfBounds(canvasWidth) {
		g.ball = NewBall(canvasHeight)
		g.score = 0
	}

	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	ebitenutil.DebugPrint(s, fmt.Sprintf("Pon-Go!\nScore: %d\nHigh Score: %d", g.score, g.highScore))

	g.paddle.Draw(s)
	g.ball.Draw(s)

	for _, echo := range g.echoes {
		echo.Draw(s)
	}
}

func main() {
	ebiten.SetWindowSize(int(canvasWidth*windowAspectRatio), int(canvasHeight*windowAspectRatio))
	ebiten.SetWindowTitle("Pon-Go!")

	g := Game{
		ball:      NewBall(canvasHeight),
		paddle:    NewPaddle(canvasWidth, canvasHeight),
		score:     0,
		highScore: 0,
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
