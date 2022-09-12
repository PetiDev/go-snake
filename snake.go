package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Snake struct {
	snake     []Snakenode
	direction string
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

	for int32(s.snake[0].x)%int32(snakeSize) != 0 || int32(s.snake[0].y)%int32(snakeSize) != 0 {
	}
	s.direction = dir
	s.snake[0].heading = dir

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

	if rl.IsKeyPressed(rl.KeyW) {
		go s.turn("up")
	}
	if rl.IsKeyPressed(rl.KeyS) {
		go s.turn("down")
	}
	if rl.IsKeyPressed(rl.KeyA) {
		go s.turn("left")
	}
	if rl.IsKeyPressed(rl.KeyD) {
		go s.turn("right")
	}
	//TODO remove in production
	if rl.IsKeyPressed(rl.KeySpace) {
		s.grow(rl.DarkGreen, snakeSize*2)
	}

	s.move()
}
