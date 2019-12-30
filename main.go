package main

import (
	"log"
	"os"

	"github.com/calebdoxsey/diagrams/animate"
	"github.com/calebdoxsey/diagrams/geometry"
	"github.com/calebdoxsey/diagrams/objects"
)

func main() {
	log.SetFlags(0)

	_ = os.MkdirAll("out", 0755)

	//renderBasicMessageQueue()
	//renderPipelineMessageQueue()
	//renderBatchMessageQueue()
	renderPartitionMessageQueue()
}

func renderBasicMessageQueue() {
	var objs []objects.Object
	q := objects.NewQueue("queue", 10)
	objs = append(objs, q)

	c1 := objects.NewConsumer("consumer", geometry.At(210-40, 70), geometry.At(210+40, 130))
	objs = append(objs, c1)

	var msgs []*objects.Message
	for i := 0; i < 24; i++ {
		msg := objects.NewMessage(24 - i)
		msgs = append(msgs, msg)
		objs = append(objs, msg)
	}

	anim := q.LayoutMessages(20, msgs)
	anim = animate.Delay(anim, 20)

	for i := 0; i < 8; i++ {
		i := i
		msg := objects.NewMessage(i + 1)
		objs = append(objs, msg)
		msg.SetVisibility(0)

		anim = animate.InSequence(anim,
			animate.Func(1, func(frame int) {
				msg.SetPosition(msgs[len(msgs)-(i+1)].GetPosition())
				msg.SetVisibility(1)
				msgs[len(msgs)-(i+1)].SetVisibility(0.25)
			}),
			c1.AnimatePreGetMessage(20, msg),
			c1.AnimateGetMessage(20, msg),
			c1.AnimateProcessMessage(60, msg),
			c1.AnimateCommitMessage(20, msg),
		)
	}

	s := NewScene(420, 160, objs, anim)
	if err := s.Render("./out/basic-message-queue.mp4"); err != nil {
		log.Fatalln(err)
	}
}

func renderPipelineMessageQueue() {
	var objs []objects.Object
	q := objects.NewQueue("queue", 10)
	objs = append(objs, q)

	c1 := objects.NewConsumer("consumer", geometry.At(210-40, 70), geometry.At(210+40, 130))
	objs = append(objs, c1)

	var msgs []*objects.Message
	for i := 0; i < 24; i++ {
		msg := objects.NewMessage(24 - i)
		msgs = append(msgs, msg)
		objs = append(objs, msg)
	}

	anim := q.LayoutMessages(20, msgs)
	anim = animate.Delay(anim, 20)

	var manim animate.Animator

	for i := 0; i < 8; i++ {
		i := i
		msg := objects.NewMessage(i + 1)
		objs = append(objs, msg)
		msg.SetVisibility(0)

		initialize := animate.Func(1, func(frame int) {
			msg.SetPosition(msgs[len(msgs)-(i+1)].GetPosition())
			msg.SetVisibility(1)
			msgs[len(msgs)-(i+1)].SetVisibility(0.25)
		})

		if i == 0 {
			manim = animate.InSequence(
				initialize,
				c1.AnimatePreGetMessage(20, msg),
				c1.AnimateGetMessage(20, msg),
				c1.AnimateProcessMessage(60, msg),
				c1.AnimateCommitMessage(20, msg),
			)
		} else {
			delay := 20
			if i > 1 {
				delay += (i - 1) * 60
			}
			manim = animate.InParallel(manim,
				animate.InSequence(
					animate.NoOp(delay),
					initialize,
					c1.AnimatePreGetMessage(20, msg),
					c1.AnimateGetMessage(20, msg),
					animate.NoOp(40),
					c1.AnimateProcessMessage(60, msg),
					c1.AnimateCommitMessage(20, msg),
				),
			)
		}
	}

	anim = animate.InSequence(anim, manim)

	s := NewScene(420, 160, objs, anim)
	if err := s.Render("./out/pipeline-message-queue.mp4"); err != nil {
		log.Fatalln(err)
	}
}

