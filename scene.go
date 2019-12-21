package main

import (
	"fmt"
	"github.com/calebdoxsey/diagrams/draw"
	"github.com/calebdoxsey/diagrams/objects"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"math"
)

var (
	fontFace font.Face
	scene    draw.Group

	imageWidth  = 420.0
	imageHeight = 240.0
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

	scene = append(scene, processorBlock{
		x: 165,
		y: 10,
	})

	reqLine := newArrowLine(180, 40, 100, 60)
	resLine := newArrowLine(100, 80, 180, 70)

	//scene = append(scene, objects.NewKafka(10, 10, imageWidth-20, 50))

	for i := 0; i < 20; i++ {
		// x := float64(34)
		// y := float64(52 + 8*i)
		// obj := &msgBlock{
		// 	number: i + 1,
		// 	x:      x,
		// 	y:      y,
		// }
		obj := objects.NewMessage(i + 1)
		a := draw.NewStatic(obj, 60)
		// if i == 7 {
		// 	a = draw.NewSequence(a,
		// 		draw.NewBasicAnimation(obj, 10, func(v float64) {
		// 			obj.x = x + v*100
		// 			obj.y = y + v*100
		// 		}),
		// 		draw.NewBasicAnimation(obj, 60, func(v float64) {
		// 			obj.pctLoaded = v
		// 		}),
		// 	)
		// }
		scene = append(scene, a)
	}

	scene = append(scene, reqLine, resLine)

}

type arrowLine struct {
	x1, y1, x2, y2 float64
	hw, hh         float64
	text           string
}

func newArrowLine(x1, y1, x2, y2 float64) *arrowLine {
	return &arrowLine{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,

		hw: 5,
		hh: 5,
	}
}

func (obj *arrowLine) Render(ggctx *gg.Context) {
	x1, x2, y1, y2 := obj.x1, obj.x2, obj.y1, obj.y2
	hw, hh := obj.hw, obj.hh

	dx := x2 - x1
	dy := y2 - y1
	angle := math.Atan2(dy, dx)
	length := math.Sqrt(dx*dx + dy*dy)

	ggctx.Push()
	ggctx.SetStrokeStyle(gg.NewSolidPattern(colorBlack))
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

	// text
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBlack))
	if dx < 0 {
		ggctx.Scale(-1, -1)
		ggctx.DrawStringAnchored("test", -length/2, -3, 0.5, 0)
	} else {
		ggctx.DrawStringAnchored("test", length/2, -3, 0.5, 0)
	}

	ggctx.Stroke()
	ggctx.Pop()
}

type processorBlock struct {
	x, y float64
}

func (pb processorBlock) Render(ggctx *gg.Context) {
	x, y := pb.x, pb.y
	const (
		w = 90
		h = 48
	)

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorWhite))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(colorBorder))
	ggctx.DrawRoundedRectangle(align(x), align(y), w, h, 3)
	ggctx.FillPreserve()
	ggctx.Stroke()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBlack))
	ggctx.DrawStringAnchored("processor #1", x+w/2, y+4, 0.5, 1)
	ggctx.Fill()
	ggctx.Pop()

}

type msqQueueBlock struct {
	commit     int
	x, y, w, h float64
}

func newMsgQueueBlock(commit int) *msqQueueBlock {
	return &msqQueueBlock{
		commit: commit,
		x:      10,
		y:      10,
		w:      80,
		h:      220,
	}
}

func (b msqQueueBlock) Render(ggctx *gg.Context) {
	x, y, w, h := b.x, b.y, b.w, b.h

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorWhite))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(colorBorder))
	ggctx.DrawRoundedRectangle(align(x), align(y), w, h, 3)
	ggctx.FillPreserve()
	ggctx.Stroke()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBlack))
	ggctx.DrawStringAnchored("kafka", x+w/2, y+4, 0.5, 1)
	ggctx.DrawStringAnchored(fmt.Sprintf("offset: %02d", b.commit), x+w/2, y+20, 0.5, 1)
	ggctx.Fill()
	ggctx.Pop()

}

type msgBlock struct {
	number    int
	x, y      float64
	pctLoaded float64
}

func (mb msgBlock) Render(ggctx *gg.Context) {
	const (
		w = 32
		h = 16
	)

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorBG))
	ggctx.DrawRoundedRectangle(align(mb.x), align(mb.y), w, h, 3)
	ggctx.Fill()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(colorLightBlue))
	ggctx.DrawRoundedRectangle(align(mb.x), align(mb.y), w*mb.pctLoaded, h, 3)
	ggctx.Fill()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetStrokeStyle(gg.NewSolidPattern(colorBorder))
	ggctx.DrawRoundedRectangle(align(mb.x), align(mb.y), w, h, 3)
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
