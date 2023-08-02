package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450
	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - models loading")
	camera := rl.NewCamera3D(
		rl.NewVector3(0, 0, 100),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1, 0),
		45,
		rl.CameraProjection(rl.CameraCustom),
	)

	model1 := rl.LoadModel("assets/models/card-spades-01.obj")
	model2 := rl.LoadModel("assets/models/card-spades-02.obj")

	texture1 := rl.LoadTexture("assets/images/card-spades-01.png")
	texture2 := rl.LoadTexture("assets/images/card-spades-02.png")

	model1.Materials.Maps.Texture = texture1
	model2.Materials.Maps.Texture = texture2

	var xAngle float32 = -90
	matrixRotateZ90 := rl.MatrixRotateZ(90 * rl.Deg2rad)

	position1 := rl.NewVector3(0, 0, 0)
	position2 := rl.NewVector3(20, 0, 0)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		matrixRotateX90 := rl.MatrixRotateX(xAngle * rl.Deg2rad)
		matrixRotateXZ90 := rl.MatrixMultiply(matrixRotateX90, matrixRotateZ90)

		model1.Transform = matrixRotateXZ90
		model2.Transform = matrixRotateXZ90

		// rl.UpdateCamera(&camera, rl.CameraFirstPerson)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		rl.DrawModel(model1, position1, 1, rl.White)
		rl.DrawModel(model2, position2, 1, rl.White)

		rl.DrawGrid(20, 10)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()

		xAngle += 70 * dt
	}
	rl.UnloadTexture(texture1)
	rl.UnloadTexture(texture2)
	rl.UnloadModel(model1)
	rl.UnloadModel(model2)
	rl.CloseWindow()
}
