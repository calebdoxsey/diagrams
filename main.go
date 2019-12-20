package main

import (
	"image/color"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

var (
	colorBG = color.RGBA{
		R: 0xF5,
		G: 0xF2,
		B: 0xF0,
		A: 0xFF,
	}
	colorBorder = color.RGBA{
		R: 0x66,
		G: 0x66,
		B: 0x66,
		A: 0xFF,
	}
	colorWhite = color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0xFF,
		A: 0xFF,
	}
	colorBlack = color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0x00,
		A: 0xFF,
	}
	colorLightBlue = color.RGBA{
		R: 0xcf,
		G: 0xe2,
		B: 0xf3,
		A: 0xff,
	}
)

func main() {
	log.SetFlags(0)

	_ = os.MkdirAll("out", 0755)

	apngname := filepath.Join(os.TempDir(), "tmp.apng")
	h264name := filepath.Join(os.TempDir(), "tmp.mp4")
	defer os.Remove(apngname)
	defer os.Remove(h264name)
	//av1name := "example.av1.mp4"

	apngf, err := os.Create(apngname)
	if err != nil {
		log.Fatalln(err)
	}
	err = render(apngf)
	apngf.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var eg errgroup.Group
	eg.Go(func() error {
		return apngToH264(h264name, apngname)
	})
	//eg.Go(func() error {
	//	return apngToAV1(av1name, apngname)
	//})
	err = eg.Wait()
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Rename(h264name, "./out/example.h264.mp4")
	if err != nil {
		log.Fatalln(err)
	}
}
