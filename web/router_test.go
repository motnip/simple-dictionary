package web

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	mock_controller "github.com/motnip/sermo/mocks/controller"
	"net/http"
	"testing"
)

func TestRouter_InitRoute(t *testing.T) {

	//given
	ctr := gomock.NewController(t)
	mockController := mock_controller.NewMockDictionaryController(ctr)

	newRoute := &Route{
		Path:     "/foo",
		Function: mockController.CreateDictionary,
		Method:   http.MethodGet,
		Name:     "fooGet",
	}

	expected := mux.NewRouter().StrictSlash(true)
	expected.HandleFunc(newRoute.Path, newRoute.Function).Name(newRoute.Name).Methods(newRoute.Method)

	//when
	sut := NewRouter()

	//than
	sut.InitRoute(newRoute)

	if sut.router.GetRoute("fooGet") == nil {
		t.Errorf("No path returend: got %v want %v", nil, "not nil")
	}

	if path, err := sut.router.GetRoute("fooGet").GetPathTemplate(); err == nil && path != "/foo" {
		t.Errorf("No path returend: got %v want %v", path, "/foo")
	}

	if sut.router.GetRoute("fooGet").GetHandler() == nil {
		t.Errorf("No path returend: got %v want %v", nil, "not nil")
	}
}
