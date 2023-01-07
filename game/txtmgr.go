package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextManager struct {
	yPosition int
}

func (tm *TextManager) DrawNextLine(content string, size int, color color.RGBA) {
	rl.DrawText(content, 10, int32(tm.yPosition), int32(size), color)
	tm.yPosition += size
}

func (tm *TextManager) Reset() {
	tm.yPosition = 10
}
