package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

const WindowWidth = 800
const WindowHeight = 600
const InitAmount = 800
const SpawnRate = 2

type Point struct {
	x      float32
	y      float32
	speed  float32
	radius float32
}

func newPoint(x, y float32) Point {
	var speed float32 = rand.Float32() + 1
	var radius float32 = speed * 0.5
	return Point{
		x:      x,
		y:      y,
		speed:  speed,
		radius: radius,
	}
}

func newTopPoint() Point {
	x := rand.Int31n(WindowWidth)
	return newPoint(float32(x), 0)
}

func newInitPoint() Point {
	x := rand.Int31n(WindowWidth)
	y := rand.Int31n(WindowHeight)
	return newPoint(float32(x), float32(y))
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.SetTraceLogLevel(rl.LogNone)
	rl.SetTargetFPS(60)

	BackgroundColor := rl.NewColor(10, 11, 39, 255)
	rl.InitWindow(WindowWidth, WindowHeight, "Falling Snow")
	defer rl.CloseWindow()

	var points []Point
	for range InitAmount {
		points = append(points, newInitPoint())
	}

	for !rl.WindowShouldClose() {
		for range SpawnRate {
			points = append(points, newTopPoint())
		}

		rl.BeginDrawing()
		rl.ClearBackground(BackgroundColor)

		for _, point := range points {
			rl.DrawCircle(int32(point.x), int32(point.y), point.radius, rl.RayWhite)
		}
		rl.EndDrawing()

		var next []Point
		for i, point := range points {
			if points[i].y < WindowHeight {
				point.y = point.y + point.speed
				next = append(next, point)
			}
		}
		points = next
	}
}
