package flow

import (
	"github.com/coralproject/core/log"
)

type flow struct {
	name  string
	steps []*step
}

type step struct {
	name         string
	startTrigger string
	endTrigger   string

	action func() // needs an interface definition
	meta   interface{}
}

type stepHandler struct {
	handler func()
}

func (s step) handle() error {

	log.Write("todo: broadcast", s.startTrigger)

	s.action()

	log.Write("todo: broadcast", s.endTrigger)

	return nil

}

var (
	flows []*flow
)

func (f flow) Run() {

	log.Write("Flow: Running ", f)

	for _, s := range f.steps {
		s.handle()
	}

}

func Create(n string) *flow {

	f := new(flow)
	f.name = n

	flows = append(flows, f)

	log.Write("Flow: create", n)

	return f

}

func (f *flow) AddStep(n string, st string, et string, a func()) *step {

	s := new(step)
	s.name = n
	s.startTrigger = st
	s.endTrigger = et
	s.action = a

	f.steps = append(f.steps, s)

	log.Write("Step: create", *s, "in", *f)

	return s

}
