package objects

import (
	"io"

	"github.com/calebdoxsey/diagrams/graphics"
)

type Arrow struct {
	Line   graphics.Line
	hw, hh float64
}

func NewArrow(line graphics.Line) *Arrow {
	return &Arrow{
		Line: line,
		hw:   4,
		hh:   4,
	}
}

func (obj *Arrow) Render(w io.Writer) {
	render(w, `
  <line
    marker-end='url(#arrow-head)'
    stroke-width='1' fill='none' stroke='black'
    x1="{{.X1}}" x2="{{.X2}}" y1="{{.Y1}}" y2="{{.Y2}}"
    />
`, struct {
		X1, X2, Y1, Y2 float64
	}{
		X1: obj.Line[0].X,
		Y1: obj.Line[0].Y,
		X2: obj.Line[1].X,
		Y2: obj.Line[1].Y,
	})
}
