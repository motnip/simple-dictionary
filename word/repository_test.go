package word

import (
	"reflect"
	"testing"
)

func TestCreateDictionary(t *testing.T) {

	repo := NewRepository()

	dictionary := &Dictionary{
		Language: "en",
	}

	result := repo.CreateDictionary("en")

	if !reflect.DeepEqual(result, dictionary) {
		t.Errorf("expected %v got %v", dictionary, result)
	}
}

func TestAddWord(t *testing.T) {

	repo := NewRepository()

	newWord := Word{
		Label:   "hello",
		Meaning: "ciao",
	}

	result := repo.AddWord(&newWord)

	if len(result.Dictionary.Words) < 1 {
		t.Error("word has not been persisted", result)
	}

	if result.Dictionary.Words[0].Label != newWord.Label {
		t.Errorf("expected %v got %v", newWord.Label, result.Dictionary.Words[0].Label)
	}
}

func TestListWord(t *testing.T) {

	repo := NewRepository()

	sentence := "Hello world!"
	newWord := Word{
		Label:    "hello",
		Meaning:  "ciao",
		Sentence: sentence,
	}

	_ = repo.AddWord(&newWord)

	result := repo.ListWords()

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
