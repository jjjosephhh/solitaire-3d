package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450
	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - models loading")
	camera := rl.NewCamera3D(
		rl.NewVector3(0, 0, 150),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1, 0),
		45,
		rl.CameraProjection(rl.CameraCustom),
	)

	modelCardSpades01 := rl.LoadModel("assets/models/card-spades-01.obj")
	modelCardSpades02 := rl.LoadModel("assets/models/card-spades-02.obj")
	modelCardSpades03 := rl.LoadModel("assets/models/card-spades-03.obj")

	textureCardSpades01 := rl.LoadTexture("assets/images/card-spades-01.png")
	textureCardSpades02 := rl.LoadTexture("assets/images/card-spades-02.png")
	textureCardSpades03 := rl.LoadTexture("assets/images/card-spades-03.png")

	modelCardSpades01.Materials.Maps.Texture = textureCardSpades01
	modelCardSpades02.Materials.Maps.Texture = textureCardSpades02
	modelCardSpades03.Materials.Maps.Texture = textureCardSpades03

	var xAngle float32 = -90
	matrixRotateZ90 := rl.MatrixRotateZ(90 * rl.Deg2rad)

	positionCardSpades01 := rl.NewVector3(0, 0, 0)
	positionCardSpades02 := rl.NewVector3(20, 0, 0)
	positionCardSpades03 := rl.NewVector3(40, 0, 0)

	cardSpades01 := NewCard(&modelCardSpades01, &textureCardSpades01, &positionCardSpades01)
	cardSpades02 := NewCard(&modelCardSpades02, &textureCardSpades02, &positionCardSpades02)
	cardSpades03 := NewCard(&modelCardSpades03, &textureCardSpades03, &positionCardSpades03)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		dt := rl.GetFrameTime()

		matrixRotateX90 := rl.MatrixRotateX(xAngle * rl.Deg2rad)
		matrixRotateXZ90 := rl.MatrixMultiply(matrixRotateX90, matrixRotateZ90)

		modelCardSpades01.Transform = matrixRotateXZ90
		modelCardSpades02.Transform = matrixRotateXZ90
		modelCardSpades03.Transform = matrixRotateXZ90

		// rl.UpdateCamera(&camera, rl.CameraFirstPerson)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		cardSpades01.Draw()
		cardSpades02.Draw()
		cardSpades03.Draw()

		rl.DrawGrid(20, 10)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()

		xAngle += 70 * dt
	}
	rl.UnloadTexture(textureCardSpades01)
	rl.UnloadTexture(textureCardSpades02)
	rl.UnloadTexture(textureCardSpades03)
	rl.UnloadModel(modelCardSpades01)
	rl.UnloadModel(modelCardSpades02)
	rl.UnloadModel(modelCardSpades03)
	rl.CloseWindow()
}

type Card struct {
	Model    *rl.Model
	Texture  *rl.Texture2D
	Position *rl.Vector3
}

func NewCard(model *rl.Model, texture *rl.Texture2D, position *rl.Vector3) *Card {
	return &Card{
		Model:    model,
		Texture:  texture,
		Position: position,
	}
}

func (c *Card) Draw() {
	rl.DrawModel(*c.Model, *c.Position, 1, rl.White)
}
