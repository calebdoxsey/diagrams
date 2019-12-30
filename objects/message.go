package objects

import (
	"fmt"
	"io"

	"github.com/calebdoxsey/diagrams/geometry"
)

const (
	messageWidth  = 26.0
	messageHeight = 17.0
)

type Message struct {
	Number    int
	visiblity float64
	position  geometry.Point
	offset    geometry.Point
	pctLoaded float64
}

func NewMessage(number int) *Message {
	msg := &Message{
		visiblity: 1,
		Number:    number,
	}
	return msg
}

func (msg *Message) GetPosition() geometry.Point {
	return msg.position
}

func (msg *Message) SetPosition(point geometry.Point) {
	msg.position = point
}

func (msg *Message) GetVisibility() float64 {
	return msg.visiblity
}

func (msg *Message) SetVisibility(visibility float64) {
	msg.visiblity = visibility
}

func (msg *Message) SetOffset(point geometry.Point) {
	msg.offset = point
}

func (msg *Message) Render(w io.Writer) {
	render(w, `
<g>
	<rect x="{{.X}}" y="{{.Y}}" width="{{.Width}}" height="{{.Height}}" rx="4" ry="4" fill="#FFF" stroke="#333"  opacity="{{.Opacity}}" />
	<text font-family="Iosevka" font-size="12px" x="{{.TextX}}" y="{{.TextY}}" text-anchor="middle" opacity="{{.Opacity}}">{{.Text}}</text>
</g>
	`, struct {
		X, Y, Width, Height float64
		TextX, TextY        float64
		Text                string
		Opacity             float64
	}{
		X:       msg.position.X + msg.offset.X,
		Y:       msg.position.Y + msg.offset.Y,
		Width:   messageWidth,
		Height:  messageHeight,
		TextX:   msg.position.X + msg.offset.X + (messageWidth / 2),
		TextY:   msg.position.Y + msg.offset.Y + (messageHeight) - 4,
		Text:    fmt.Sprintf("%03d", msg.Number),
		Opacity: msg.visiblity,
	})
}
