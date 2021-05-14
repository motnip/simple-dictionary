package model

import (
	"errors"
	"fmt"

	"github.com/motnip/sermo/system"
)

const NO_DICTIONARY = "no dictionary available"

type Repository interface {
	CreateDictionary(language string) (*Dictionary, error)
	DeleteDictionary()
	AddWord(word *Word) error
	ListWords() ([]*Word, error)
	ListDictionary() []*Dictionary
}

type repository struct {
	Dictionary *Dictionary
	logger     *system.SermoLog
}

func NewRepository() Repository {

	return &repository{
		logger: system.NewLog(),
	}
}

func (r *repository) CreateDictionary(language string) (*Dictionary, error) {

	if r.existsDictionaryOfLanguage(language) {
		errorMsg := fmt.Sprintf("dictionary %s already exists", language)
		r.logger.LogErr(errorMsg)
		return nil, errors.New(errorMsg)
	}

	r.Dictionary = &Dictionary{
		Language: language,
	}

	return r.Dictionary, nil
}

func (r *repository) DeleteDictionary() {
	r.Dictionary = nil
}

func (r *repository) existsDictionaryOfLanguage(language string) bool {

	return r.existsDictionary() && r.Dictionary.Language == language
}

func (r *repository) existsDictionary() bool {
	return r.Dictionary != nil
}

func (r *repository) ListDictionary() []*Dictionary {
	dictionaryList := make([]*Dictionary, 0)
	return append(dictionaryList, r.Dictionary)
}

func (r *repository) AddWord(word *Word) error {
	if !r.existsDictionary() {
		r.logger.LogErr(NO_DICTIONARY)
		return errors.New(NO_DICTIONARY)
	}

	r.logger.LogInfo("Added new word " + word.Label)
	r.Dictionary.Words = append(r.Dictionary.Words, word)
	return nil
}

func (r *repository) ListWords() ([]*Word, error) {
	if !r.existsDictionary() {
		r.logger.LogErr(NO_DICTIONARY)
		return nil, errors.New(NO_DICTIONARY)
	}
	return r.Dictionary.Words, nil
}
