package usecase

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/rymiyamoto/memer-server/service"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld_NewHelloWorld(t *testing.T) {
	t.Parallel()

	got := NewHelloWorld()
	assert.Equal(t, got, &HelloWorld{
		helloWorldService: service.NewHelloWorld(),
	})
}

func TestHelloWorld_Get(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	helloWorldService := service.NewMockIHelloWorld(mockCtrl)
	s := &HelloWorld{
		helloWorldService: helloWorldService,
	}

	helloWorldService.EXPECT().
		Get().
		Return("test")

	got := s.Get()
	assert.Equal(t, got, "test")
}
