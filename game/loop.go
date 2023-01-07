package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PreDraw(state *State) {
	if rl.IsKeyPressed(rl.KeyF3) {
		state.IsShowingDebugInfo = !state.IsShowingDebugInfo
	}

	state.Text.Reset()
}

func Draw(state *State) {
	rl.ClearBackground(rl.RayWhite)
	maxX, maxY := rl.GetScreenHeight(), rl.GetScreenWidth()

	rl.DrawCircle(int32(maxY/2), int32(maxX/2), 100, rl.Red)

	DrawDebugInfo(state)
}
