package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const MOVEMENT_SPEED = 5.0
const GRAVITY = 20.0
const JUMP_FORCE = 8.0

type Player struct {
	Position     rl.Vector3
	Velocity     rl.Vector3
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
	rl.DrawCubeWires(p.Position, 1, 1, 1, rl.Black)
}

func (p *Player) IsGrounded() bool {
	if p.Position.Y <= 0 {
		p.Position.Y = 0
		return true
	}
	return false
}
func (p *Player) UpdatePlayer(deltaTime float32) {
	direction := rl.NewVector2(0, 0)
	if rl.IsKeyDown(rl.KeyW) {
		direction.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		direction.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		direction.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		direction.X += 1
	}
	if p.IsGrounded() {
		p.Velocity.Y = 0
		if rl.IsKeyDown(rl.KeySpace) {
			p.Velocity.Y = JUMP_FORCE
		}
	} else {
		p.Velocity.Y -= GRAVITY * deltaTime
	}

	normalizedDirection := rl.Vector2Normalize(direction)
	p.Velocity.X = normalizedDirection.X * MOVEMENT_SPEED
	p.Velocity.Z = normalizedDirection.Y * MOVEMENT_SPEED
	p.Position.X += p.Velocity.X * deltaTime
	p.Position.Y += p.Velocity.Y * deltaTime
	p.Position.Z += p.Velocity.Z * deltaTime
	if p.Position.Y < 0 {
		p.Position.Y = 0
		p.Velocity.Y = 0
	}

	p.PlayerCamera.UpdatePlayerCamera(p.Position)
}
