package objects

import (
	"html/template"
	"io"
	"io/ioutil"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Object interface {
	Render(w io.Writer)
}

var fontFace font.Face

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
		Size:    11,
		DPI:     72,
		Hinting: font.HintingNone,
	})
}

func align(v float64) float64 {
	return math.Round(v*2) / 2
}

func render(w io.Writer, tpl string, data interface{}) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
