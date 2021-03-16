package model

import (
	"reflect"
	"testing"
)

var newWord = Word{
	Label:   "hello",
	Meaning: "ciao",
}

func TestCreateDictionary(t *testing.T) {

	repo := NewRepository()

	dictionary := &Dictionary{
		Language: "en",
	}

	result, _ := repo.CreateDictionary("en")

	if !reflect.DeepEqual(result, dictionary) {
		t.Errorf("expected %v got %v", dictionary, result)
	}
}

func TestCreateDictionary_dictionaryAlreadyExists_Fail(t *testing.T) {

	repo := NewRepository()

	_, err := repo.CreateDictionary("en")
	if err != nil {
		t.Errorf("Error while create first dictionary for test")
	}

	secondDictionaryResult, err := repo.CreateDictionary("en")

	if err == nil {
		t.Errorf("expected %v got %v", "dictionary already exists", secondDictionaryResult)
	}

	if err.Error() != "dictionary already exists" {
		t.Errorf("expected %v got %v", "dictionary already exists", secondDictionaryResult)
	}
}

func TestExistsDictionary_dictionaryAlreadyExists_Fail(t *testing.T) {

	repo := NewRepository()

	_, err := repo.CreateDictionary("en")
	if err != nil {
		t.Errorf("Error while create first dictionary for test")
	}

	existsDictionary := repo.existsDictionaryOfLanguage("en")

	if !existsDictionary {
		t.Errorf("expected %v got %v", true, false)
	}
}

func TestListDictionary(t *testing.T) {

	repo := NewRepository()

	_, _ = repo.CreateDictionary("en")

	result := repo.ListDictionary()

	if len(result) != 1 {
		t.Errorf("expected %v got %v", 1, len(result))
	}
}

func TestAddWord(t *testing.T) {

	repo := NewRepository()
	repo.CreateDictionary("en")
	err := repo.AddWord(&newWord)

	if err != nil {
		t.Errorf("expected %v go %v", nil, err)

	}
}

func TestAddWord_noDictionaryExists_Failed(t *testing.T) {

	repo := NewRepository()
	err := repo.AddWord(&newWord)

	if err == nil {
		t.Errorf("expected %v go %v", "no dictionary available", err)

	}
}

func TestListWord(t *testing.T) {

	repo := NewRepository()
	repo.CreateDictionary("en")

	sentence := "Hello world!"
	newWord := Word{
		Label:    "hello",
		Meaning:  "ciao",
		Sentence: sentence,
	}

	err := repo.AddWord(&newWord)
	if err != nil {
		t.Errorf("Error not expected %v", err)
	}

	result,_ := repo.ListWords()

	if len(result) < 1 {
		t.Error("no list of words have been returned", result)
	}

	if result[0].Label != newWord.Label {
		t.Errorf("expected %s, got %s ", newWord.Label, result[0].Label)
	}
	if result[0].Meaning != newWord.Meaning {
		t.Errorf("expected %s, got %s ", newWord.Meaning, result[0].Meaning)
	}
	if result[0].Sentence != newWord.Sentence {
		t.Errorf("expected %s, got %s ", newWord.Sentence, result[0].Sentence)
	}
}

func TestListWord_noDictionaryExists_Failed(t *testing.T) {

	repo := NewRepository()


	result,err := repo.ListWords()

	if err == nil {
		t.Errorf("expected %v got %v", "no dictionary available", err)

	}

	if result != nil {
		t.Errorf("expected %v got %v", nil, result)

	}
}
