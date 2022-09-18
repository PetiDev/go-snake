package main

import (
	"encoding/json"
	"net/http"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AuthScreen struct {
	submit   *Button
	register *Button
	menu     *Button
	name     *Input
	pass     *Input
}

func (a *AuthScreen) draw() {
	a.submit.draw()
	a.register.draw()
	a.menu.draw()
	a.name.draw()
	a.pass.draw()
}

func (a *AuthScreen) init() {
	a.menu = &Button{
		x:              screenWidt / 2,
		y:              screenHeight/2 + 40,
		fontSize:       30,
		color:          rl.White,
		defaultColor:   rl.White,
		activeColor:    rl.Beige,
		text:           "MENU",
		grownFontSize:  45,
		normalFontSize: 30,
		callback: func() {
			menuscreen.init()
			state = 0
		},
	}

	a.name = &Input{
		x:            screenWidt / 2,
		y:            20,
		width:        screenWidt/2 + 80,
		height:       60,
		color:        rl.White,
		defaultColor: rl.White,
		hoverColor:   rl.Beige,
		activeColor:  rl.Green,
		isActive:     false,
		placeholder:  "name",
		isPass:       false,
	}
	a.pass = &Input{
		x:            screenWidt / 2,
		y:            100,
		width:        screenWidt/2 + 80,
		height:       60,
		color:        rl.White,
		defaultColor: rl.White,
		hoverColor:   rl.Beige,
		activeColor:  rl.Green,
		isActive:     false,
		placeholder:  "password",
		isPass:       true,
	}
	a.submit = &Button{
		x:              screenWidt / 2,
		y:              screenHeight/2 - 40,
		fontSize:       30,
		color:          rl.White,
		defaultColor:   rl.White,
		activeColor:    rl.Beige,
		text:           "LOGIN",
		grownFontSize:  45,
		normalFontSize: 30,
		callback: func() {
			if score == 0 || string(a.name.biteSlice) == "" || string(a.pass.biteSlice) == "" {
				return
			}
			username = string(a.name.biteSlice)
			password = string(a.pass.biteSlice)
			a := map[string]interface{}{
				"name":   string(a.name.biteSlice),
				"pass":   string(a.pass.biteSlice),
				"points": score,
			}
			data, err := json.Marshal(a)
			if err != nil {
				panic("Something went wrong")
			}
			http.Post("https://go-snake-backend.fly.dev/write?version="+gameVersion, "application/json", strings.NewReader(string(data)))
		},
	}
	a.register = &Button{
		x:              screenWidt / 2,
		y:              screenHeight / 2,
		fontSize:       30,
		color:          rl.White,
		defaultColor:   rl.White,
		activeColor:    rl.Beige,
		text:           "REGISTER",
		grownFontSize:  45,
		normalFontSize: 30,
		callback: func() {
			if string(a.name.biteSlice) == "" || string(a.pass.biteSlice) == "" {
				return
			}
			username = string(a.name.biteSlice)
			password = string(a.pass.biteSlice)
			a := map[string]interface{}{
				"name":   string(a.name.biteSlice),
				"pass":   string(a.pass.biteSlice),
				"points": score,
			}
			data, err := json.Marshal(a)
			if err != nil {
				panic("Something went wrong")
			}
			http.Post("https://go-snake-backend.fly.dev/register?version="+gameVersion, "application/json", strings.NewReader(string(data)))
		},
	}
}
