package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock_model "github.com/motnip/sermo/mocks/model"
	"github.com/motnip/sermo/model"
	"testing"
)

func Test_saveWord_successful(t *testing.T) {
	word := model.Word{
		Label:    "label",
		Meaning:  "Meaning",
		Sentence: "sentence",
	}

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(true)
	repositoryMock.EXPECT().AddWord(gomock.Any()).Return(nil)

	sut := NewWordService(repositoryMock)
	err := sut.SaveWord(&word)

	if err != nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}
}

func Test_saveWord_errorOnSave_Failed(t *testing.T) {
	word := model.Word{
		Label:    "label",
		Meaning:  "meaning",
		Sentence: "sentence",
	}
	expectedError := errors.New("forced error - save word")

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(true)
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(1).Return(expectedError)

	sut := NewWordService(repositoryMock)
	err := sut.SaveWord(&word)

	if err == nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Router returned unexpected value: got %v want %v", err.Error(), expectedError.Error())
	}
}

func Test_SaveWord_noDictionaryAvailable_Failed(t *testing.T) {
	word := model.Word{
		Label:    "label",
		Meaning:  "meaning",
		Sentence: "sentence",
	}
	expectedError := model.NO_DICTIONARY

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(false)
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0)

	sut := NewWordService(repositoryMock)
	err := sut.SaveWord(&word)

	if err == nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Router returned unexpected value: got %v want %v", err.Error(), expectedError)
	}
}

func Test_saveWord_errorOnValidate_Failed(t *testing.T) {
	word := model.Word{
		Label:    "",
		Meaning:  "meaning",
		Sentence: "sentence",
	}
	expectedError := errors.New("forced error - validation failed")

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0).Return(expectedError)

	sut := NewWordService(repositoryMock)
	err := sut.SaveWord(&word)

	if err == nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}
}

func Test_ListWords_succeed(t *testing.T) {
	word := model.Word{
		Label:    "",
		Meaning:  "meaning",
		Sentence: "sentence",
	}

	words := make([]*model.Word, 0)
	words = append(words, &word)

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(true)
	repositoryMock.EXPECT().ListWords().Times(1).Return(words, nil)

	sut := NewWordService(repositoryMock)
	savedWord, err := sut.ListWords()

	if err != nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if len(savedWord) != len(words) {
		t.Errorf("Router returned unexpected value: got %v want %v", len(savedWord), len(words))
	}
}

func Test_ListWords_noDictionaryAvailable_failed(t *testing.T) {

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(false)
	repositoryMock.EXPECT().ListWords().Times(0)

	sut := NewWordService(repositoryMock)
	savedWord, err := sut.ListWords()

	if err == nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if err.Error() != model.NO_DICTIONARY {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if len(savedWord) != 0 {
		t.Errorf("Router returned unexpected value: got %v want %v", len(savedWord), 0)
	}
}
