package main

import rl "github.com/gen2brain/raylib-go/raylib"

type DieScreen struct {
	menu *Button
}

func (d *DieScreen) draw() {
	drawCenteredText(dieReason, screenWidt/2, screenHeight/2-20, 32, rl.Red)
	d.menu.draw()
}
func (d *DieScreen) init() {
	d.menu = &Button{
		x:              screenWidt / 2,
		y:              screenHeight/2 + 30,
		fontSize:       40,
		grownFontSize:  50,
		normalFontSize: 40,
		color:          rl.Beige,
		text:           "MENU",
		callback: func() {
			menuscreen.init()
			state = 0
		},
	}
}
