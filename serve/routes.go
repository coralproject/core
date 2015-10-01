package routes

import (
	"github.com/gorilla/mux"
)

func get(r *mux.Routes) {
	r.Handlefunc("/", controllers.Index).Methods("GET")
}
