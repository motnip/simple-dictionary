package main

import (
	"fmt"
	w "sermo/word"
)

func main() {

	newWord := w.CreateWord("sto", "cazzo")

	fmt.Println("This is the new word %v% wich means %v%", newWord.Label, newWord.Meaning)
}
