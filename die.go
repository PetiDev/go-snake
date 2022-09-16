package main

import (
	"encoding/json"
	"net/http"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DieScreen struct {
	menu *Button
}

func (d *DieScreen) draw() {
	drawCenteredText(dieReason, screenWidt/2, screenHeight/2-20, 32, rl.Red)
	d.menu.draw()
}
func (d *DieScreen) init() {
	if score != 0 {
		if username != "" || password != "" {
			a := map[string]interface{}{
				"name":   username,
				"pass":   password,
				"points": score,
			}
			data, err := json.Marshal(a)
			if err != nil {
				panic("Something went wrong")
			}
			go http.Post("https://go-snake-backend.fly.dev/write", "application/json", strings.NewReader(string(data)))
		}
	}

	d.menu = &Button{
		x:              screenWidt / 2,
		y:              screenHeight/2 + 30,
		fontSize:       40,
		grownFontSize:  50,
		normalFontSize: 40,
		color:          rl.White,
		defaultColor:   rl.White,
		activeColor:    rl.Beige,
		text:           "MENU",
		callback: func() {
			menuscreen.init()
			state = 0
		},
	}
}
