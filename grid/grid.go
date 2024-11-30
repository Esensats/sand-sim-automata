package grid

type Cell byte

const (
	Empty Cell = iota
	Sand
)

type Grid struct {
	Width, Height int
	Cells         []Cell
}

func New(width, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
		Cells:  make([]Cell, width*height),
	}
}

func (g *Grid) Get(x, y int) Cell {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return Sand // Treat outside as solid
	}
	return g.Cells[y*g.Width+x]
}

func (g *Grid) Set(x, y int, cell Cell) {
	if x < 0 || x >= g.Width || y >= g.Height {
		return
	}
	g.Cells[y*g.Width+x] = cell
}
