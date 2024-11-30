package simulation

import (
	"math/rand"

	"github.com/esensats/sand-sim-automata/grid"
)

type Simulation struct {
	Grid *grid.Grid
	rand *rand.Rand
}

func New(width, height int, randSource *rand.Rand) *Simulation {
	return &Simulation{
		Grid: grid.New(width, height),
		rand: randSource,
	}
}

func (s *Simulation) Update() {
	// Update from bottom to top, right to left
	for y := s.Grid.Height - 1; y >= 0; y-- {
		for x := s.Grid.Width - 1; x >= 0; x-- {
			if s.Grid.Get(x, y) == grid.Sand {
				s.updateSand(x, y)
			}
		}
	}
}

func (s *Simulation) updateSand(x, y int) {
	// Try to move down
	if s.Grid.Get(x, y+1) == grid.Empty {
		s.Grid.Set(x, y, grid.Empty)
		s.Grid.Set(x, y+1, grid.Sand)
		return
	}

	// Try to move diagonally
	direction := []int{-1, 1}[s.rand.Intn(2)]
	if s.Grid.Get(x+direction, y+1) == grid.Empty {
		s.Grid.Set(x, y, grid.Empty)
		s.Grid.Set(x+direction, y+1, grid.Sand)
	}
}
