package objects

import (
	"io/ioutil"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

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
