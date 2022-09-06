package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidt   = int32(800)
	screenHeight = int32(450)
	gamescreen   = &Gamescreen{}
	menuscreen   = &Menu{}
	diescreen    = &DieScreen{}
	snakeSpeed   = 1
	snakeSize    = 10

	state     = 0
	score     = 0
	highScore = 0
	dieReason = ""
)

func drawCenteredText(text string, x int32, y int32, fontSize int32, color rl.Color) {
	rl.DrawText(text, x-rl.MeasureText(text, fontSize)/2, y-fontSize/2, fontSize, color)
}

func main() {
	rl.InitWindow(int32(screenWidt), int32(screenHeight), "Super Epic Snake (Ami még nem működik)")
	rl.SetTargetFPS(60)

	menuscreen.init()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch state {
		case 1:
			gamescreen.draw()
		case 2:
			diescreen.draw()
		default:
			menuscreen.draw()
		}

		rl.DrawFPS(2, 0)
		rl.EndDrawing()
	}
}
