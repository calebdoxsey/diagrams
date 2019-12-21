package objects

import (
	"image/color"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/fogleman/ease"
	"github.com/fogleman/gg"
)

// Kafka is a Kafka Queue
type Kafka struct {
	W, H     float64
	position graphics.Point
}

// NewKafka creates a new Kafka queue.
func NewKafka(w, h float64) *Kafka {
	k := &Kafka{
		W: w,
		H: h,
	}
	k.position = graphics.Point{X: 10, Y: 10}
	return k
}

// LayoutMessages creates an animator that will move messages to appear
// like a queue.
func (k *Kafka) LayoutMessages(frames int, msgs []*Message) animate.Animator {
	var animators []animate.Animator
	for i, msg := range msgs {
		src := graphics.Point{
			X: k.position.X + k.W - messageWidth,
			Y: 34,
		}
		dst := graphics.Point{
			X: k.position.X + 10 + float64(i)*(k.W-24)/messageWidth,
			Y: 34,
		}
		msg.SetPosition(src)
		animator := animate.MoveTo(frames, msg, dst, ease.OutQuart)
		animators = append(animators, animator)
	}
	return animate.InParallel(animators...)
}

// Render renders the Kafka queue.
func (k *Kafka) Render(ggctx *gg.Context) {
	x, y, w, h := k.position.X, k.position.Y, k.W, k.H

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(color.White))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(Color(0x666666)))
	ggctx.DrawRoundedRectangle(align(x), align(y), w, h, 4)
	ggctx.SetLineJoinRound()
	ggctx.FillPreserve()
	ggctx.Stroke()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(color.Black))
	ggctx.DrawStringAnchored("kafka", x+6, y+4, 0, 1)
	ggctx.Fill()
	ggctx.Pop()
}
