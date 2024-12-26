package w8_view

type InputOutput interface {
	Start()
	Terminate()
	SetPixel(i, j int, val bool)
	FlipPixel(i, j int)
}
