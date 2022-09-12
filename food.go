package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	x     int32
	y     int32
	color rl.Color
}

func (f *Food) draw() {
	rl.DrawRectangle(f.x, f.y, int32(snakeSize), int32(snakeSize), f.color)

}

func (f *Food) init() {
	f.x = (rand.Int31n((screenWidt)/int32(snakeSize)-4) + 2) * int32(snakeSize)
	f.y = (rand.Int31n((screenHeight)/int32(snakeSize)-4) + 2) * int32(snakeSize)

	f.color = rl.ColorFromHSV(rand.Float32()*360, 1, 1)
}
