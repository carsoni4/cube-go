package ui

import (
	"fmt"
	"math"

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

func (ui *DebugUI) DrawFacingDirection(direction rl.Vector2) {
	directionStr := fmt.Sprintf("Facing Direction: (%.2f, %.2f) %s",
		direction.X,
		direction.Y,
		_facingDirectionName(direction))

	rl.DrawText(directionStr, 10, 115, ui.FontSize, ui.TextColor)
}

func _facingDirectionName(direction rl.Vector2) string {
	if math.Abs(float64(direction.X)) >= math.Abs(float64(direction.Y)) {
		if direction.X >= 0 {
			return "X"
		}
		return "-X"
	}

	if direction.Y >= 0 {
		return "Z"
	}
	return "-Z"
}
