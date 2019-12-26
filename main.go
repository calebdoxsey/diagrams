package main

import (
	"log"
	"os"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/graphics"
	"github.com/calebdoxsey/diagrams/objects"
	"github.com/fogleman/ease"
)

func main() {
	log.SetFlags(0)

	_ = os.MkdirAll("out", 0755)

	var objs []objects.Object
	q := objects.NewQueue("queue")
	objs = append(objs, q)

	var msgs []*objects.Message
	for i := 0; i < 24; i++ {
		msg := objects.NewMessage(24 - i)
		msgs = append(msgs, msg)
		objs = append(objs, msg)
	}

	anim := q.LayoutMessages(30, msgs)
	anim = animate.Delay(anim, 30)

	m1 := objects.NewMessage(1)
	objs = append(objs, m1)
	m1.SetVisibility(0)

	a1 := objects.NewArrow(graphics.Line{graphics.At(100, 40), graphics.At(100, 60)})
	objs = append(objs, a1)

	anim = animate.InSequence(anim,
		animate.Func(1, func(frame int) {
			m1.SetPosition(msgs[len(msgs)-1].GetPosition())
			m1.SetVisibility(1)
			msgs[len(msgs)-1].SetVisibility(0.5)
		}),
		animate.MoveTo(10, m1, graphics.At(100, 100), ease.Linear),
	)

	anim = animate.InSequence(anim, animate.NoOp(30))

	s := NewScene(objs, anim)
	if err := s.Render("./out/example.mp4"); err != nil {
		log.Fatalln(err)
	}
}
