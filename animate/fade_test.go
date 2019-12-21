package animate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testVisibility struct {
	visibility float64
}

func (tv *testVisibility) GetVisibility() float64 {
	return tv.visibility
}

func (tv *testVisibility) SetVisibility(visibility float64) {
	tv.visibility = visibility
}

func TestFadeIn(t *testing.T) {
	v := &testVisibility{0}
	frames := 10

	a := FadeIn(frames, v)
	for i := 0; i < frames; i++ {
		a.Step(i)
		assert.Equal(t, float64(i+1)/float64(frames), v.visibility)
	}
	assert.Equal(t, 1.0, v.visibility)
}

func TestFadeOut(t *testing.T) {
	v := &testVisibility{1}
	frames := 10

	a := FadeOut(frames, v)
	for i := 0; i < frames; i++ {
		a.Step(i)
		assert.InDelta(t, float64(frames-(i+1))/float64(frames), v.visibility, 0.001)
	}
	assert.Equal(t, 0.0, v.visibility)
}
