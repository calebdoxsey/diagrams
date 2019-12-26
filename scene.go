package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/objects"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type Scene struct {
	objects  []objects.Object
	animator animate.Animator
	tmp      string
}

func NewScene(objects []objects.Object, animator animate.Animator) *Scene {
	s := &Scene{
		objects:  objects,
		animator: animator,
		tmp:      filepath.Join(os.TempDir(), uuid.New().String()),
	}
	if err := os.MkdirAll(s.tmp, 0755); err != nil {
		panic(err)
	}
	return s
}

func (s *Scene) Close() error {
	return os.RemoveAll(s.tmp)
}

func (s *Scene) Render(dst string) error {
	if err := s.renderSVGs(); err != nil {
		return err
	}

	if err := s.renderPNGs(); err != nil {
		return err
	}

	if err := s.renderMP4(dst); err != nil {
		return err
	}

	return nil
}

func (s *Scene) renderMP4(dst string) error {
	log.Println("converting pngs to", dst)
	cmd := exec.Command("ffmpeg",
		"-hide_banner", "-loglevel", "panic",
		"-r", "60",
		"-f", "image2",
		"-s", "420x240",
		"-i", filepath.Join(s.tmp, "pngs", "%05d.png"),
		"-vcodec", "libx264",
		"-crf", "17",
		"-pix_fmt", "yuv420p",
		"-profile:v", "main",
		"-tune", "animation",
		"-movflags", "+faststart",
		"-y",
		dst,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *Scene) renderPNGs() error {
	os.MkdirAll(filepath.Join(s.tmp, "pngs"), 0755)

	type payload struct {
		dst, src string
	}
	payloads := make(chan payload)

	var eg errgroup.Group
	eg.Go(func() error {
		defer close(payloads)
		for i := 0; i < s.animator.Frames(); i++ {
			payloads <- payload{
				dst: pngName(s.tmp, i),
				src: svgName(s.tmp, i),
			}
		}
		return nil
	})
	for i := 0; i < 8; i++ {
		eg.Go(func() error {
			for p := range payloads {
				dst, _ := filepath.Abs(p.dst)
				src, _ := filepath.Abs(p.src)
				log.Println("converting", p.src, "to", p.dst)
				// cmd := exec.Command("rsvg-convert",
				// 	"--background-color", "#FFFFFF",
				// 	"--output", p.dst,
				// 	p.src)
				cmd := exec.Command("cairosvg",
					"--output="+dst,
					"--width=420",
					"--height=240",
					src)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					return err
				}
			}
			return nil
		})
	}

	return eg.Wait()
}

func (s *Scene) renderSVGs() error {
	os.MkdirAll(filepath.Join(s.tmp, "svgs"), 0755)
	for frame := 0; frame < s.animator.Frames(); frame++ {
		s.animator.Step(frame)

		f, err := os.Create(svgName(s.tmp, frame))
		if err != nil {
			return err
		}
		f.WriteString(`<svg width="420" height="240" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
		f.WriteString(`<rect width="100%" height="100%" fill="white"/>`)
		f.WriteString(`
    <marker id='arrow-head' orient='auto' markerWidth='6' markerHeight='8'
            refX='0.1' refY='2'>
      <path d='M0,0 V4 L2,2 Z' fill='black' />
		</marker>
		`)

		for _, obj := range s.objects {
			obj.Render(f)
		}
		f.WriteString(`</svg>`)
		f.Close()
	}

	return nil
}

func svgName(dir string, frame int) string {
	return filepath.Join(dir, "svgs", fmt.Sprintf("%05d.svg", frame))
}

func pngName(dir string, frame int) string {
	return filepath.Join(dir, "pngs", fmt.Sprintf("%05d.png", frame))
}
