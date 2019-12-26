package objects

import (
	"io"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/fogleman/ease"
)

type Queue struct {
	TopLeft, BottomRight graphics.Point
	Name                 string
}

func NewQueue(name string) *Queue {
	return &Queue{
		TopLeft:     graphics.At(10, 10),
		BottomRight: graphics.At(410, 34),
		Name:        name,
	}
}

func (obj *Queue) Render(w io.Writer) {
	render(w, `
<g>
	<rect x="{{.X}}" y="{{.Y}}" width="{{.Width}}" height="{{.Height}}" rx="8" ry="8" fill="#f5f2f0" stroke="#000000" />
	<text font-family="Iosevka" font-size="10px" x="{{.TextX}}" y="{{.TextY}}">{{.Text}}</text>
</g>
	`, struct {
		X, Y, Width, Height float64
		TextX, TextY        float64
		Text                string
	}{
		X:      obj.TopLeft.X,
		Y:      obj.TopLeft.Y,
		Width:  obj.BottomRight.X - obj.TopLeft.X,
		Height: obj.BottomRight.Y - obj.TopLeft.Y,
		TextX:  obj.TopLeft.X + 8,
		TextY:  obj.TopLeft.Y + 14,
		Text:   obj.Name,
	})
}

func (obj *Queue) LayoutMessages(frames int, msgs []*Message) animate.Animator {
	w := obj.BottomRight.X - obj.TopLeft.X - 20 - (10 * 2) - messageWidth
	dw := w / float64(len(msgs))

	var animators []animate.Animator
	for i, msg := range msgs {
		msg.SetVisibility(0)
		src := graphics.Point{
			X: obj.TopLeft.X + 40,
			Y: obj.TopLeft.Y + 4,
		}
		dst := graphics.Point{
			X: obj.BottomRight.X - 10 - messageWidth - float64(i)*dw,
			Y: obj.TopLeft.Y + 4,
		}
		msg.SetPosition(src)
		anim := animate.InParallel(
			animate.MoveTo(frames, msg, dst, ease.InOutQuint),
			animate.FadeIn(frames/2, msg),
		)
		animators = append(animators, anim)
	}
	return animate.InParallel(animators...)
}
