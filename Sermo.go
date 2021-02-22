package main

import (
	"fmt"
	w "sermo/word"
)

func main() {

	newWord := w.CreateWord("sto", "cazzo")

	var label string
	label = newWord.Label
	fmt.Println("This is the new word %v% wich means %v%", label, newWord.Meaning)
}
