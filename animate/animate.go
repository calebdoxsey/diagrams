package animate

// An Animator is used to implementation animation. An animator will be run
// for the given number of frames by calling the step function for each frame.
type Animator interface {
	Frames() int
	Step(frame int)
}

type multiAnimator struct {
	animators []Animator
	ranges    [][2]int
	frames    int
}

func (m *multiAnimator) Frames() int {
	return m.frames
}

func (m *multiAnimator) Step(frame int) {
	for i, a := range m.animators {
		s, e := m.ranges[i][0], m.ranges[i][1]
		if s <= frame && frame < e {
			a.Step(frame - s)
		}
	}
}

// InSequence returns a new Animator that composes sub-animators in sequence.
func InSequence(animators ...Animator) Animator {
	m := &multiAnimator{}
	for _, a := range animators {
		if a == nil {
			continue
		}
		fs := a.Frames()
		m.animators = append(m.animators, a)
		m.ranges = append(m.ranges, [2]int{m.frames, m.frames + fs})
		m.frames += fs
	}
	return m
}

// InParallel returns a new Animator that composes sub-animators in parallel.
func InParallel(animators ...Animator) Animator {
	m := &multiAnimator{}
	for _, a := range animators {
		if a == nil {
			continue
		}
		fs := a.Frames()
		m.animators = append(m.animators, a)
		m.ranges = append(m.ranges, [2]int{0, fs})
		if fs > m.frames {
			m.frames = fs
		}
	}
	return m
}

type noop struct {
	frames int
}

// NoOp returns an Animator that runs the for the given number of frames, but
// doesn't do anything for 'Step'.
func NoOp(frames int) Animator {
	return noop{frames}
}

func (n noop) Frames() int {
	return n.frames
}

func (_ noop) Step(frame int) {}

func Delay(animator Animator, frames int) Animator {
	return InSequence(NoOp(frames), animator)
}
