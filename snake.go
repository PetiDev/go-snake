package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Snake struct {
	snake     []Snakenode
	direction string
	turnDir   string
}

func (s *Snake) grow(color rl.Color, count int) {
	for i := 0; i < count; i++ {
		s.snake = append(s.snake, Snakenode{x: s.snake[len(s.snake)-1].x, y: s.snake[len(s.snake)-1].y, color: color, heading: s.snake[len(s.snake)-1].heading})
	}
}

func (s *Snake) die(reason string) {
	diescreen.init()
	state = 2
	dieReason = reason
	if score > highScore {
		highScore = score
	}
}

func (s *Snake) getHeading() rl.Vector2 {
	switch s.direction {
	case "up":
		return rl.NewVector2(float32(s.snake[0].x)+float32(snakeSize)/2, float32(s.snake[0].y)-1)
	case "down":
		return rl.NewVector2(float32(s.snake[0].x)+float32(snakeSize)/2, float32(s.snake[0].y)+1)
	case "left":
		return rl.NewVector2(float32(s.snake[0].x)-1, float32(s.snake[0].y)+float32(snakeSize)/2)
	case "right":
		return rl.NewVector2(float32(s.snake[0].x)+1, float32(s.snake[0].y)+float32(snakeSize)/2)
	default:
		return rl.NewVector2(0, 0)
	}
}
func (s *Snake) turn(dir string) {
	if s.direction == dir || (s.direction == "up" && dir == "down") || (s.direction == "down" && dir == "up") || (s.direction == "left" && dir == "right") || (s.direction == "right" && dir == "left") {
		return
	}

	if int32(s.snake[0].x)%int32(snakeSize) != 0 || int32(s.snake[0].y)%int32(snakeSize) != 0 {
		return
	}
	s.direction = dir
	s.snake[0].heading = dir
	s.turnDir = ""

}
func (s *Snake) move() {
	for i := len(s.snake); i > 0; i-- {
		nthSnakeNode := i - 1
		rl.DrawRectangle(int32(s.snake[nthSnakeNode].x), int32(s.snake[nthSnakeNode].y), int32(snakeSize), int32(snakeSize), s.snake[nthSnakeNode].color)

		if nthSnakeNode != 0 {
			s.snake[nthSnakeNode].x = s.snake[nthSnakeNode-1].x
			s.snake[nthSnakeNode].y = s.snake[nthSnakeNode-1].y
			if nthSnakeNode > snakeSize*6 && rl.CheckCollisionRecs(rl.NewRectangle(s.getHeading().X, s.getHeading().Y, 1, 1), rl.NewRectangle(s.snake[nthSnakeNode].x, s.snake[nthSnakeNode].y, float32(snakeSize), float32(snakeSize))) {
				s.die("Magadnak mentél")
			}
		}
	}

	switch s.direction {
	case "up":
		if s.snake[0].y-snakeSpeed*rl.GetFrameTime() > 0 {
			s.snake[0].y = s.snake[0].y - snakeSpeed*rl.GetFrameTime()
		} else {
			s.die("Beütötted a fejed")
		}
	case "down":
		if s.snake[0].y+snakeSpeed*rl.GetFrameTime() < float32(screenHeight) {
			s.snake[0].y = s.snake[0].y + snakeSpeed*rl.GetFrameTime()
		} else {
			s.die("YOU DIED")
		}
	case "left":
		if s.snake[0].x-snakeSpeed*rl.GetFrameTime() > 0 {
			s.snake[0].x = s.snake[0].x - snakeSpeed*rl.GetFrameTime()
		} else {
			s.die("Nem arra van a pálya")
		}
	case "right":
		if s.snake[0].x-snakeSpeed*rl.GetFrameTime() < float32(screenWidt) {
			s.snake[0].x = s.snake[0].x + snakeSpeed*rl.GetFrameTime()
		} else {
			s.die("Nem vagyok túl kreatít, de meghaltál")
		}
	}
}

func (s *Snake) draw() {

	if rl.IsKeyDown(rl.KeyW) {
		s.turnDir = "up"
	}
	if rl.IsKeyDown(rl.KeyS) {
		s.turnDir = "down"
	}
	if rl.IsKeyDown(rl.KeyA) {
		s.turnDir = "left"
	}
	if rl.IsKeyDown(rl.KeyD) {
		s.turnDir = "right"
	}

	if s.turnDir != "" {
		s.turn(s.turnDir)
	}

	//TODO remove in production
	if rl.IsKeyPressed(rl.KeySpace) {
		s.grow(rl.DarkGreen, snakeSize*int(rl.GetFPS()/60))
	}
	if rl.IsKeyDown(rl.KeyE) {
		snakeSpeed = snakeSpeed + 1
	}
	if rl.IsKeyDown(rl.KeyQ) {
		snakeSpeed = snakeSpeed - 1
	}

	s.move()
}
