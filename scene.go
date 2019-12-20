package main

import (
	"fmt"
	"github.com/calebdoxsey/diagrams/draw"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"math"
)

var (
	fontFace font.Face
	scene    draw.Group
)

func init() {
	bs, err := ioutil.ReadFile("iosevka-medium.ttf")
	if err != nil {
		panic(err)
	}

	f, err := truetype.Parse(bs)
	if err != nil {
		panic(f)
	}
	fontFace = truetype.NewFace(f, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	scene = append(scene, draw.ObjectFunc(func(ggctx *gg.Context) {
		renderBG(ggctx)
	}))

	for i := 0; i < 20; i++ {
		obj := &msgBlock{
			number: i + 1,
			x:      float64(10),
			y:      float64(40+8*i),
		}
		a := draw.NewStatic(obj, 60)
		if i == 7 {
			a = draw.NewSequence(a,
				draw.NewBasicAnimation(obj, 60, func(v float64) {
					obj.x = 10 + v*100
				}),
			)
		}
		scene = append(scene, a)
	}

	scene = append(scene, commitBlock{
		number: 7,
		x:      10,
		y:      10,
	})

}

type commitBlock struct {
	number int
	x, y   float64
}

func (cb commitBlock) Render(ggctx *gg.Context) {
	const (
		w = 80
		h = 24
	)

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorWhite))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(colorBorder))
	ggctx.DrawRoundedRectangle(align(cb.x), align(cb.y), w, h, 3)
	ggctx.FillPreserve()
	ggctx.Stroke()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBlack))
	ggctx.DrawStringAnchored(fmt.Sprintf("offset: %02d", cb.number), cb.x+w/2, cb.y+h/2, 0.5, 0.3)
	ggctx.Fill()
	ggctx.Pop()

}

type msgBlock struct {
	number int
	x, y   float64
}

func (mb msgBlock) Render(ggctx *gg.Context) {
	const (
		w = 32
		h = 16
	)

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorWhite))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(colorBorder))
	ggctx.DrawRoundedRectangle(align(mb.x), align(mb.y), w, h, 3)
	ggctx.FillPreserve()
	ggctx.Stroke()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBlack))
	ggctx.DrawStringAnchored(fmt.Sprintf("%02d", mb.number), mb.x+w/2, mb.y+h/2, 0.5, 0.3)
	ggctx.Fill()
	ggctx.Pop()

}

func renderBG(ggctx *gg.Context) {
	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBG))
	ggctx.Clear()
	ggctx.Pop()
}

func align(v float64) float64 {
	return math.Round(v*2) / 2
}
