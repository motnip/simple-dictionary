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
	dictionaryController := controller.NewController(repository)
	wordController := controller.NewWordController(repository)

	router := web.NewRouter()
	router.InitRoute(dictionaryController.GetCreateDictionaryRoute())
	router.InitRoute(wordController.GetAddWordRoute())
	router.InitRoute(wordController.GetListWordRoute())

	router.RouterStart()
}
