package main

import (
	"fmt"
	"net/http"
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
	router.InitRoute(&web.Route{
		Path:     "/dictionary",
		Function: dictionaryController.CreateDictionary,
		Method:   http.MethodPost,
		Name:     "createDictionary",
	})
	router.InitRoute(wordController.GetAddWordRoute())
	router.InitRoute(&web.Route{
		Path:     "/word",
		Function: wordController.ListWords,
		Method:   http.MethodGet,
		Name:     "listWords",
	})

	router.RouterStart()
}
