package objects

import (
	"math"

	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/fogleman/gg"
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

func (obj *Arrow) Render(ggctx *gg.Context) {
	x1, x2, y1, y2 := obj.Line[0].X, obj.Line[1].X, obj.Line[0].Y, obj.Line[1].Y
	hw, hh := obj.hw, obj.hh

	dx := x2 - x1
	dy := y2 - y1
	angle := math.Atan2(dy, dx)
	length := math.Sqrt(dx*dx + dy*dy)

	ggctx.Push()
	ggctx.SetStrokeStyle(gg.NewSolidPattern(Color(0x000000)))
	ggctx.SetLineCap(gg.LineCapButt)
	ggctx.Translate(x1, y1)
	ggctx.Rotate(angle)

	ggctx.MoveTo(0, 0)
	ggctx.LineTo(length, 0)

	// start arrow
	//ggctx.MoveTo(hh, -hw)
	//ggctx.LineTo(0, 0)
	//ggctx.LineTo(hh, hw)

	// end arrow
	ggctx.MoveTo(length-hh, -hw)
	ggctx.LineTo(length, 0)
	ggctx.LineTo(length-hh, hw)

	ggctx.Stroke()
	ggctx.Pop()
}
