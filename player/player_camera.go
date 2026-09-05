package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CAMERA_DISTANCE   = 14.0
	MOUSE_SENSITIVITY = 0.003
	MIN_PITCH         = 0.1
	MAX_PITCH         = 1.5
)

type PlayerCamera struct {
	Camera rl.Camera3D
	Yaw    float32
	Pitch  float32
}

func NewPlayerCamera() PlayerCamera {
	return PlayerCamera{
		Camera: rl.Camera3D{
			Position: rl.NewVector3(0, 0, 0),
			Target:   rl.NewVector3(0, 0, 0),
			Up:       rl.NewVector3(0, 1, 0),
			Fovy:     45,
		},
		Yaw:   math.Pi / 2,
		Pitch: math.Pi / 4,
	}
}

func (pc *PlayerCamera) UpdatePlayerCamera(player_position rl.Vector3) {
	mouse_delta := rl.GetMouseDelta()
	pc.Yaw -= mouse_delta.X * MOUSE_SENSITIVITY
	pc.Pitch -= mouse_delta.Y * MOUSE_SENSITIVITY
	if pc.Pitch < MIN_PITCH {
		pc.Pitch = MIN_PITCH
	}
	if pc.Pitch > MAX_PITCH {
		pc.Pitch = MAX_PITCH
	}

	horizontal_distance := float32(math.Cos(float64(pc.Pitch))) * CAMERA_DISTANCE
	pc.Camera.Position.X = player_position.X + float32(math.Cos(float64(pc.Yaw)))*horizontal_distance
	pc.Camera.Position.Y = player_position.Y + float32(math.Sin(float64(pc.Pitch)))*CAMERA_DISTANCE
	pc.Camera.Position.Z = player_position.Z + float32(math.Sin(float64(pc.Yaw)))*horizontal_distance

	pc.Camera.Target = player_position
}

func (pc *PlayerCamera) MovementDirection(input rl.Vector2) rl.Vector2 {
	return rl.NewVector2(
		input.X*float32(math.Sin(float64(pc.Yaw)))+
			input.Y*float32(math.Cos(float64(pc.Yaw))),
		-input.X*float32(math.Cos(float64(pc.Yaw)))+
			input.Y*float32(math.Sin(float64(pc.Yaw))),
	)
}

func (pc *PlayerCamera) FacingDirection() rl.Vector2 {
	return rl.NewVector2(
		-float32(math.Cos(float64(pc.Yaw))),
		-float32(math.Sin(float64(pc.Yaw))),
	)
}
