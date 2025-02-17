package main

// welcome to the controller of the MVC

import (
	"os"
	"time"

	"wumeno.8/w8_model"
	"wumeno.8/w8_view"
)

var rayl = &w8_view.Raylib{}

var interpreter = &w8_model.Interpreter{}

func main() {
	// use tickers for interpreter https://gobyexample.com/tickers
	modelTick := time.NewTicker(time.Second / 2000)
	inputTick := time.NewTicker(time.Second / 60)
	done := make(chan bool)
	program, err := os.ReadFile("roms/6-keypad.ch8")
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
