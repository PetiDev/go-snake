package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snakenode struct {
	x       float32
	y       float32
	heading string
	color   rl.Color
}
