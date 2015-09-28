/*

Package flow allows flows of step to be composed together in a micro-service friendly manner.

Status: Experimental

A step is a func that is triggered by a notification and triggers a notification when complete.
A flow is a series of steps set up to accomplish a task, such as posting form input

Notifications may be broadcast via an internal or external dispatcher.

By sewing steps together via notifications, each step of each flow can naturally span multiple micro-services.  In addition, the registration of flows can be described in a clean declaritave way.

Example

	f := flow.Create("Post")

	f.AddStep("cfsr", "HandleFlow", "RequetVerified", cfsr)
	f.AddStep("authenticate", "RequetVerified", "UserAuthenticated", authenticate)
	f.AddStep("validation", "UserAuthenticated", "ContentValidated", validation)
	f.AddStep("store", "ContentValidated", "DataWritten", store)
	f.AddStep("updateSearch", "ContentValidated", "SearchUpdated", updateSearch)

	go flow.Start()

Issues:

* the action of each step at this point is func(), which is limiting.  A technique to either standardize a step action or to make it flexible would make this more useful.

* Differentiating between notifications going to a local bus vs. a shared microservice notification service could cut down on excess notification traffic.
*/

package flow

import (
	"net/http"

	"github.com/coralproject/core/log"
)

// top level Flow type
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

// each step is called through this handle func
func (s step) handle() error {

	// execute the action
	s.action()

	// broadcast the end trigger
	Broadcast(s.endTrigger)

	return nil

}

type FlowHandler struct {
	F Flow
}

// Handle a flow, assume http
func (fh FlowHandler) Handle(w http.ResponseWriter, r *http.Request) {

	log.Write("Handling", r, "Running", fh.F)

	go fh.F.Run()

}

var (
	flows []*Flow
)

// Kick off a flow
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

// Create a new flow
func Create(n string) *Flow {

	f := new(Flow)
	f.name = n

	flows = append(flows, f)

	log.Write("Flow: create", n)

	return f

}

// Add a step to a flow
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
