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
	controller := controller.NewController(repository)

	router := web.NewRouter()
	router.InitRoute(&web.Route{
		Path:     "/dictionary",
		Function: controller.CreateDictionary,
		Method:   http.MethodPost,
		Name:     "createDictionary",
	})
	router.InitRoute(&web.Route{
		Path:     "/word",
		Function: controller.AddWord,
		Method:   http.MethodPost,
		Name:     "addWord",
	})
	router.InitRoute(&web.Route{
		Path:     "/word",
		Function: controller.ListWords,
		Method:   http.MethodGet,
		Name:     "listWords",
	})

	router.RouterStart()
}
