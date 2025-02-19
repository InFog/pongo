package main

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
