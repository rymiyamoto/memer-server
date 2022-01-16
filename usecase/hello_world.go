package usecase

import (
	"github.com/rymiyamoto/memer-server/model"
	"github.com/rymiyamoto/memer-server/service"
)

type (
	// IHelloWorld interface
	IHelloWorld interface {
		Get() *model.HelloWorld
	}

	// HelloWorld struct
	HelloWorld struct {
		helloWorldService service.IHelloWorld
	}
)

// NewHelloWorld initialize
func NewHelloWorld() IHelloWorld {
	return &HelloWorld{
		helloWorldService: service.NewHelloWorld(),
	}
}

// Get get
func (u *HelloWorld) Get() *model.HelloWorld {
	return u.helloWorldService.Get()
}
