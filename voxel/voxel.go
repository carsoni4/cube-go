package voxel

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const POSITIVE_X = 1
const NEGATIVE_X = 2
const POSITIVE_Y = 3
const NEGATIVE_Y = 4
const POSITIVE_Z = 5
const NEGATIVE_Z = 6

func DrawFace(position rl.Vector3, size float32, color rl.Color, face_direction int8) {
	half_size := size / 2.0
	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	switch face_direction {
	case POSITIVE_X:
		rl.Normal3f(1, 0, 0)
		rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z-half_size)
		rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z+half_size)
		rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z+half_size)
		rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z-half_size)
	case NEGATIVE_X:
		rl.Normal3f(-1, 0, 0)
		rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z+half_size)
		rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z-half_size)
		rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z-half_size)
		rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z+half_size)
	case POSITIVE_Y:
		rl.Normal3f(0, 1, 0)
		rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z-half_size)
		rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z-half_size)
		rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z+half_size)
		rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z+half_size)
	case NEGATIVE_Y:
		rl.Normal3f(0, -1, 0)
		rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z+half_size)
		rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z+half_size)
		rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z-half_size)
		rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z-half_size)
	case POSITIVE_Z:
		rl.Normal3f(0, 0, 1)
		rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z+half_size)
		rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z+half_size)
		rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z+half_size)
		rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z+half_size)
	case NEGATIVE_Z:
		rl.Normal3f(0, 0, -1)
		rl.Vertex3f(position.X+half_size, position.Y-half_size, position.Z-half_size)
		rl.Vertex3f(position.X-half_size, position.Y-half_size, position.Z-half_size)
		rl.Vertex3f(position.X-half_size, position.Y+half_size, position.Z-half_size)
		rl.Vertex3f(position.X+half_size, position.Y+half_size, position.Z-half_size)
	}
	rl.End()
}

func DrawCube(position rl.Vector3, visible [6]bool) {
	for face_direction := POSITIVE_X; face_direction <= NEGATIVE_Z; face_direction++ {
		if visible[face_direction-1] {
			DrawFace(position, 1.0, rl.Green, int8(face_direction))
		}
	}
}
