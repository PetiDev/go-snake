package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gamescreen struct {
	snake *Snake
	food  *Food
}

func (g *Gamescreen) eatFood() {
	g.snake.grow(g.food.color, snakeSize)
	score = score + 1
	g.food.init()
}

func (g *Gamescreen) draw() {
	drawCenteredText("Score: "+strconv.Itoa(score), screenWidt/2, 10, 20, rl.White)
	if g.snake.snake[0].x == g.food.x && g.snake.snake[0].y == g.food.y {
		g.eatFood()
	}

	g.food.draw()
	g.snake.draw()

	//TODO remove in production
	if rl.IsKeyPressed(rl.KeyR) {
		g.food.init()
	}
}

func (g *Gamescreen) init() {
	rand.Seed(time.Now().Unix())

	g.snake = &Snake{
		direction: "right",
	}
	g.food = new(Food)
	g.food.init()

	g.snake.snake = append(g.snake.snake, Snakenode{x: 100, y: 100, color: rl.Red})
	g.snake.snake = append(g.snake.snake, Snakenode{x: 100, y: 100, color: rl.ColorFromHSV(rand.Float32()*360, 1, 1)})
}
