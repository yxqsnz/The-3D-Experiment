package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextManager struct {
	yPosition int32
	padding   int32
}

func (tm *TextManager) DrawNextLine(content string, size int32, color color.RGBA) {
	rl.DrawText(content, tm.padding, tm.yPosition, size, color)
	tm.yPosition += size
}

func (tm *TextManager) Reset() {
	tm.padding = 10
	tm.yPosition = 10
}
