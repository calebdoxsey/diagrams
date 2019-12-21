package animate

type funcAnimator struct {
	frames int
	step   func(frame int)
}

// Func creates an animator from a function.
func Func(frames int, step func(frame int)) Animator {
	return funcAnimator{
		frames: frames,
		step:   step,
	}
}

func (a funcAnimator) Frames() int {
	return a.frames
}

func (a funcAnimator) Step(frame int) {
	a.step(frame)
}
