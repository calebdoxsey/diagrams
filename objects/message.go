package objects

import (
	"fmt"

	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/fogleman/gg"
)

const (
	messageWidth  = 23.0
	messageHeight = 15.0
)

type Message struct {
	visiblity float64
	position  graphics.Point
	number    int
	pctLoaded float64
}

func NewMessage(number int) *Message {
	msg := &Message{
		visiblity: 1,
		number:    number,
	}
	return msg
}

func (msg *Message) GetPosition() graphics.Point {
	return msg.position
}

func (msg *Message) SetPosition(point graphics.Point) {
	msg.position = point
}

func (msg *Message) GetVisibility() float64 {
	return msg.visiblity
}

func (msg *Message) SetVisibility(visibility float64) {
	msg.visiblity = visibility
}

func (msg *Message) Render(ggctx *gg.Context) {
	x, y, w, h := msg.position.X, msg.position.Y, messageWidth, messageHeight

	clear := Color(0x000000)
	clear.A = 0
	bg := Color(0xFFFFFF)
	border := Color(0x666666)
	fg := Color(0x000000)

	if msg.visiblity != 1 {
		bg.A = GetAlpha(msg.visiblity)
		border.A = GetAlpha(msg.visiblity)
		fg.A = GetAlpha(msg.visiblity)
	}

	fmt.Println(bg)

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(bg))
	ggctx.DrawRoundedRectangle(align(x), align(y), w, h, 4)
	ggctx.Fill()
	ggctx.Pop()

	// ggctx.Push()
	// ggctx.SetFillStyle(gg.NewSolidPattern(color.Black))
	// ggctx.DrawRoundedRectangle(align(x), align(y), w*msg.pctLoaded, h, 2)
	// ggctx.Fill()
	// ggctx.Pop()

	ggctx.Push()
	ggctx.SetFillStyle(gg.NewSolidPattern(clear))
	ggctx.SetStrokeStyle(gg.NewSolidPattern(border))
	ggctx.DrawRoundedRectangle(align(x), align(y), w, h, 4)
	ggctx.Stroke()
	ggctx.Pop()

	ggctx.Push()
	ggctx.SetFontFace(fontFace)
	ggctx.SetFillStyle(gg.NewSolidPattern(fg))
	ggctx.DrawStringAnchored(fmt.Sprintf("%03d", msg.number), x+w/2, y+h/2, 0.5, 0.3)
	ggctx.Fill()
	ggctx.Pop()
}
