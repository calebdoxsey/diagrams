package main

import (
	"github.com/fogleman/gg"
	"github.com/kettek/apng"
	"io"
	"log"
	"os"
	"os/exec"
)

const fps = 60

func render(w io.Writer) error {
	log.Println("rendering")
	a := apng.APNG{}
	fs := scene.Frames()
	for i := 0; i < fs; i++ {
		ggctx := gg.NewContext(420, 240)
		scene.Update(float64(i) / float64(fs))
		scene.Render(ggctx)
		aframe := apng.Frame{
			Image:            ggctx.Image(),
			DelayNumerator:   1,
			DelayDenominator: fps,
			IsDefault:        i == 0,
		}
		a.Frames = append(a.Frames, aframe)
	}
	return apng.Encode(w, a)
}

func apngToH264(dstname string, srcname string) error {
	return apngToVideo(dstname, srcname,
		"-c:v", "libx264",
		"-crf", "24",
		"-preset", "veryslow",
		"-profile:v", "main",
	)
}

func apngToAV1(dstname string, srcname string) error {
	return apngToVideo(dstname, srcname,
		"-c:v", "libaom-av1",
		"-crf", "34",
		"-b:v", "0",
		"-strict", "experimental",
	)
}

func apngToVideo(dstname string, srcname string, options ...string) error {
	args := []string{
		"-f", "apng", "-r", "60", "-i", srcname,
		"-f", "mp4", "-r", "60",
		"-map_metadata", "-1",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		"-vf", "scale=trunc(iw/2)*2:trunc(ih/2)*2",
	}
	args = append(args, options...)
	args = append(args, []string{
		dstname,
		"-y",
	}...)
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
