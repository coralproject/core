package flow

import (
	"github.com/coralproject/core/log"
)

var (
	bus         = make(chan string, 1000)
	killMessage = "kill-that-channel-dead"
)

func Start() {
	listen()
}

func Stop() {
	bus <- killMessage
}

func listen() {

	for {
		t := <-bus

		if t == killMessage {
			//			break
		}

		log.Write("Flow: heard", t)

		Trigger(t)
	}

}

func Broadcast(m string) {

	log.Write("Flow: Broadcast", m)

	bus <- m
}
