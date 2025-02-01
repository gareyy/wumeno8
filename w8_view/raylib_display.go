package w8_view

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	con "wumeno.8/constants"
)

type Raylib struct {
	// implements InputOutput interface
	Matrix   [con.WIDTH][con.HEIGHT]bool
	heldKeys [16]bool
	Beep     rl.Sound
}

func (rayl *Raylib) Start() {
	rl.SetTraceLogLevel(rl.LogError) // i dont care about raylib rn
	rl.InitWindow(con.WIDTH*con.PIXEL_SIZE, con.HEIGHT*con.PIXEL_SIZE, "wumeno 8 chip 8 interpreter")
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	rayl.Beep = rl.LoadSound("blipSelect.wav")
	defer rayl.Terminate()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rayl.UpdateMatrix()
		rl.EndDrawing()
		for i, v := range con.KNOWN_KB {
			rayl.heldKeys[i] = rl.IsKeyDown(v)
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
	rl.UnloadSound(rayl.Beep)
	rl.CloseAudioDevice()
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

type inputFunc func([16]bool)

func (rayl *Raylib) TrasmitHeldKeys(fn inputFunc) {
	fn(rayl.heldKeys)
}

func GetHeldKey(key int32) bool {
	return rl.IsKeyDown(key)
}

func GetReleasedKey(key int32) bool {
	return rl.IsKeyReleased(key)
}

func (rayl *Raylib) PlayBeep() {
	rl.PlaySound(rayl.Beep)
}
