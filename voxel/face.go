package voxel

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawFace(position rl.Vector3, size float32, color rl.Color, face_direction rl.Vector3) {
	half_size := size / 2.0
	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)
	rl.Normal3f(face_direction.X, face_direction.Y, face_direction.Z)

	// Four Verticies
	rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z+half_size)
	rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z+half_size)
	rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z+half_size)
	rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z+half_size)
	rl.End()
}
