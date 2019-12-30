package animate

import "github.com/calebdoxsey/diagrams/geometry"
import "github.com/fogleman/ease"

type moveAnimator struct {
	frames   int
	obj      geometry.Positionable
	src, dst geometry.Point
	easing   ease.Function
}

// MoveTo creates an animator that will move an object to the specified
// position.
func MoveTo(frames int, obj geometry.Positionable, dst geometry.Point, easing ease.Function) Animator {
	if easing == nil {
		easing = ease.Linear
	}
	return &moveAnimator{
		frames: frames,
		obj:    obj,
		dst:    dst,
		easing: easing,
	}
}

func (a *moveAnimator) Frames() int {
	return a.frames
}

func (a *moveAnimator) Step(frame int) {
	if frame == 0 {
		a.src = a.obj.GetPosition()
	}
	pct := float64(frame+1) / float64(a.frames)
	adjusted := a.easing(pct)
	dx, dy := a.dst.X-a.src.X, a.dst.Y-a.src.Y
	a.obj.SetPosition(geometry.Point{
		X: a.src.X + dx*adjusted,
		Y: a.src.Y + dy*adjusted,
	})
}
