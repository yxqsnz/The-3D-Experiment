package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawDebugInfo(state *State) {
	if state.IsShowingDebugInfo {
		state.Text.DrawNextLine(fmt.Sprintf("Frame time: %.3f", rl.GetFrameTime()), 20, rl.Brown)
		state.Text.DrawNextLine(fmt.Sprintf("FPS: %d", rl.GetFPS()), 20, rl.Brown)
	}
}
