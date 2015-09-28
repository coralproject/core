package flow

import (
	"net/http"

	"github.com/coralproject/core/log"
)

type Flow struct {
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

	s.action()
	Broadcast(s.endTrigger)

	return nil

}

type FlowHandler struct {
	F Flow
}

func (fh FlowHandler) Handle(w http.ResponseWriter, r *http.Request) {

	log.Write("Handling", r, "Running", fh.F)

	go fh.F.Run()

}

var (
	flows []*Flow
)

func (f Flow) Run() {

	log.Write("Flow: Running ", f)

	f.steps[0].handle()

}

// trigger a step
func Trigger(t string) {

	for _, f := range flows {
		for _, s := range f.steps {

			if s.startTrigger == t {
				s.handle()
			}

		}
	}

}

func Create(n string) *Flow {

	f := new(Flow)
	f.name = n

	flows = append(flows, f)

	log.Write("Flow: create", n)

	return f

}

func (f *Flow) AddStep(n string, st string, et string, a func()) *step {

	s := new(step)
	s.name = n
	s.startTrigger = st
	s.endTrigger = et
	s.action = a

	f.steps = append(f.steps, s)

	log.Write("Step: create", *s, "in", *f)

	return s

}
