package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld_NewHelloWorld(t *testing.T) {
	t.Parallel()

	got := NewHelloWorld()
	assert.Equal(t, got, &HelloWorld{
		text: "hello world",
	})
}

func TestHelloWorld_Get(t *testing.T) {
	t.Parallel()

	r := NewHelloWorld()
	got := r.Get()
	assert.Equal(t, got, "hello world")
}
