package draw

import "github.com/fogleman/gg"

type Object interface {
	Render(ggctx *gg.Context)
}

type ObjectFunc func(ggctx *gg.Context)

func (f ObjectFunc) Render(ggctx *gg.Context) {
	f(ggctx)
}

type Group []Object

func (g Group) Frames() int {
	var max int
	for _, obj := range g {
		if a, ok := obj.(Animation); ok {
			fs := a.Frames()
			if fs > max {
				max = fs
			}
		}
	}
	return max
}

func (g Group) Render(ggctx *gg.Context) {
	for _, obj := range g {
		obj.Render(ggctx)
	}
}

func (g Group) Update(pct float64) {
	for _, obj := range g {
		if a, ok := obj.(Animation); ok {
			a.Update(pct)
		}
	}
}
