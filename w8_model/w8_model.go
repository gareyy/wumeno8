package w8_model

type Model interface {
	Start()
	UpdateCycle()
	Terminate()
	ReceiveInput([16]bool)
}

/*
REFERENCE: input

Input works by having an array of booleans corresponding to the press or release
of each key, in order of KNOWN_KB

For instruction FX0A, priority is given in order of KNOWN_KB, which relates to
hexadecimal order
*/
