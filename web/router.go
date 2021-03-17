package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	c "sermo/controller"
)

type Router struct {
	router     *mux.Router
	controller c.Controller
}

func NewRouter(controller c.Controller) *Router {
	return &Router{
		router:     mux.NewRouter().StrictSlash(true),
		controller: controller,
	}
}

func (r *Router) RouterStart() {
	r.router.HandleFunc("/dictionary", r.controller.CreateDictionary).Methods(http.MethodPost)
	r.router.HandleFunc("/word", r.controller.AddWord).Methods(http.MethodPost)
	r.router.HandleFunc("/word", r.controller.ListWords).GetMethods()
	fmt.Println("Server started... ")
	log.Fatal(http.ListenAndServe(":8080", r.router))
}
