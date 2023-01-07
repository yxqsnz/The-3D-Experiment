package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	IsShowingDebugInfo bool
	Text               TextManager
	Camera             rl.Camera3D
}

func InitializeState() State {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 2.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.8, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	return State{Text: TextManager{}, Camera: camera}
}
