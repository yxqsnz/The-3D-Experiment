package game

import (
	"fmt"

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

	if state.IsShowingDebugInfo {
		state.Text.DrawNextLine(fmt.Sprintf("FPS: %d", rl.GetFPS()), 20, rl.Pink)
	}
}
