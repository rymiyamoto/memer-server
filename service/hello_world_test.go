package service

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/rymiyamoto/memer-server/model"
	"github.com/rymiyamoto/memer-server/repository"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld_NewHelloWorld(t *testing.T) {
	t.Parallel()

	got := NewHelloWorld()
	assert.Equal(t, got, &HelloWorld{
		helloWorldRepository: repository.NewHelloWorld(),
	})
}

func TestHelloWorld_Get(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	helloWorldRepository := repository.NewMockIHelloWorld(mockCtrl)
	s := &HelloWorld{
		helloWorldRepository: helloWorldRepository,
	}
	m := &model.HelloWorld{
		Text: "test",
	}

	helloWorldRepository.EXPECT().
		Get().
		Return(m)

	got := s.Get()
	assert.Equal(t, got, m)
}
