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
		Meaning:  "",
		Sentence: "sentence",
	}

	controller := gomock.NewController(t)
	wordValidatorMock := mock_model.NewMockValidator(controller)
	repositoryMock := mock_model.NewMockRepository(controller)

	wordValidatorMock.EXPECT().Validate(gomock.Any()).Times(1).Return(true, nil)
	repositoryMock.EXPECT().AddWord(gomock.Any()).Return(nil)

	sut := NewWordService(repositoryMock, wordValidatorMock)
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
	wordValidatorMock := mock_model.NewMockValidator(controller)

	wordValidatorMock.EXPECT().Validate(gomock.Any()).Times(1).Return(true, nil)
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(1).Return(expectedError)

	sut := NewWordService(repositoryMock, wordValidatorMock)
	err := sut.SaveWord(&word)

	if err == nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Router returned unexpected value: got %v want %v", err.Error(), expectedError.Error())
	}
}

func Test_saveWord_errorOnValidate_Failed(t *testing.T) {
	word := model.Word{
		Label:    "label",
		Meaning:  "meaning",
		Sentence: "sentence",
	}
	expectedError := errors.New("forced error - validation failed")

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	wordValidatorMock := mock_model.NewMockValidator(controller)

	wordValidatorMock.EXPECT().Validate(gomock.Any()).Times(1).Return(false, expectedError)
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0).Return(expectedError)

	sut := NewWordService(repositoryMock, wordValidatorMock)
	err := sut.SaveWord(&word)

	if err == nil {
		t.Errorf("Router returned unexpected value: got %v want %v", err, "nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Router returned unexpected value: got %v want %v", err.Error(), expectedError.Error())
	}
}
