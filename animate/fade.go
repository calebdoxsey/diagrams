package animate

import "github.com/calebdoxsey/diagrams/graphics"

type fader struct {
	obj        graphics.HasVisibility
	frames     int
	start, end float64
}

// FadeTo creates an animator that fades an object to the requested visibility.
func FadeTo(frames int, obj graphics.HasVisibility, visibility float64) Animator {
	return &fader{
		obj:    obj,
		frames: frames,
		end:    visibility,
	}
}

func (a *fader) Frames() int {
	return a.frames
}

func (a *fader) Step(frame int) {
	if frame == 0 {
		a.start = a.obj.GetVisibility()
	}
	pct := float64(frame+1) / float64(a.frames)
	a.obj.SetVisibility(a.start + (a.end-a.start)*pct)
}

// FadeIn creates an animator that fades in an object.
func FadeIn(frames int, obj graphics.HasVisibility) Animator {
	return FadeTo(frames, obj, 1)
}

// FadeOut creates an animator that fades out an object.
func FadeOut(frames int, obj graphics.HasVisibility) Animator {
	return FadeTo(frames, obj, 0)
}
