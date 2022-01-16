package repository_test

import (
	"testing"

	"github.com/rymiyamoto/memer-server/repository"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld_NewHelloWorld(t *testing.T) {
	t.Parallel()

	got := repository.NewHelloWorld()
	assert.IsType(t, got, &repository.HelloWorld{})
}

func TestHelloWorld_Get(t *testing.T) {
	t.Parallel()

	r := repository.NewHelloWorld()
	got := r.Get()
	assert.Equal(t, got, "hello world")
}
