/*

	Example of registering a flow with a basic http handler

*/
package flow

import (
	"net/http"

	"github.com/coralproject/core/flow"
	"github.com/coralproject/core/log"
)

const (
	port = "8080"
)

func RegisterHandler(path string, handler http.HandlerFunc) {

	log.Write("Registering Handler", path, handler)

	http.HandleFunc(path, handler)

}

func RegisterFlow(p string, f *flow.Flow) {

	log.Write("Registering Flow", p, *f)

	fh := new(flow.FlowHandler)
	fh.F = *f

	http.HandleFunc(p, fh.Handle)

}

func Start() {

	log.Write("ListenAndServe on :%n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
