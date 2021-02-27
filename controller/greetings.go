package controller

import (
	"fmt"
	"net/http"
)

func NewGreetings() *route {
	return &route{
		path: "/greetings",
	}
}

func (r *route) Greetings(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(httpResponse, "Welcome home!")
}

func (r *route) Path() string {
	return r.path
}
