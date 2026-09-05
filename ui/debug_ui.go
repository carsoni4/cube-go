package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DebugUI struct {
	TextColor rl.Color
	FontSize  int32
}

func NewDebugUI(color rl.Color, size int32) DebugUI {
	return DebugUI{
		TextColor: color,
		FontSize:  size,
	}
}

func (ui *DebugUI) DrawCameraCoords(camera rl.Camera3D) {
	// Position String
	posStr := fmt.Sprintf("Camera Position: (%.2f, %.2f, %.2f)",
		camera.Position.X,
		camera.Position.Y,
		camera.Position.Z)

	// Target String
	targetStr := fmt.Sprintf("Camera Target: (%.2f, %.2f, %.2f)",
		camera.Target.X,
		camera.Target.Y,
		camera.Target.Z)

	// Draw the strings on the screen
	rl.DrawText(posStr, 10, 40, ui.FontSize, ui.TextColor)
	rl.DrawText(targetStr, 10, 65, ui.FontSize, ui.TextColor)
}
