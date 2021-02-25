package controller

import (
	"fmt"
	"net/http"
)

func NewGoodbye() *route {
	return &route{
		path: "/goodbye",
	}
}

func (r *route) Goodbye(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(httpResponse, "Goodbye!")
}
