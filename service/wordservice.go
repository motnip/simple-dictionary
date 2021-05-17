package service

import (
	"github.com/motnip/sermo/model"
	"github.com/motnip/sermo/system"
)

type wordService struct {
	repository model.Repository
	validator  model.Validator
	log        *system.SermoLog
}

func NewWordService(repository model.Repository, validator model.Validator) wordService {
	return wordService{
		repository: repository,
		validator:  validator,
		log:        system.NewLog(),
	}
}

func (s *wordService) SaveWord(w *model.Word) error {
	if val, err := s.validator.Validate(*w); val == false && err != nil {
		return err
	}
	return s.repository.AddWord(w)
}
