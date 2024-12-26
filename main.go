package main

// welcome to the controller of the MVC

import (
	"time"

	"wumeno.8/w8_model"
	"wumeno.8/w8_view"
)

var rayl = &w8_view.Raylib{}
var gol = &w8_model.GameOfLife{}

func main() {
	// use tickers for interpreter https://gobyexample.com/tickers
	ticker := time.NewTicker(10 * time.Millisecond)
	gol.Start()

	go func() {
		for {
			select {
			case <-ticker.C:
				gol.UpdateCycle()
				rayl.CopyMatrix(gol.LifeMatrix)
				gol.Terminate()
			}
		}
	}()
	rayl.Start()
}
