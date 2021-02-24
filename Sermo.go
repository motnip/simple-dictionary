package main

import (
	"fmt"
	w "sermo/word"
)

func main() {

	newWord := w.CreateWord("label", "meaning")

	fmt.Println("This is the new word %v% wich means %v%", newWord.Label, newWord.Meaning)
}
