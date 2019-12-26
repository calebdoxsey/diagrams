package objects

import (
	"fmt"
	"io"

	"github.com/calebdoxsey/diagrams/graphics"
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

func (msg *Message) Render(w io.Writer) {
	render(w, `
<g>
	<rect x="{{.X}}" y="{{.Y}}" width="{{.Width}}" height="{{.Height}}" rx="4" ry="4" fill="#FFF" stroke="#333"  opacity="{{.Opacity}}" />
	<text font-family="Iosevka" font-size="10px" x="{{.TextX}}" y="{{.TextY}}" text-anchor="middle" opacity="{{.Opacity}}">{{.Text}}</text>
</g>
	`, struct {
		X, Y, Width, Height float64
		TextX, TextY        float64
		Text                string
		Opacity             float64
	}{
		X:       msg.position.X,
		Y:       msg.position.Y,
		Width:   messageWidth,
		Height:  messageHeight,
		TextX:   msg.position.X + (messageWidth / 2),
		TextY:   msg.position.Y + (messageHeight) - 4,
		Text:    fmt.Sprintf("%03d", msg.number),
		Opacity: msg.visiblity,
	})
}
