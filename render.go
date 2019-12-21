package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/calebdoxsey/diagrams/objects"
	"github.com/fogleman/gg"
	"github.com/kettek/apng"
)

const fps = 60

func render(w io.Writer) error {
	log.Println("rendering")
	a := apng.APNG{}

	kafka := objects.NewKafka(imageWidth-20, 50)

	var msgs []*objects.Message
	for i := 0; i < 20; i++ {
		msgs = append(msgs, objects.NewMessage(i+1))
	}

	consumer := objects.NewConsumer(graphics.At(imageWidth/2-30, 110), 60, 40)

	animator := kafka.LayoutMessages(20, msgs)
	animator = animate.InSequence(animator,
		animate.NoOp(10),
		consumer.ProcessMessage(msgs[8]),
	)
	fs := animator.Frames()
	for i := 0; i < fs; i++ {
		ggctx := gg.NewContext(int(imageWidth), int(imageHeight))
		objects.NewBackground(imageWidth, imageHeight).Render(ggctx)
		kafka.Render(ggctx)
		consumer.Render(ggctx)
		for j := len(msgs) - 1; j >= 0; j-- {
			msgs[j].Render(ggctx)
		}
		animator.Step(i)
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
