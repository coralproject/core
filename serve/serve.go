package serve

import (
	"net/http"

	"github.com/coralproject/core/log"
)

const (
	port = "8080"
)

func RegisterEndpoint(path string, handler http.HandlerFunc) {

	log.Write("Registering Endpoint", path, handler)

	http.HandleFunc(path, handler)

}

func Start() {

	log.Write("ListenAndServe on :%n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
