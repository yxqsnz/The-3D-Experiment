package game

import (
	"fmt"
	"runtime"

	"github.com/dustin/go-humanize"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawDebugInfo(state *State) {
	if state.IsShowingDebugInfo {
		mem := runtime.MemStats{}
		runtime.ReadMemStats(&mem)

		state.Text.DrawNextLine(fmt.Sprintf("Frame time: %.3f", rl.GetFrameTime()), 20, rl.Brown)
		state.Text.DrawNextLine(fmt.Sprintf("FPS: %d", rl.GetFPS()), 20, rl.Brown)
		state.Text.DrawNextLine(fmt.Sprintf("Memory: %s", humanize.Bytes(mem.Alloc)), 20, rl.Brown)

	}
}
