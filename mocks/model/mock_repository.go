// Code generated by MockGen. DO NOT EDIT.
// Source: model/repository.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/motnip/sermo/model"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddWord mocks base method.
func (m *MockRepository) AddWord(word *model.Word) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWord", word)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWord indicates an expected call of AddWord.
func (mr *MockRepositoryMockRecorder) AddWord(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWord", reflect.TypeOf((*MockRepository)(nil).AddWord), word)
}

// CreateDictionary mocks base method.
func (m *MockRepository) CreateDictionary(language string) (*model.Dictionary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDictionary", language)
	ret0, _ := ret[0].(*model.Dictionary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDictionary indicates an expected call of CreateDictionary.
func (mr *MockRepositoryMockRecorder) CreateDictionary(language interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDictionary", reflect.TypeOf((*MockRepository)(nil).CreateDictionary), language)
}

// ListDictionary mocks base method.
func (m *MockRepository) ListDictionary() []*model.Dictionary {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDictionary")
	ret0, _ := ret[0].([]*model.Dictionary)
	return ret0
}

// ListDictionary indicates an expected call of ListDictionary.
func (mr *MockRepositoryMockRecorder) ListDictionary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDictionary", reflect.TypeOf((*MockRepository)(nil).ListDictionary))
}

// ListWords mocks base method.
func (m *MockRepository) ListWords() ([]*model.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWords")
	ret0, _ := ret[0].([]*model.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWords indicates an expected call of ListWords.
func (mr *MockRepositoryMockRecorder) ListWords() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWords", reflect.TypeOf((*MockRepository)(nil).ListWords))
}
