package w8_model

type Model interface {
	Start()
	UpdateCycle()
	Terminate()
	// SOON: figure out how to do inputs
}
