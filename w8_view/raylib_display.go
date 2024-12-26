package w8_view

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	con "wumeno.8/constants"
)

type Raylib struct {
	// implements InputOutput interface
	Matrix [con.WIDTH][con.HEIGHT]bool
}

func (rayl *Raylib) Start() {
	rl.InitWindow(con.WIDTH*con.PIXEL_SIZE, con.HEIGHT*con.PIXEL_SIZE, "wumeno 8 chip 8 interpreter")
	rl.SetTargetFPS(60)
	defer rayl.Terminate()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rayl.UpdateMatrix()
		rl.EndDrawing()
		if jj := rl.GetKeyPressed(); jj != 0 {
			fmt.Println(string(jj))
		}
	}
}

func (rayl *Raylib) UpdateMatrix() {
	for i := range con.WIDTH {
		for j := range con.HEIGHT {
			if rayl.Matrix[i][j] {
				rl.DrawRectangle(i*con.PIXEL_SIZE, j*con.PIXEL_SIZE, con.PIXEL_SIZE, con.PIXEL_SIZE, rl.Lime)
			}
		}
	}
}

func (rayl Raylib) Terminate() {
	rl.CloseWindow()
}

func (rayl *Raylib) SetPixel(i, j int, val bool) {
	rayl.Matrix[i][j] = val
}
func (rayl *Raylib) FlipPixel(i, j int) {
	rayl.Matrix[i][j] = !rayl.Matrix[i][j]
}

func (rayl *Raylib) CopyMatrix(newMatrix [con.WIDTH][con.HEIGHT]bool) {
	for v := range con.WIDTH {
		rayl.Matrix[v] = newMatrix[v]
	}
}
