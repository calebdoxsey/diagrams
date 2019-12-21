package objects

import "github.com/fogleman/gg"

type Background struct {
	w, h float64
}

func NewBackground(w, h float64) *Background {
	return &Background{
		w: w,
		h: h,
	}
}

func (obj *Background) Render(ggctx *gg.Context) {
	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(Color(0xF5F2F0)))
	ggctx.Clear()
	ggctx.Pop()
}
