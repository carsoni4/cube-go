package main

import (
	ui "github.com/carsoni4/cube-go/ui"

	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Hello Carson :P")
	rl.InitWindow(1920, 1200, "Cube-go")
	rl.SetTargetFPS(60)

	camera := rl.Camera3D{
		Position: rl.NewVector3(0, 10, 10),
		Target:   rl.NewVector3(0, 0, 0),
		Up:       rl.NewVector3(0, 1, 0),
		Fovy:     45,
	}

	debugUI := ui.NewDebugUI(rl.Red, 32)

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginDrawing()
		debugUI.DrawCameraCoords(camera)
		rl.BeginMode3D(camera)
		rl.DrawGrid(16, 1)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
