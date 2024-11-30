package renderer

import (
	"strings"
	"testing"

	"github.com/esensats/sand-sim-automata/grid"
)

func BenchmarkRenderWithCellChars(b *testing.B) {
	gr := grid.New(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderWithArrayMap(gr)
	}
}

func BenchmarkRenderWithRunes(b *testing.B) {
	gr := grid.New(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderWithArrayMapRune(gr)
	}
}

func BenchmarkRenderWithSwitchCase(b *testing.B) {
	gr := grid.New(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderWithSwitchCase(gr)
	}
}

func BenchmarkRenderWithSwitchCaseRune(b *testing.B) {
	gr := grid.New(100, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderWithSwitchCaseRune(gr)
	}
}

var cellStrs = []string{" ", "▒"}

func renderWithArrayMap(gr *grid.Grid) string {
	var sb strings.Builder

	for y := 0; y < gr.Height; y++ {
		for x := 0; x < gr.Width; x++ {
			sb.WriteString(cellStrs[gr.Get(x, y)])
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

var cellRunes = []rune{' ', '▒'}

func renderWithArrayMapRune(gr *grid.Grid) string {
	var sb strings.Builder

	for y := 0; y < gr.Height; y++ {
		for x := 0; x < gr.Width; x++ {
			sb.WriteRune(cellRunes[gr.Get(x, y)])
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func renderWithSwitchCase(gr *grid.Grid) string {
	var sb strings.Builder

	for y := 0; y < gr.Height; y++ {
		for x := 0; x < gr.Width; x++ {
			// switch cell {
			// case grid.Empty:
			// 	sb.WriteString(" ")
			// case grid.Sand:
			// 	sb.WriteString("▒")
			// }
			switch gr.Get(x, y) {
			case grid.Empty:
				sb.WriteString(" ")
			case grid.Sand:
				sb.WriteString("▒")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func renderWithSwitchCaseRune(gr *grid.Grid) string {
	var sb strings.Builder

	for y := 0; y < gr.Height; y++ {
		for x := 0; x < gr.Width; x++ {
			switch gr.Get(x, y) {
			case grid.Empty:
				sb.WriteByte(' ')
			case grid.Sand:
				sb.WriteRune('▒')
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
