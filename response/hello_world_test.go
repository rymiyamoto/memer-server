package response

import (
	"testing"

	"github.com/rymiyamoto/memer-server/model"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld_NewHelloWorld(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m      *model.HelloWorld
		expect *HelloWorld
	}{
		{m: nil, expect: nil},
		{m: &model.HelloWorld{Text: "test"}, expect: &HelloWorld{Text: "test"}},
	}

	for _, test := range tests {
		got := NewHelloWorld(test.m)
		assert.Equal(t, got, test.expect)
	}
}

func TestHelloWorld_NewHelloWorlds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m      *model.HelloWorlds
		expect *HelloWorlds
	}{
		{m: nil, expect: &HelloWorlds{}},
		{m: &model.HelloWorlds{}, expect: &HelloWorlds{}},
		{m: &model.HelloWorlds{{Text: "test"}}, expect: &HelloWorlds{{Text: "test"}}},
	}

	for _, test := range tests {
		got := NewHelloWorlds(test.m)
		assert.Equal(t, got, test.expect)
	}
}
