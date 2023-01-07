package game

import (
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Setup(state *State) {
	rl.SetCameraMode(state.Camera, rl.CameraFirstPerson)
}

func PreDraw(state *State) {
	if rl.IsKeyPressed(rl.KeyF3) {
		state.IsShowingDebugInfo = !state.IsShowingDebugInfo
	}

	if rl.IsKeyPressed(rl.KeyF10) {
		state.FrameGC = !state.FrameGC
	}

	if state.FrameGC {
		runtime.GC()
	}

	state.Text.Reset()
}

func Draw(state *State) {
	rl.UpdateCamera(&state.Camera)
	rl.ClearBackground(rl.Black)

	rl.BeginMode3D(state.Camera)
	Draw3D(state)
	rl.EndMode3D()

	rl.DrawCircle(int32(rl.GetScreenWidth())/2, int32(rl.GetScreenHeight())/2, 2, rl.RayWhite)

	DrawDebugInfo(state)
}
