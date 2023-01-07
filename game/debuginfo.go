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
		refrate := rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())
		fps := rl.GetFPS()
		state.Text.DrawNextLine(fmt.Sprintf("Frame time: %.5f", rl.GetFrameTime()), 20, rl.Brown)
		state.Text.DrawNextLine(fmt.Sprintf("FPS: %d (%v%%)", fps, (100*fps)/int32(refrate)), 20, rl.Brown)
		state.Text.DrawNextLine(fmt.Sprintf("Memory (Go): %s (Frame GC Enabled: %v)", humanize.Bytes(mem.Alloc), state.FrameGC), 20, rl.Brown)
		state.Text.DrawNextLine(fmt.Sprintf("Position (X, Y, Z): %f/%f/%f", state.Camera.Position.X, state.Camera.Position.Y, state.Camera.Position.Z), 20, rl.Brown)
	}
}
