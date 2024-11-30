package renderer

import (
	"fmt"
	"strings"

	"github.com/esensats/sand-sim-automata/grid"
)

type Renderer struct {
	Width, Height int
}

func New(width, height int) *Renderer {
	return &Renderer{
		Width:  width,
		Height: height,
	}
}

func (r *Renderer) Render(gr *grid.Grid) {
	fmt.Print("\033[H") // Move cursor to home position
	var sb strings.Builder

	for y := 0; y < gr.Height; y++ {
		for x := 0; x < gr.Width; x++ {
			switch gr.Get(x, y) {
			case grid.Empty:
				sb.WriteString(" ")
			case grid.Sand:
				sb.WriteString("▒")
			}
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}
