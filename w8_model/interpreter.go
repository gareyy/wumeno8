package w8_model

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"

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

/*
0x000-0x1FF - Chip 8 interpreter (contains font set in emu)
0x050-0x0A0 - Used for the built in 4x5 pixel font set (0-F)
0x200-0xFFF - Program ROM and work RAM

big endian (so for test_opcode.ch8, 124e is the first opcode, 13dc is last)
*/

func (in *Interpreter) Start() {
	// empty them all, just in case
	in.memory = [4096]byte{}
	in.DisplayMatrix = [con.WIDTH][con.HEIGHT]bool{}
	in.registerV = [16]byte{}
	in.indexRegister, in.stackPointer = 0, 0
	in.programCounter = 0x200
	in.delayTimer, in.soundTimer = 0x0, 0x0
	in.stack = [16]uint16{}
	in.inputs = [16]bool{}

	// load fontset
	for i := 0; i < len(con.FONTSET); i++ {
		in.memory[i] = con.FONTSET[i]
	}

	// load program
	program, err := os.ReadFile("6-keypad.ch8")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(program); i++ {
		in.memory[i+int(in.programCounter)] = program[i]
	}

}

func (in *Interpreter) UpdateCycle() {
	// get opcode
	opcode := uint16(in.memory[in.programCounter])<<8 | uint16(in.memory[in.programCounter+1])
	fmt.Printf("0x%.3X: %.4X\n", in.programCounter, opcode)
	increase := 2 // default
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode & 0x0FFF {
		case 0x00E0:
			// clear screen
			in.DisplayMatrix = [con.WIDTH][con.HEIGHT]bool{}
		case 0x00EE:
			//return from subroutine
			increase = 2
			in.stackPointer--
			in.programCounter = in.stack[in.stackPointer]
		default:
			// execute machine subroutine, ignore
			break
		}
	case 0x1000:
		// jump
		increase = 0
		in.programCounter = opcode & 0x0FFF
		//fmt.Printf("jumped to %.3X\n", in.programCounter)
	case 0x2000:
		// execute subroutine at NNN
		in.stack[in.stackPointer] = in.programCounter
		increase = 0
		in.stackPointer++
		in.programCounter = opcode & 0x0FFF
		//fmt.Printf("subroutined to %.3X\n", in.programCounter)
	case 0x3000:
		// skip next instruction if vx == nn
		vx := in.registerV[(opcode&0x0F00)>>8]
		nn := byte(opcode & 0x00FF)
		if vx == nn {
			increase = 4
		}
	case 0x4000:
		// skip next instruction if vx != nn
		vx := in.registerV[(opcode&0x0F00)>>8]
		nn := byte(opcode & 0x00FF)
		if vx != nn {
			increase = 4
		}
	case 0x5000:
		// skip next instruction if vx == vy
		vx := in.registerV[(opcode&0x0F00)>>8]
		vy := in.registerV[(opcode&0x00F0)>>4]
		if vx == vy {
			increase = 4
		}
	case 0x6000:
		// set reg X to NN
		register := (opcode & 0x0F00) >> 8
		val := byte(opcode & 0x00FF)
		in.registerV[register] = val
	case 0x7000:
		// add value to register X
		register := (opcode & 0x0F00) >> 8
		val := byte(opcode & 0x00FF)
		in.registerV[register] += val
	case 0x8000:
		// uh oh
		switch opcode & 0x000F {
		case 0x0000:
			in.registerV[(opcode&0x0F00)>>8] = in.registerV[(opcode&0x00F0)>>4]
		case 0x0001:
			in.registerV[(opcode&0x0F00)>>8] |= in.registerV[(opcode&0x00F0)>>4]
		case 0x0002:
			in.registerV[(opcode&0x0F00)>>8] &= in.registerV[(opcode&0x00F0)>>4]
		case 0x0003:
			in.registerV[(opcode&0x0F00)>>8] ^= in.registerV[(opcode&0x00F0)>>4]
		case 0x0004:
			vx := &in.registerV[(opcode&0x0F00)>>8]
			vy := &in.registerV[(opcode&0x00F0)>>4]
			if int(*vx)+int(*vy) > 255 {
				in.registerV[0xF] = 1
			}
			*vx += *vy
		case 0x0005:
			vx := &in.registerV[(opcode&0x0F00)>>8]
			vy := &in.registerV[(opcode&0x00F0)>>4]
			if int(*vx) > int(*vy) {
				in.registerV[0xF] = 1
			} else {
				in.registerV[0xF] = 0
			}
			*vx -= *vy
		case 0x0007:
			vx := &in.registerV[(opcode&0x0F00)>>8]
			vy := &in.registerV[(opcode&0x00F0)>>4]
			if int(*vx) < int(*vy) {
				in.registerV[0xF] = 1
			} else {
				in.registerV[0xF] = 0
			}
			*vx = *vy - *vx
		case 0x0006:
			// right shift
			vx := &in.registerV[(opcode&0x0F00)>>8]
			vy := in.registerV[(opcode&0x00F0)>>4]
			in.registerV[0xF] = vy & 0x01
			*vx = vy >> 1
		case 0x000E:
			// left shift
			vx := &in.registerV[(opcode&0x0F00)>>8]
			vy := in.registerV[(opcode&0x00F0)>>4]
			in.registerV[0xF] = vy & 0x80
			*vx = vy << 1
		default:
			fmt.Printf("Unknown: %.4X at pc %.3X\n", opcode, in.programCounter)
			break
		}
	case 0x9000:
		// skip next instruction if vx != vy
		vx := in.registerV[(opcode&0x0F00)>>8]
		vy := in.registerV[(opcode&0x00F0)>>4]
		if vx != vy {
			increase = 4
		}
	case 0xA000:
		// set I to NNN
		in.indexRegister = opcode & 0x0FFF
	case 0xB000:
		// jump with offset, uses cosmac implementation
		vx := in.registerV[0x0]
		increase = 0
		in.programCounter = (opcode & 0x0FFF) + uint16(vx)
		//fmt.Printf("jumped to %.3X\n", in.programCounter)
	case 0xC000:
		// set X to random
		in.registerV[(opcode&0x0F00)>>8] = byte((opcode & 0x00FF) & uint16(rand.IntN(100)))
	case 0xD000:
		// draw time!!
		x := int32(in.registerV[int32((opcode&0x0F00)>>8)]) % con.WIDTH
		y := int32(in.registerV[int32((opcode&0x00F0)>>4)]) % con.HEIGHT
		height := opcode & 0x000F
		in.registerV[0xF] = 0
		for yoff := range int32(height) {
			line := byteToBoolArray(in.memory[in.indexRegister+uint16(yoff)])
			for xoff := range int32(8) {
				if x+int32(xoff) >= con.WIDTH || y+int32(yoff) >= con.HEIGHT {
					continue
				}
				pos := &in.DisplayMatrix[x+int32(xoff)][y+int32(yoff)]
				if *pos && line[xoff] {
					in.registerV[0xF] = 1
				}
				*pos = bool(line[xoff] != *pos)
			}
		}
	case 0xE000:
		switch opcode & 0x00FF {
		case 0x009E:
			vx := in.registerV[(opcode&0x0F00)>>8]
			key := in.inputs[vx]
			if key {
				increase = 4
			}
		case 0x00A1:
			vx := in.registerV[(opcode&0x0F00)>>8]
			key := in.inputs[vx]
			if !key {
				increase = 4
			}
		}
	case 0xF000:
		fmt.Printf("0x%.3X: %.4X\n", in.programCounter, opcode)
		switch opcode & 0x00FF {
		case 0x0007:
			in.registerV[(opcode&0x0F00)>>8] = in.delayTimer
		case 0x000A:
			increase = 0
			for i := range byte(0xF) {
				if in.inputs[i] {
					in.registerV[(opcode&0x0F00)>>8] = i
					increase = 2
					break
				}
			}
		case 0x0015:
			in.delayTimer = in.registerV[(opcode&0x0F00)>>8]
		case 0x0018:
			in.soundTimer = in.registerV[(opcode&0x0F00)>>8]
		case 0x001E:
			in.indexRegister += uint16(in.registerV[(opcode&0x0F00)>>8])
		case 0x0029:
			// fontset starts from mem address 0
			vx := in.registerV[(opcode&0x0F00)>>8]
			fmt.Println(vx)
			in.indexRegister = uint16(5 * vx)
			fmt.Printf("Hex to %.3X", in.indexRegister)
		case 0x0033:
			vx := in.registerV[(opcode&0x0F00)>>8]
			in.memory[in.indexRegister] = vx / 100
			in.memory[in.indexRegister+1] = (vx / 10) % 10
			in.memory[in.indexRegister+2] = (vx % 100) % 10
			for i := range uint16(3) {
				fmt.Println(in.memory[in.indexRegister+i])
			}
		case 0x0055:
			// use temp offset, save
			totalregisters := uint16((opcode & 0x0F00) >> 8)
			for i := range totalregisters + 1 {
				in.memory[in.indexRegister+i] = in.registerV[i]
			}
		case 0x0065:
			// load
			fmt.Println("FF")
			totalregisters := uint16((opcode & 0x0F00) >> 8)
			for i := range totalregisters + 1 {
				in.registerV[i] = in.memory[in.indexRegister+i]
			}
		}
	default:
		fmt.Printf("Unknown: %.4X at pc %.3X\n", opcode, in.programCounter)
		break
	}

	// then increase
	in.programCounter += uint16(increase)
	if in.programCounter < 0x200 || in.programCounter > 0xFFF {
		log.Panicf("Invalid position for pc %.3X\n", in.programCounter)
	}
}

func (in *Interpreter) ReceiveInput(received [16]bool) {
	in.inputs = received
	in.TimerUpdate() // because input is at 60hz, we also do this
}

func (in *Interpreter) Terminate() {
	os.Exit(0)
}

func (in *Interpreter) TimerUpdate() {
	// timers
	if in.delayTimer > 0 {
		in.delayTimer--
	}
	if in.soundTimer > 0 {
		if in.soundTimer == 1 {
			fmt.Println("BEEP!!!!")
		}
		in.soundTimer--
	}
}

func byteToBoolArray(b byte) [8]bool {
	var new [8]bool
	for i := range 8 {
		focus := b & 0b10000000
		new[i] = bool(focus > 0)
		b = b << 1
	}
	return new
}
