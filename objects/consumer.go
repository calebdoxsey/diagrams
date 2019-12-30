package objects

import (
	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/geometry"
	"github.com/fogleman/ease"
	"io"
)

type Consumer struct {
	Name                 string
	TopLeft, BottomRight geometry.Point
}

func NewConsumer(name string, topLeft, bottomRight geometry.Point) *Consumer {
	return &Consumer{
		Name:        name,
		TopLeft:     topLeft,
		BottomRight: bottomRight,
	}
}

func (obj *Consumer) AnimatePreGetMessage(frames int, msg *Message) animate.Animator {
	return animate.InSequence(
		animate.MoveTo(frames, msg, obj.lineForGetArrow()[0].Translate(-messageWidth/2, -messageHeight/2), ease.InOutQuad),
	)
}

func (obj *Consumer) AnimateGetMessage(frames int, msg *Message) animate.Animator {
	return animate.InSequence(
		animate.MoveTo(frames, msg, obj.lineForGetArrow()[1].Translate(-messageWidth/2, -messageHeight/2), ease.InOutQuad),
	)
}

func (obj *Consumer) AnimateProcessMessage(frames int, msg *Message) animate.Animator {
	w := obj.BottomRight.X - obj.TopLeft.X
	h := obj.BottomRight.Y - obj.TopLeft.Y
	return animate.InSequence(
		animate.MoveTo(10, msg, geometry.At(obj.TopLeft.X+w/2, obj.TopLeft.Y+h/2).Translate(-messageWidth/2, -messageHeight/2), ease.InOutQuad),
		animate.NoOp(frames-20),
		animate.MoveTo(10, msg, obj.lineForCommitArrow()[0].Translate(-messageWidth/2, -messageHeight/2), ease.InOutQuad),
	)
}

func (obj *Consumer) AnimateCommitMessage(frames int, msg *Message) animate.Animator {
	return animate.InParallel(
		animate.MoveTo(frames, msg, obj.lineForCommitArrow()[1].Translate(-messageWidth/2, -messageHeight/2), ease.InOutQuad),
		animate.Delay(animate.FadeOut(frames/2, msg), frames/2),
	)
}

func (obj *Consumer) Render(out io.Writer) {
	w := obj.BottomRight.X - obj.TopLeft.X

	NewArrow(obj.lineForGetArrow()).Render(out)
	NewArrow(obj.lineForCommitArrow()).Render(out)

	render(out, `
<text font-family="Iosevka" font-size="12px" x="{{.GetTextX}}" y="{{.GetTextY}}" dominant-baseline="middle" text-anchor="end">get</text>
<text font-family="Iosevka" font-size="12px" x="{{.CommitTextX}}" y="{{.CommitTextY}}" dominant-baseline="middle" text-anchor="start">commit</text>
<rect x="{{.X}}" y="{{.Y}}" width="{{.Width}}" height="{{.Height}}" rx="8" ry="8" fill="#f5f2f0" stroke="#000000" />
<text font-family="Iosevka" font-size="12px" x="{{.TextX}}" y="{{.TextY}}" text-anchor="middle">{{.Text}}</text>
	`, struct {
		GetTextX, GetTextY       float64
		CommitTextX, CommitTextY float64
		X, Y, Width, Height      float64
		TextX, TextY             float64
		Text                     string
	}{
		GetTextX:    obj.TopLeft.X + w/2 - 8,
		GetTextY:    obj.TopLeft.Y - 14,
		CommitTextX: obj.TopLeft.X + w/2 + 8,
		CommitTextY: obj.TopLeft.Y - 14,
		X:           obj.TopLeft.X,
		Y:           obj.TopLeft.Y,
		Width:       obj.BottomRight.X - obj.TopLeft.X,
		Height:      obj.BottomRight.Y - obj.TopLeft.Y,
		TextX:       obj.TopLeft.X + w/2,
		TextY:       obj.TopLeft.Y + 14,
		Text:        obj.Name,
	})
}

func (obj *Consumer) lineForGetArrow() geometry.Line {
	w := obj.BottomRight.X - obj.TopLeft.X
	return geometry.Line{
		geometry.At(obj.TopLeft.X+w/2-4, obj.TopLeft.Y-30),
		geometry.At(obj.TopLeft.X+w/2-4, obj.TopLeft.Y-7),
	}
}

func (obj *Consumer) lineForCommitArrow() geometry.Line {
	w := obj.BottomRight.X - obj.TopLeft.X
	return geometry.Line{
		geometry.At(obj.TopLeft.X+w/2+4, obj.TopLeft.Y-4),
		geometry.At(obj.TopLeft.X+w/2+4, obj.TopLeft.Y-27),
	}
}
