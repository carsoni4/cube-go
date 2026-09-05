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

func (ui *DebugUI) DrawCameraCoords(camera_pos rl.Vector3) {
	// Position String
	posStr := fmt.Sprintf("Camera Position: (%.2f, %.2f, %.2f)",
		camera_pos.X,
		camera_pos.Y,
		camera_pos.Z)

	// Target String
	targetStr := fmt.Sprintf("Camera Target: (%.2f, %.2f, %.2f)",
		camera_pos.X,
		camera_pos.Y,
		camera_pos.Z)

	// Draw the strings on the screen
	rl.DrawText(posStr, 10, 40, ui.FontSize, ui.TextColor)
	rl.DrawText(targetStr, 10, 65, ui.FontSize, ui.TextColor)
}

func (ui *DebugUI) DrawPlayerCoords(player_pos rl.Vector3) {
	// Position String
	posStr := fmt.Sprintf("Player Position: (%.2f, %.2f, %.2f)",
		player_pos.X,
		player_pos.Y,
		player_pos.Z)

	// Draw the string on the screen
	rl.DrawText(posStr, 10, 90, ui.FontSize, ui.TextColor)
}
