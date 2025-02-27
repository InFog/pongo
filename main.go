package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	canvasWidth, canvasHeight float32 = 320, 320
	windowAspectRatio                 = 2
)

type Obstacle struct {
	x, y          float32
	width, height float32
	speed         float32
}

func (o *Obstacle) Move() {
	o.y += o.speed
}

func (o *Obstacle) Draw(s *ebiten.Image) {
	vector.DrawFilledRect(s, o.x, o.y, o.width, o.height, color.White, true)
}

func NewObstacle() Obstacle {
	var height float32 = float32(rand.Intn(40) + 10)
	var x float32 = float32(rand.Intn(100) + 100)
	var speed float32 = float32(rand.Intn(2) + 1)
	return Obstacle{
		x: x, y: -height,
		width: 10, height: height,
		speed: speed,
	}
}

type Game struct {
	ball      Ball
	echoes    []Echo
	obstacles []Obstacle
	paddle    Paddle
	keys      []ebiten.Key
	score     int
	highScore int
}

// This is more useful when the window is resizeable.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(canvasWidth), int(canvasHeight)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.ball.Move(canvasHeight)
	var newEchoes []Echo = []Echo{}
	var newObstacles []Obstacle = []Obstacle{}

	for i := len(g.echoes) - 1; i >= 0; i-- {
		g.echoes[i].Update()

		if g.echoes[i].radius < 40 {
			newEchoes = append(newEchoes, g.echoes[i])
		}
	}

	var lastObstacleY float32 = canvasHeight

	for i := len(g.obstacles) - 1; i >= 0; i-- {
		g.obstacles[i].Move()

		if g.obstacles[i].y < canvasHeight {
			newObstacles = append(newObstacles, g.obstacles[i])
		}

		if g.obstacles[i].y < lastObstacleY {
			lastObstacleY = g.obstacles[i].y
		}

		if g.ball.CheckObstacleCollision(g.obstacles[i]) {
			newEchoes = append(g.echoes, Echo{
				g.ball.x, g.ball.y, g.ball.radius,
			})
		}
	}

	if lastObstacleY > canvasHeight/2 {
		newObstacles = append(newObstacles, NewObstacle())
	}

	g.echoes = newEchoes
	g.obstacles = newObstacles

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

	for _, e := range g.echoes {
		e.Draw(s)
	}

	for _, o := range g.obstacles {
		o.Draw(s)
	}
}

func main() {
	ebiten.SetWindowSize(int(canvasWidth*windowAspectRatio), int(canvasHeight*windowAspectRatio))
	ebiten.SetWindowTitle("Pon-Go!")

	g := Game{
		ball:      NewBall(canvasHeight),
		paddle:    NewPaddle(canvasWidth, canvasHeight),
		obstacles: []Obstacle{NewObstacle()},
		score:     0,
		highScore: 0,
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
