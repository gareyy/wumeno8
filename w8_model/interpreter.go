package w8_model

import (
	con "wumeno.8/constants"
)

type Interpreter struct {
	memory         [4096]byte
	DisplayMatrix  [con.WIDTH][con.HEIGHT]bool
	registerV      [16]byte
	indexRegister  uint16
	programCounter uint16
	delayTimer     byte
	soundTimer     byte
	stack          [16]uint16
	stackPointer   uint16
	inputs         [16]bool
}

func (in *Interpreter) Start() {
}

func (in *Interpreter) UpdateCycle() {

}

func (in *Interpreter) ReceiveInput(received [16]bool) {

}

func (in *Interpreter) Terminate() {

}
