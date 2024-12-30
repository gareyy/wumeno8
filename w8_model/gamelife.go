package w8_model

import (
	"os"

	con "wumeno.8/constants"
)

type GameOfLife struct {
	// inherits some fucking interface for games or whatever
	LifeMatrix [con.WIDTH][con.HEIGHT]bool
	inputs     [16]bool
}

/*
wraparound count =
 (pos (x or y) + i (delta) + num of rows or cols ) % num of cols or rows

 i.e for rows: row := (y + i + rows) % rows

*/

func (gol *GameOfLife) Start() {
	// make block
	gol.LifeMatrix[10][1] = true
	gol.LifeMatrix[10][2] = true
	gol.LifeMatrix[10][3] = true

	// make spaceship (glider)
	gol.LifeMatrix[3][3] = true
	gol.LifeMatrix[4][3] = true
	gol.LifeMatrix[5][3] = true
	gol.LifeMatrix[5][2] = true
	gol.LifeMatrix[4][1] = true

	// TODO (but never): add ability to pause and play, place cells, etc
}

func (gol *GameOfLife) UpdateCycle() {

	var nextMatrix [con.WIDTH][con.HEIGHT]bool

	for i, row := range gol.LifeMatrix {
		for j, cell := range row {
			total := gol.countNeighbours(i, j)
			nextCell := &nextMatrix[i][j]
			if cell && (total < 2 || total > 3) {
				*nextCell = false
			} else if cell && total >= 2 && total <= 3 {
				*nextCell = true
			} else if !cell && total == 3 {
				*nextCell = true
			}
		}
	}
	// then deep copy
	for i := range con.WIDTH {
		gol.LifeMatrix[i] = nextMatrix[i]
	}
}

func (gol *GameOfLife) Terminate() {
	os.Exit(0)
}

func (gol GameOfLife) countNeighbours(x, y int) int {
	aliveNeighbours := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ { // TODO FIX
			if i == 0 && j == 0 { // prevent counting self
				continue
			}
			col := (x + i + int(con.WIDTH)) % int(con.WIDTH)
			row := (y + j + int(con.HEIGHT)) % int(con.HEIGHT)
			if gol.LifeMatrix[col][row] {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func (gol *GameOfLife) ReceiveInput(received [16]bool) {
	gol.inputs = received
}
