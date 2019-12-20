package draw

import "github.com/fogleman/gg"

// An Animation is an Object for the given number of frames.
type Animation interface {
	Object
	Frames() int
	Update(float64)
}

type staticAnimation struct {
	Object
	frames int
}

// NewStatic creates a new static animation for the given number of frames.
func NewStatic(obj Object, frames int) Animation {
	return staticAnimation{obj, frames}
}

// Frames returns the number of frames for the animation.
func (s staticAnimation) Frames() int {
	return s.frames
}

func (s staticAnimation) Update(float64) {}

type basic struct {
	object Object
	frames int
	updater func(float64)
}

func NewBasicAnimation(obj Object, frames int, updater func(float64)) Animation {
	return basic{
		object: obj,
		frames: frames,
		updater: updater,
	}
}

func (b basic) Render(ggctx *gg.Context) {
	b.object.Render(ggctx)
}

func (b basic) Frames() int {
	return b.frames
}

func (b basic) Update(pct float64) {
	b.updater(pct)
}

// A Curve changes a percentage completed for a tween.
type Curve = func(completed float64) float64

type curved struct {
	Animation
	curve Curve
}

func NewCurved(animation Animation, curve Curve) Animation {
	return curved{
		Animation: animation,
		curve:     curve,
	}
}

func (c curved) Update(pct float64) {
	c.Animation.Update(c.curve(pct))
}

type sequence struct {
	steps  []Animation
	frames int
}

// NewSequence creates a new sequence of animations that are run one after the other.
func NewSequence(steps ...Animation) Animation {
	s := sequence{
		steps: steps,
	}
	for _, step := range steps {
		s.frames += step.Frames()
	}
	return s
}

func (s sequence) Frames() int {
	return s.frames
}

func (s sequence) Render(ggctx *gg.Context) {
	for _, step := range s.steps {
		step.Render(ggctx)
	}
}

func (s sequence) Update(pct float64) {
	frame := int(float64(s.frames)*pct)
	prev := 0
	for _, step := range s.steps {
		next := prev + step.Frames()
		switch {
		case frame > next:
			step.Update(1)
		case frame < prev:
			step.Update(0)
		default:
			step.Update(float64(frame-prev)/float64(next-prev))
		}
		prev = next
	}
}
