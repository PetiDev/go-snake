package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	x              int32
	y              int32
	fontSize       int32
	color          rl.Color
	text           string
	grownFontSize  int32
	normalFontSize int32
	callback       func()
}

func (b *Button) draw() {
	drawCenteredText(b.text, b.x, b.y, b.fontSize, b.color)

	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()

	if mouseY >= b.y-b.fontSize/2 && mouseY <= b.y+b.fontSize/2 && mouseX >= b.x-rl.MeasureText(b.text, b.fontSize)/2 && mouseX <= b.x+rl.MeasureText(b.text, b.fontSize)/2 {
		b.fontSize = b.grownFontSize
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			b.callback()
		}
	} else {
		b.fontSize = b.normalFontSize
	}
}
