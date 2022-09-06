package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct {
	play *Button
}

func (m *Menu) draw() {
	m.play.draw()

	drawCenteredText("Score: "+strconv.Itoa(score), int32(screenWidt)/2, 10, 20, rl.White)
	drawCenteredText("Highscore: "+strconv.Itoa(highScore), screenWidt/2, 40, 30, rl.White)

}

func (m *Menu) init() {
	m.play = &Button{
		x:              screenWidt / 2,
		y:              100,
		text:           "PLAY",
		color:          rl.White,
		fontSize:       30,
		grownFontSize:  40,
		normalFontSize: 30,
		callback: func() {
			gamescreen.init()
			state = 1
			score = 0
		},
	}
}
