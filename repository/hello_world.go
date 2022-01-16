package repository

import "github.com/rymiyamoto/memer-server/model"

type (
	// IHelloWorld interface
	IHelloWorld interface {
		Get() *model.HelloWorld
	}

	// HelloWorld struct
	HelloWorld struct {
		text string
	}
)

// NewHelloWorld initialize
func NewHelloWorld() IHelloWorld {
	return &HelloWorld{
		text: "hello world",
	}
}

// Get get
func (r *HelloWorld) Get() *model.HelloWorld {
	return &model.HelloWorld{
		Text: r.text,
	}
}
