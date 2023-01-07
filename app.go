package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"me.yxqsnz.The-3D-Experiment/game"
)

func main() {
	rl.InitWindow(800, 450, "The 3D Experiment")
	rl.SetTargetFPS(60)
	state := game.InitializeState()

	for !rl.WindowShouldClose() {
		game.PreDraw(&state)
		rl.BeginDrawing()
		game.Draw(&state)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
