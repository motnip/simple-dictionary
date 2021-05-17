package main

import (
	"fmt"
	"github.com/motnip/sermo/service"

	"github.com/motnip/sermo/controller"
	"github.com/motnip/sermo/model"
	"github.com/motnip/sermo/web"
)

var BuildVersion = "development"

func main() {

	fmt.Println("Sermo version " + BuildVersion)

	fmt.Println("Server starting... ")

	repository := model.NewRepository()
	wordService := service.NewWordService(repository)
	dictionaryController := controller.NewController(repository)
	wordController := controller.NewWordController(wordService)

	router := web.NewRouter()
	router.InitRoute(dictionaryController.GetCreateDictionaryRoute())
	router.InitRoute(dictionaryController.GetListAllDictionary())
	router.InitRoute(wordController.GetAddWordRoute())
	router.InitRoute(wordController.GetListWordRoute())

	router.RouterStart()
}
