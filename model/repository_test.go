package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

	result, err := repo.CreateDictionary("en")

	assert.NoError(t, err)
	assert.Equal(t, dictionary, result, "expected %v got %v", dictionary, result)
}

func TestCreateDictionary_dictionaryAlreadyExists_Fail(t *testing.T) {

	repo := NewRepository()

	language := "en"
	expectedErroMsg := fmt.Errorf("dictionary %s already exists", language)
	_, err := repo.CreateDictionary("en")
	assert.NoError(t, err)

	secondDictionaryResult, err := repo.CreateDictionary("en")
	assert.Error(t, err, "expected %v got %v", expectedErroMsg, secondDictionaryResult)
	assert.EqualError(t, expectedErroMsg, err.Error(), "expected %v got %v", expectedErroMsg, secondDictionaryResult)

}

func TestListDictionary(t *testing.T) {

	repo := NewRepository()

	_, _ = repo.CreateDictionary("en")

	result := repo.ListDictionary()

	assert.Equal(t, 1, len(result))
}

func TestAddWord(t *testing.T) {

	repo := NewRepository()
	repo.CreateDictionary("en")
	err := repo.AddWord(&newWord)

	assert.NoError(t, err)
}

func TestExistsDictionary_Succeed(t *testing.T) {

	repo := NewRepository()
	repo.CreateDictionary("en")

	result := repo.ExistsDictionary()

	assert.True(t, result)
}

func TestExistsDictionary_noDictionaryAvailable_Failed(t *testing.T) {

	repo := NewRepository()
	result := repo.ExistsDictionary()

	assert.False(t, result)
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
	assert.NoError(t, err)

	result, _ := repo.ListWords()

	assert.GreaterOrEqualf(t, 1, len(result), "no list of words have been returned", result)
	assert.Equal(t, newWord.Label, result[0].Label)
	assert.Equal(t, newWord.Meaning, result[0].Meaning)
	assert.Equal(t, newWord.Sentence, result[0].Sentence)
}
