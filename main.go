package main

import "github.com/gen2brain/raylib-go/raylib"

const FontSize = 18
const Scale = 1.5
const ParticalRadius = 4
const ParticalSpeed = 2

type Point struct {
	X int32
	Y int32
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(800, 600, "Falling Sand")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var points []Point

	for !rl.WindowShouldClose() {
		mouseX := int32(float32(rl.GetMouseX()) * Scale)
		mouseY := int32(float32(rl.GetMouseY()) * Scale)

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			points = append(points, Point{mouseX, mouseY})
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawCircle(mouseX, mouseY, ParticalRadius, rl.RayWhite)

		for i, point := range points {
			rl.DrawCircle(point.X, point.Y, ParticalRadius, rl.RayWhite)
			if points[i].Y < 400 {
				points[i].Y = point.Y + ParticalSpeed
			}
		}
		rl.EndDrawing()
	}
}
