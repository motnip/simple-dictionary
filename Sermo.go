package main

import (
	"fmt"
	controller "sermo/controller"
	web "sermo/web"
)

func main() {

	/*newWord := w.CreateWord("label", "meaning", "")
	fmt.Printf("This is the new word %v wich means %v", newWord.Label, newWord.Meaning)
	repo := w.NewRepository()
	_ = repo.AddWord(newWord)
	result := repo.ListWords()
	fmt.Printf("This is the new word %v wich means %v ", result[0].Label, result[0].Meaning)*/

	fmt.Println("Server starting ")

	router := web.NewRouter()

	greetings := controller.NewGreetings()
	router.AddPath(greetings)
	router.RouterStart()

}
