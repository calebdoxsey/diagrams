package objects

import "image/color"

func Color(rgb uint64) color.RGBA {
	r := byte((rgb >> 16) & 0xFF)
	g := byte((rgb >> 8) & 0xFF)
	b := byte((rgb >> 0) & 0xFF)
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 0xFF,
	}
}

func GetAlpha(pct float64) byte {
	a := pct * 255
	i := int(a)
	switch {
	case i >= 255:
		return 255
	case i <= 0:
		return 0
	default:
		return byte(i)
	}
}
