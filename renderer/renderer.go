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

var cellChars = []string{" ", "▒"}

func (r *Renderer) Render(gr *grid.Grid) {
	fmt.Print("\033[H") // Move cursor to home position
	var sb strings.Builder

	for y := 0; y < gr.Height; y++ {
		for x := 0; x < gr.Width; x++ {
			sb.WriteString(cellChars[gr.Get(x, y)])
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}