func renderBatchMessageQueue() {
	var objs []objects.Object
	q := objects.NewQueue("queue", 10)
	objs = append(objs, q)

	c1 := objects.NewConsumer("consumer", geometry.At(210-40, 70), geometry.At(210+40, 130))
	objs = append(objs, c1)

	var msgs []*objects.Message
	for i := 0; i < 24; i++ {
		msg := objects.NewMessage(24 - i)
		msgs = append(msgs, msg)
		objs = append(objs, msg)
	}

	anim := q.LayoutMessages(20, msgs)
	anim = animate.Delay(anim, 20)

	for i := 0; i < 12; i += 3 {
		var manim animate.Animator
		for j := i; j < i+3; j++ {
			j := j
			msg := objects.NewMessage(j + 1)
			objs = append(objs, msg)
			msg.SetVisibility(0)
			msg.SetOffset(geometry.At(float64(j-i-1)*4, float64(j-i-1)*4))

			manim = animate.InParallel(manim,
				animate.InSequence(
					animate.Func(1, func(frame int) {
						msg.SetPosition(msgs[len(msgs)-(j+1)].GetPosition())
						msg.SetVisibility(1)
						msgs[len(msgs)-(j+1)].SetVisibility(0.25)
					}),
					c1.AnimatePreGetMessage(20, msg),
					c1.AnimateGetMessage(20, msg),
					c1.AnimateProcessMessage(120, msg),
					c1.AnimateCommitMessage(20, msg),
				),
			)
		}
		anim = animate.InSequence(anim, manim)
	}

	s := NewScene(420, 160, objs, anim)
	if err := s.Render("./out/batch-message-queue.mp4"); err != nil {
		log.Fatalln(err)
	}
}

func renderPartitionMessageQueue() {
	var objs []objects.Object
	q1 := objects.NewQueue("queue 1", 10)
	objs = append(objs, q1)
	q2 := objects.NewQueue("queue 2", 44)
	objs = append(objs, q2)
	q3 := objects.NewQueue("queue 3", 78)
	objs = append(objs, q3)

	c1 := objects.NewConsumer("consumer 1", geometry.At(80-40, 140), geometry.At(80+40, 200))
	objs = append(objs, c1)
	c2 := objects.NewConsumer("consumer 2", geometry.At(210-40, 140), geometry.At(210+40, 200))
	objs = append(objs, c2)
	c3 := objects.NewConsumer("consumer 3", geometry.At(340-40, 140), geometry.At(340+40, 200))
	objs = append(objs, c3)

	cs := []*objects.Consumer{c1, c2, c3}

	var msgs1, msgs2, msgs3 []*objects.Message
	for i := 0; i < 72; i++ {
		msg := objects.NewMessage(72 - i)
		objs = append(objs, msg)
		switch i % 3 {
		case 0:
			msgs3 = append(msgs3, msg)
		case 1:
			msgs2 = append(msgs2, msg)
		case 2:
			msgs1 = append(msgs1, msg)
		}
	}

	anim := animate.InParallel(
		q1.LayoutMessages(20, msgs1),
		q2.LayoutMessages(20, msgs2),
		q3.LayoutMessages(20, msgs3),
	)
	anim = animate.Delay(anim, 20)

	for i := 0; i < 8; i++ {
		i := i
		var manim animate.Animator
		for j, msgs := range [][]*objects.Message{msgs1, msgs2, msgs3} {
			msgs := msgs
			c := cs[j]
			msg := objects.NewMessage(msgs[len(msgs)-(i+1)].Number)
			objs = append(objs, msg)
			msg.SetVisibility(0)

			manim = animate.InParallel(manim,
				animate.InSequence(
					animate.Func(1, func(frame int) {
						msg.SetPosition(msgs[len(msgs)-(i+1)].GetPosition())
						msg.SetVisibility(1)
						msgs[len(msgs)-(i+1)].SetVisibility(0.25)
					}),
					c.AnimatePreGetMessage(20, msg),
					c.AnimateGetMessage(20, msg),
					c.AnimateProcessMessage(60, msg),
					c.AnimateCommitMessage(20, msg),
				),
			)
		}
		anim = animate.InSequence(anim, manim)
	}

	s := NewScene(420, 240, objs, anim)
	if err := s.Render("./out/partition-message-queue.mp4"); err != nil {
		log.Fatalln(err)
	}
}
