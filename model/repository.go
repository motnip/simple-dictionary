package model

import "errors"

type Repository interface {
	CreateDictionary(language string) (*Dictionary, error)
	AddWord(word *Word) error
	ListWords() ([]*Word, error)
}

type repository struct {
	Dictionary *Dictionary
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) CreateDictionary(language string) (*Dictionary, error) {

	if r.existsDictionaryOfLanguage(language) {
		return nil, errors.New("dictionary already exists")
	}

	r.Dictionary = &Dictionary{
		Language: language,
	}

	return r.Dictionary, nil
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
		return errors.New("no dictionary available")
	}

	r.Dictionary.Words = append(r.Dictionary.Words, word)
	return nil
}

func (r *repository) ListWords() ([]*Word, error) {
	if !r.existsDictionary() {
		return nil, errors.New("no dictionary available")
	}
	return r.Dictionary.Words, nil
}
