package main

import (
	"fmt"
	"sermo/web"
)

func main() {

	fmt.Println("Server starting ")

	router := web.NewRouter()
	router.RouterStart()
}
