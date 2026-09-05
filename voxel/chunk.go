package voxel

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func isSolid(x, y, z int) bool {
	return x >= -8 && x < 8 &&
		y >= -8 && y < 8 &&
		z >= -8 && z < 8
}

func DrawChunk() {
	for x := -8; x < 8; x++ {
		for y := -8; y < 8; y++ {
			for z := -8; z < 8; z++ {
				visible := [6]bool{
					!isSolid(x-1, y, z), // POSITIVE_X
					!isSolid(x+1, y, z), // NEGATIVE_X
					!isSolid(x, y-1, z), // POSITIVE_Y
					!isSolid(x, y+1, z), // NEGATIVE_Y
					!isSolid(x, y, z+1), // POSITIVE_Z
					!isSolid(x, y, z-1), // NEGATIVE_Z
				}
				DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), visible)
			}
		}
	}
}
