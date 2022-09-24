package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	gameVersion = "v2.3.1"
)

var (
	currentVersion = ""
	screenWidt     = int32(500)
	screenHeight   = int32(500)
	gamescreen     = &Gamescreen{}
	menuscreen     = &Menu{}
	diescreen      = &DieScreen{}
	authscreen     = &AuthScreen{}
	snakeSpeed     = float32(70)
	snakeSize      = 10
	password       = ""
	username       = ""
	gameOpened     = int64(0)
	lastGametime   = int64(0)
	deaths         = 0
	state          = 0
	score          = 0
	highScore      = 0
	dieReason      = ""
)

func drawCenteredText(text string, x int32, y int32, fontSize int32, color rl.Color) {
	rl.DrawText(text, x-rl.MeasureText(text, fontSize)/2, y-fontSize/2, fontSize, color)
}

func getCurrentVersion() string {
	res, err := http.Get("https://go-snake-backend.fly.dev/get/version/")

	if err != nil {
		fmt.Println("Error getting current version")
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error getting current version")
	}

	return string(data)
}

func main() {

	gameOpened = time.Now().Unix()

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(screenWidt), int32(screenHeight), "Super Epic Snake (Ami még nem működik)")

	currentVersion = getCurrentVersion()
	menuscreen.init()

	fmt.Println(`----------------
	Never gonna give you up
	Never gonna let you down
	Never gonna run around and desert you
	Never gonna make you cry
	Never gonna say goodbye
	Never gonna tell a lie and hurt you`)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch state {
		case 1:
			gamescreen.draw()
		case 2:
			diescreen.draw()
		case 3:
			authscreen.draw()
		default:
			menuscreen.draw()
		}

		rl.DrawFPS(2, 0)
		rl.EndDrawing()
	}
}
