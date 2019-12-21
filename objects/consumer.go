package objects

import (
	"image/color"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/fogleman/ease"
	"github.com/fogleman/gg"
)

type Consumer struct {
	Position   graphics.Point
	W, H       float64
	Send, Recv *Arrow

	arrowHorizontalPadding float64
	arrowVerticalPadding   float64
	arrowHeight            float64
}

// NewConsumer creates a new Consumer.
func NewConsumer(topLeft graphics.Point, w, h float64) *Consumer {
	obj := &Consumer{
		Position: topLeft,
		W:        w,
		H:        h,

		arrowHorizontalPadding: 5.0,
		arrowVerticalPadding:   6.0,
		arrowHeight:            38.0,
	}

	x, y, w := obj.Position.X, obj.Position.Y, obj.W
	ahp, avp := obj.arrowHorizontalPadding, obj.arrowVerticalPadding
	ah := obj.arrowHeight
	obj.Recv = NewArrow(graphics.Line{
		graphics.At(x+w/2-ahp, y-ah-avp),
		graphics.At(x+w/2-ahp, y-avp),
	})
	obj.Send = NewArrow(graphics.Line{
		graphics.At(x+w/2+ahp, y-avp),
		graphics.At(x+w/2+ahp, y-ah-avp),
	})

	return obj
}

// Render renders the Consumer.
func (obj *Consumer) Render(ggctx *gg.Context) {
	x, y, w, h := obj.Position.X, obj.Position.Y, obj.W, obj.H
	ahp, avp := obj.arrowHorizontalPadding, obj.arrowVerticalPadding
	ah := obj.arrowHeight

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(color.White))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(Color(0x666666)))
	ggctx.DrawRoundedRectangle(align(x), align(y), w, h, 4)
	ggctx.SetLineJoinRound()
	ggctx.FillPreserve()
	ggctx.Stroke()
	ggctx.Pop()

	obj.Recv.Render(ggctx)
	obj.Send.Render(ggctx)

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(color.Black))
	ggctx.DrawStringAnchored("consumer", x+w/2, y+4, 0.5, 1)
	ggctx.DrawStringAnchored("get", x+w/2-ahp-4, y-ah/2-avp, 1, 0.5)
	ggctx.DrawStringAnchored("commit", x+w/2+ahp+4, y-ah/2-avp, 0, 0.5)
	ggctx.Fill()
	ggctx.Pop()
}

func (obj *Consumer) ProcessMessage(msg *Message) animate.Animator {
	mid := obj.Position.Translate(obj.W/2, obj.H/2)

	return animate.InSequence(
		animate.MoveTo(15, msg, obj.Recv.Line[0].Translate(-messageWidth/2, -messageHeight/2), ease.Linear),
		animate.MoveTo(15, msg, obj.Recv.Line[1].Translate(-messageWidth/2, -messageHeight/2), ease.Linear),
		animate.MoveTo(15, msg, mid.Translate(-messageWidth/2, 0), ease.Linear),
		animate.NoOp(60),
		animate.MoveTo(15, msg, obj.Send.Line[0].Translate(-messageWidth/2, -messageHeight/2), ease.Linear),
		animate.InParallel(
			//animate.MoveTo(15, msg, obj.Send.Line[1].Translate(-messageWidth/2, -messageHeight/2), ease.Linear),
			animate.FadeOut(60, msg),
		),
	)
}
