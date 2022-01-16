package usecase

import (
	"github.com/rymiyamoto/memer-server/service"
)

type (
	// IHelloWorld interface
	IHelloWorld interface {
		Get() string
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
func (u *HelloWorld) Get() string {
	return u.helloWorldService.Get()
}
