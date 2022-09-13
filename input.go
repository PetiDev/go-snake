package main

import (
	"fmt"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Input struct {
	x            int32
	y            int32
	width        int32
	height       int32
	color        rl.Color
	hoverColor   rl.Color
	defaultColor rl.Color
	activeColor  rl.Color
	isActive     bool
	biteSlice    []byte
	placeholder  string
	isPass       bool
}

func (i *Input) draw() {
	rl.DrawRectangleLinesEx(rl.NewRectangle(float32(i.x)-float32(i.width)/2, float32(i.y), float32(i.width), float32(i.height)), 5, i.color)

	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()

	if mouseX >= i.x-i.width/2 && mouseY >= i.y && mouseX <= i.x+i.width-i.width/2 && mouseY <= i.y+i.height {
		i.color = i.hoverColor
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			i.isActive = true
		}
	} else {
		i.color = i.defaultColor
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			i.isActive = false
		}
	}
	if i.isActive {
		i.color = i.activeColor
		i.textInput()
	}
	if len(i.biteSlice) == 0 {
		drawCenteredText(i.placeholder, i.x, i.y+30, 30, rl.Gray)
	} else {
		if i.isPass {
			drawCenteredText(strings.Repeat("*", len(i.biteSlice)), i.x, i.y+30, 30, rl.Gray)
		} else {
			drawCenteredText(string(i.biteSlice), i.x, i.y+30, 30, rl.Gray)
		}
	}
}

func (i *Input) textInput() {
	keyPressed := rl.GetKeyPressed()
	charPressed := rl.GetCharPressed()
	if keyPressed == 0 {
		return
	}
	if keyPressed == rl.KeyBackspace {
		if len(i.biteSlice) > 0 {
			i.biteSlice = i.biteSlice[:len(i.biteSlice)-1]
		}
		fmt.Println(string(i.biteSlice))
		return
	}

	if charPressed == 0 || len(i.biteSlice) >= 16 {
		return
	}

	i.biteSlice = append(i.biteSlice, byte(charPressed))
	fmt.Println(string(i.biteSlice))
}
