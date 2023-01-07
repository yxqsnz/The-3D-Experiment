package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"me.yxqsnz.The-3D-Experiment/game"
)

func main() {
	rl.InitWindow(800, 450, "The 3D Experiment")

	refrate := rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())
	rl.SetTargetFPS(int32(refrate))

	state := game.InitializeState()
	game.Setup(&state)

	for !rl.WindowShouldClose() {
		game.PreDraw(&state)
		rl.BeginDrawing()
		game.Draw(&state)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
