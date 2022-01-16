package service

import "github.com/rymiyamoto/memer-server/repository"

type (
	// IHelloWorld interface
	IHelloWorld interface {
		Get() string
	}

	// HelloWorld struct
	HelloWorld struct {
		helloWorldRepository repository.IHelloWorld
	}
)

// NewHelloWorld initialize
func NewHelloWorld() IHelloWorld {
	return &HelloWorld{
		helloWorldRepository: repository.NewHelloWorld(),
	}
}

// Get get
func (s *HelloWorld) Get() string {
	return s.helloWorldRepository.Get()
}
