package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/esensats/sand-sim-automata/grid"
	"github.com/esensats/sand-sim-automata/renderer"
	"github.com/esensats/sand-sim-automata/simulation"
)

const (
	WIDTH  = 40
	HEIGHT = 20
	FPS    = 30
)

func main() {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	sim := simulation.New(WIDTH, HEIGHT, randSource)
	renderer := renderer.New(WIDTH, HEIGHT)

	// Clear screen and hide cursor
	fmt.Print("\033[2J\033[?25l")
	defer fmt.Print("\033[?25h") // Show cursor on exit

	ticker := time.NewTicker(time.Second / FPS)
	defer ticker.Stop()

	// Main game loop
	for range ticker.C {
		// Randomly add sand at top
		if randSource.Float32() < 0.1 {
			sim.Grid.Set(randSource.Intn(WIDTH), 0, grid.Sand)
		}

		sim.Update()
		renderer.Render(sim.Grid)
	}
}
