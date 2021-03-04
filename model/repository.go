package model

import "errors"

type Repository interface {
	CreateDictionary(language string) *Dictionary
	AddWord(word *Word) error
	ListWords() []*Word
}

type repository struct {
	Dictionary *Dictionary
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) CreateDictionary(language string) *Dictionary {
	r.Dictionary = &Dictionary{
		Language: language,
	}
	return r.Dictionary
}

func (r *repository) AddWord(word *Word) error {
	if r.Dictionary == nil {
		return errors.New("no dictionary available")
	}

	r.Dictionary.Words = append(r.Dictionary.Words, word)
	return nil
}

func (r *repository) ListWords() []*Word {
	return r.Dictionary.Words
}

func (r *repository) ListDictionary() []*Dictionary {
	dictionaryList := make([]*Dictionary, 0)
	return append(dictionaryList, r.Dictionary)
}
