package main

import (
	ui "github.com/carsoni4/cube-go/ui"

	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	plyr "github.com/carsoni4/cube-go/player"

	vox "github.com/carsoni4/cube-go/voxel"
)

func main() {
	fmt.Println("Hello Carson :P")
	rl.InitWindow(1920, 1200, "Cube-go")
	rl.SetTargetFPS(60)
	player := plyr.NewPlayer()

	debugUI := ui.NewDebugUI(rl.Red, 32)

	rl.DisableCursor()
	for !rl.WindowShouldClose() {
		player.UpdatePlayer(rl.GetFrameTime())

		rl.ClearBackground(rl.SkyBlue)
		rl.BeginDrawing()
		debugUI.DrawCameraCoords(player.PlayerCamera.Camera.Position)
		debugUI.DrawPlayerCoords(player.Position)
		debugUI.DrawFacingDirection(player.PlayerCamera.FacingDirection())
		rl.BeginMode3D(player.PlayerCamera.Camera)
		vox.DrawFace(rl.Vector3{0, 0, 0}, 1, rl.Green, rl.Vector3{1, 0, 0})
		vox.DrawFace(rl.Vector3{0, 0, 0}, 1, rl.Green, rl.Vector3{0, 1, 0})
		vox.DrawFace(rl.Vector3{0, 0, 0}, 1, rl.Green, rl.Vector3{0, 0, 1})
		vox.DrawFace(rl.Vector3{0, 0, 0}, 1, rl.Green, rl.Vector3{0, 0, -1})
		vox.DrawFace(rl.Vector3{0, 0, 0}, 1, rl.Green, rl.Vector3{-1, 0, 0})
		vox.DrawFace(rl.Vector3{0, 0, 0}, 1, rl.Green, rl.Vector3{0, -1, 0})

		player.DrawPlayer()
		rl.DrawGrid(16, 1)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
