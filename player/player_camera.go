package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerCamera struct {
	Camera rl.Camera3D
}

func NewPlayerCamera() PlayerCamera {
	return PlayerCamera{
		Camera: rl.Camera3D{
			Position: rl.NewVector3(0, 0, 0),
			Target:   rl.NewVector3(0, 0, 0),
			Up:       rl.NewVector3(0, 1, 0),
			Fovy:     45,
		},
	}
}

func (pc *PlayerCamera) UpdatePlayerCamera(player_position rl.Vector3) {
	// Keeping Camera Offsets
	pc.Camera.Position.X = player_position.X
	pc.Camera.Position.Y = player_position.Y + 5
	pc.Camera.Position.Z = player_position.Z + 5

	// Keeping Camera Target
	pc.Camera.Target.X = player_position.X
	pc.Camera.Target.Y = player_position.Y
	pc.Camera.Target.Z = player_position.Z
}
