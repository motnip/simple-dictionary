package service

import (
	"github.com/motnip/sermo/model"
	"github.com/motnip/sermo/system"
)

type wordService struct {
	repository model.Repository
	log        *system.SermoLog
}

type WordService interface {
	SaveWord(w *model.Word) error
}

func NewWordService(repository model.Repository) wordService {
	return wordService{
		repository: repository,
		log:        system.NewLog(),
	}
}

func (s *wordService) SaveWord(w *model.Word) error {
	if val, err := w.Validate(); val == false && err != nil {
		return err
	}
	return s.repository.AddWord(w)
}
