package main

// welcome to the controller of the MVC

import (
	"time"

	"wumeno.8/w8_model"
	"wumeno.8/w8_view"
)

var rayl = &w8_view.Raylib{}

var interpreter = &w8_model.Interpreter{}

func main() {
	// use tickers for interpreter https://gobyexample.com/tickers
	modelTick := time.NewTicker(10 * time.Millisecond)
	inputTick := time.NewTicker(time.Second / 60)
	done := make(chan bool)
	interpreter.Start()

	go func() {
		for {
			select {
			case <-done:
				return
			case <-modelTick.C:
				interpreter.UpdateCycle()
				rayl.CopyMatrix(interpreter.DisplayMatrix)
			case <-inputTick.C:
				rayl.TrasmitHeldKeys(interpreter.ReceiveInput)
			}
		}
	}()

	rayl.Start()
	done <- true
}
