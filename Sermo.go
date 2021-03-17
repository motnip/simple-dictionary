package main

import (
	"fmt"
	"sermo/controller"
	"sermo/model"
	"sermo/web"
)

func main() {

	fmt.Println("Server starting... ")

	repository := model.NewRepository()
	controller := controller.NewController(repository)

	router := web.NewRouter(controller)
	router.RouterStart()
}
