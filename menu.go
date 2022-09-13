package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct {
	play       *Button
	auth       *Button
	scoreboard []ScoreboardEntry
}

type ScoreboardEntry struct {
	Id     string `json:"id"`
	Pass   string `json:"pass"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func (m *Menu) draw() {
	m.play.draw()
	m.auth.draw()
	drawCenteredText("Score: "+strconv.Itoa(score), int32(screenWidt)/2, 10, 20, rl.White)
	drawCenteredText("Highscore: "+strconv.Itoa(highScore), screenWidt/2, 40, 30, rl.White)

	for k, v := range m.scoreboard {
		if k > 10 {
			return
		}
		drawCenteredText(fmt.Sprintf("%s <---------> %d", v.Name, v.Points), screenWidt/2, int32(180+k*30), 26, rl.Red)
	}

}

func (m *Menu) getScoreboardData() {
	res, err := http.Get("https://go-snake-backend.fly.dev/get")

	if err != nil {
		panic("JAJ ne nincs scoreboard")
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic("JAJ ne nincs scoreboard")
	}
	scoreboard := make([]ScoreboardEntry, 0)
	err = json.Unmarshal(data, &scoreboard)
	if err != nil {
		panic("JAJ ne nincs scoreboard")
	}
	m.scoreboard = scoreboard
}

func (m *Menu) init() {
	m.getScoreboardData()
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
	m.auth = &Button{
		x:              screenWidt / 2,
		y:              150,
		text:           "AUTH",
		color:          rl.White,
		fontSize:       30,
		grownFontSize:  40,
		normalFontSize: 30,
		callback: func() {
			authscreen.init()
			state = 3
		},
	}
}
