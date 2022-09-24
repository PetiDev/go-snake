package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

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
	if gameVersion != currentVersion {
		drawCenteredText("Outdated version", screenWidt/2, screenHeight-20, 30, rl.Blue)
	}
	rl.DrawText(currentVersion, 4, screenHeight-24, 20, rl.White)

	for k, v := range m.scoreboard {
		if k > 10 {
			return
		}
		drawCenteredText(fmt.Sprintf("%s <---------> %d", v.Name, v.Points), screenWidt/2, int32(180+k*30), 26, rl.Red)
	}

	rl.DrawText("Session:", 20, screenHeight-120, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Deaths: %d", deaths), 30, screenHeight-100, 20, rl.White)
	rl.DrawText(fmt.Sprintf("lastGametime: %dmin %dsec", time.Unix(lastGametime, 0).Minute(), time.Unix(lastGametime, 0).Second()), 30, screenHeight-80, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Opened: %s", time.Unix(gameOpened, 0).Format("2006-01-02, 15:04:05")), 30, screenHeight-60, 20, rl.White)

}

func (m *Menu) getScoreboardData() {
	res, err := http.Get("https://go-snake-backend.fly.dev/get")

	if err != nil {
		fmt.Printf("Error fetching scoreboard: %s\n", err.Error())
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error fetching scoreboard: %s\n", err.Error())
		return
	}
	scoreboard := make([]ScoreboardEntry, 0)
	err = json.Unmarshal(data, &scoreboard)
	if err != nil {
		fmt.Printf("Error fetching scoreboard: %s\n", err.Error())
		return
	}
	m.scoreboard = scoreboard
}

func (m *Menu) init() {
	go m.getScoreboardData()
	m.play = &Button{
		x:              screenWidt / 2,
		y:              100,
		text:           "PLAY",
		color:          rl.White,
		defaultColor:   rl.White,
		activeColor:    rl.Beige,
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
		defaultColor:   rl.White,
		activeColor:    rl.Beige,
		fontSize:       30,
		grownFontSize:  40,
		normalFontSize: 30,
		callback: func() {
			authscreen.init()
			state = 3
		},
	}
}
