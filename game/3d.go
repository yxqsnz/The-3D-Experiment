package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Draw3D(state *State) {
	rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.Green) // Draw ground

}
