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
	g.snake.grow(g.food.color, 1)
	score = score + 1
	g.food.init()
}

func (g *Gamescreen) draw() {
	drawCenteredText("Score: "+strconv.Itoa(score), screenWidt/2, 10, 20, rl.White)
	if rl.CheckCollisionRecs(rl.NewRectangle(g.snake.getHeading().X, g.snake.getHeading().Y, 1, 1), rl.NewRectangle(float32(g.food.x), float32(g.food.y), float32(snakeSize), float32(snakeSize))) {
		g.eatFood()
	}

	g.food.draw()
	g.snake.draw()

}

func (g *Gamescreen) init() {
	rand.Seed(time.Now().Unix())

	g.snake = &Snake{
		direction: "right",
	}
	g.food = new(Food)
	g.food.init()

	g.snake.snake = append(g.snake.snake, Snakenode{x: 100, y: 100, color: rl.ColorFromHSV(rand.Float32()*360, 1, 1)})
}
