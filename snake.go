package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Snake struct {
	snake     []Snakenode
	direction string
	path      []Snakenode
}

func (s *Snake) grow(color rl.Color, count int) {
	for i := 0; i < count; i++ {
		s.snake = append(s.snake, Snakenode{x: -100, y: -100, color: color})
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
		return rl.NewVector2(float32(s.snake[0].x), float32(s.snake[0].y)-float32(snakeSize))
	case "down":
		return rl.NewVector2(float32(s.snake[0].x), float32(s.snake[0].y)+float32(snakeSize))
	case "left":
		return rl.NewVector2(float32(s.snake[0].x)-float32(snakeSize), float32(s.snake[0].y))
	case "right":
		return rl.NewVector2(float32(s.snake[0].x)+float32(snakeSize), float32(s.snake[0].y))
	default:
		return rl.NewVector2(0, 0)
	}
}

func (s *Snake) move(direction string) {

	switch direction {
	case "up":

		if s.snake[0].y > 0 {
			s.snake[0].y = int32(s.snake[0].y) - int32(snakeSpeed)
		} else {
			s.die("Beütötted a fejed")
		}
	case "down":

		if s.snake[0].y < screenHeight-int32(snakeSize) {
			s.snake[0].y = int32(s.snake[0].y) + int32(snakeSpeed)
		} else {
			s.die("Falnak ütköztél")
		}

	case "left":

		if s.snake[0].x > 0 {
			s.snake[0].x = int32(s.snake[0].x) - int32(snakeSpeed)
		} else {
			s.die("Nem arra van a pálya")
		}
	case "right":

		if s.snake[0].x < screenWidt-int32(snakeSize) {
			s.snake[0].x = int32(s.snake[0].x) + int32(snakeSpeed)
		} else {
			s.die("Nincs ötletem, de meghaltál")
		}
	}
}

func (s *Snake) turn(dir string) {
	if s.direction == dir || (s.direction == "up" && dir == "down") || (s.direction == "down" && dir == "up") || (s.direction == "left" && dir == "right") || (s.direction == "right" && dir == "left") {
		return
	}
	for s.snake[0].x%int32(snakeSize) != 0 || s.snake[0].y%int32(snakeSize) != 0 {

	}

	s.direction = dir

}

func (s *Snake) draw() {

	for i := len(s.snake) - 1; i > 0; i-- {
		rl.DrawRectangle(s.snake[i].x, s.snake[i].y, int32(snakeSize), int32(snakeSize), s.snake[i].color)
		if i != 0 {

			s.snake[i].x = int32(s.snake[i-1].x)
			s.snake[i].y = int32(s.snake[i-1].y)

			if i > 2 && rl.CheckCollisionRecs(rl.NewRectangle(float32(s.getHeading().X), float32(s.getHeading().Y), float32(snakeSize), float32(snakeSize)), rl.NewRectangle(float32(s.snake[i].x), float32(s.snake[i].y), float32(snakeSize), float32(snakeSize))) {
				s.die("Magadnak mentél volna")
			}

		}
	}

	s.move(s.direction)

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
		s.grow(rl.DarkGreen, snakeSize)
	}
}
