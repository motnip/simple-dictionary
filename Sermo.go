package main

import (
	"fmt"
	"sermo/controller"
	"sermo/web"
)

func main() {

	fmt.Println("Server starting ")

	router := web.NewRouter()

	goodBy := controller.NewGoodbye()
	router.AddPath(&web.Route{
		Path:           goodBy.Path(),
		ControllerFunc: goodBy.Goodbye,
	})

	greetings := controller.NewGreetings()
	router.AddPath(&web.Route{
		Path:           greetings.Path(),
		ControllerFunc: greetings.Greetings,
	})

	router.RouterStart()
}
