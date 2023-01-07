package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	forceY = 2
	forceX = 2
)

func drawWalls() {
	rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.LightGray) // Draw ground
	rl.DrawCube(rl.NewVector3(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)                // Draw a blue wall
	rl.DrawCube(rl.NewVector3(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)                 // Draw a green wall
	rl.DrawCube(rl.NewVector3(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)                 // Draw a yellow wall
	rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
	rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)
}

func drawTux(state *State) {
	frame := 0.5 * math.Pi * rl.GetFrameTime()
	state.RanUntilNow += frame
	speed := state.RanUntilNow - frame

	a, b := 5.0, 5.0
	a += forceY * math.Sin(float64(speed))
	b += forceX * math.Sin(float64(speed))
	ta, tb := float32(a), float32(b)

	rl.DrawCubeTexture(state.Texture, rl.NewVector3(10.0, 10.0, 0.0), ta, tb, ta, rl.White)
}

func Draw3D(state *State) {
	drawTux(state)
	drawWalls()
}
