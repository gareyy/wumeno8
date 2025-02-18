package main

// welcome to the controller of the MVC

import (
	"errors"
	"os"
	"time"

	"wumeno.8/w8_model"
	"wumeno.8/w8_view"
)

var rayl = &w8_view.Raylib{}

var interpreter = &w8_model.Interpreter{}

func main() {
	modelTick := time.NewTicker(time.Second / 2000)
	inputTick := time.NewTicker(time.Second / 60)
	done := make(chan bool)
	args := os.Args[1:]
	if len(args) != 1 {
		panic(errors.New("No rom provided or too many arguments"))
	}
	filename := args[0]
	program, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	interpreter.Start(program)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-modelTick.C:
				rayl.TrasmitHeldKeys(interpreter.ReceiveInput)
				interpreter.UpdateCycle()
				rayl.CopyMatrix(interpreter.DisplayMatrix)
			case <-inputTick.C:
				interpreter.TimerUpdate(rayl.PlayBeep)
			}
		}
	}()

	rayl.Start()
	done <- true
}
