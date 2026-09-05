package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position     rl.Vector3
	PlayerCamera PlayerCamera
}

func NewPlayer() Player {
	return Player{
		Position:     rl.NewVector3(0, 0, 0),
		PlayerCamera: NewPlayerCamera(),
	}
}

func (p *Player) DrawPlayer() {
	rl.DrawCube(p.Position, 1, 1, 1, rl.Red)
}

func (p *Player) UpdatePlayer(deltaTime float32) {
	if rl.IsKeyDown(rl.KeyW) {
		p.Position.Z -= 5 * deltaTime
	} else if rl.IsKeyDown(rl.KeyS) {
		p.Position.Z += 5 * deltaTime
	} else if rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= 5 * deltaTime
	} else if rl.IsKeyDown(rl.KeyD) {
		p.Position.X += 5 * deltaTime
	}
	p.PlayerCamera.UpdatePlayerCamera(p.Position)
}
