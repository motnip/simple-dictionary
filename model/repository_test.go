package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

var newWord = Word{
	Label:   "hello",
	Meaning: "ciao",
}

type RepositoryTestSuite struct {
	suite.Suite
	repo Repository
}

func TestWordTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (r *RepositoryTestSuite) SetupTest() {
	r.repo = NewRepository()
}

func (r *RepositoryTestSuite) TestCreateDictionary() {

	dictionary := &Dictionary{
		Language: "en",
	}

	result, err := r.repo.CreateDictionary("en")

	assert.NoError(r.T(), err)
	assert.Equal(r.T(), dictionary, result, "expected %v got %v", dictionary, result)
}

func (r *RepositoryTestSuite) TestCreateDictionary_dictionaryAlreadyExists_Fail() {

	language := "en"
	expectedErroMsg := fmt.Errorf("dictionary %s already exists", language)
	_, err := r.repo.CreateDictionary("en")
	assert.NoError(r.T(), err)

	secondDictionaryResult, err := r.repo.CreateDictionary("en")
	assert.Error(r.T(), err, "expected %v got %v", expectedErroMsg, secondDictionaryResult)
	assert.EqualError(r.T(), expectedErroMsg, err.Error(), "expected %v got %v", expectedErroMsg, secondDictionaryResult)

}

func (r *RepositoryTestSuite) TestListDictionary() {

	_, _ = r.repo.CreateDictionary("en")

	result := r.repo.ListDictionary()

	assert.Equal(r.T(), 1, len(result))
}

func (r *RepositoryTestSuite) TestAddWord() {

	r.repo.CreateDictionary("en")
	err := r.repo.AddWord(&newWord)

	assert.NoError(r.T(), err)
}

func (r *RepositoryTestSuite) TestExistsDictionary_Succeed() {

	r.repo.CreateDictionary("en")

	result := r.repo.ExistsDictionary()

	assert.True(r.T(), result)
}

func (r *RepositoryTestSuite) TestExistsDictionary_noDictionaryAvailable_Failed() {

	result := r.repo.ExistsDictionary()

	assert.False(r.T(), result)
}

func (r *RepositoryTestSuite) TestListWord() {

	r.repo.CreateDictionary("en")

	sentence := "Hello world!"
	newWord := Word{
		Label:    "hello",
		Meaning:  "ciao",
		Sentence: sentence,
	}

	err := r.repo.AddWord(&newWord)
	assert.NoError(r.T(), err)

	result, _ := r.repo.ListWords()

	assert.GreaterOrEqualf(r.T(), 1, len(result), "no list of words have been returned", result)
	assert.Equal(r.T(), newWord.Label, result[0].Label)
	assert.Equal(r.T(), newWord.Meaning, result[0].Meaning)
	assert.Equal(r.T(), newWord.Sentence, result[0].Sentence)
}
