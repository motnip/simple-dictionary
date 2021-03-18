package web

import (
	"github.com/golang/mock/gomock"
	mock_controller "sermo/mocks/controller"
	"testing"
)

func TestNewRouter(t *testing.T) {

	//given
	ctr := gomock.NewController(t)
	restFullController := mock_controller.NewMockController(ctr)

	//when
	sut := NewRouter(restFullController)

	//than
	sut.Init()

	if sut.router.GetRoute("createDictionary") == nil {
		t.Errorf("No path returend: got %v want %v", sut.router.Path("/dictionary"), "not nil")
	}

	if sut.router.GetRoute("createDictionary") == nil {
		t.Errorf("No path returend: got %v want %v", sut.router.Path("/dictionary"), "not nil")
	}
}
