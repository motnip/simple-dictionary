package web

import (
	"sermo/controller"
	"testing"
)

func TestAddPath(t *testing.T) {

	router := NewRouter()

	greetings := controller.NewGreetings()
	router.AddPath(greetings)

	if len(router.paths) < 1 {
		t.Error("no paths have not been persisted")
	}
}
