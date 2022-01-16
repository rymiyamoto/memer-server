package repository

type (
	// IHelloWorld interface
	IHelloWorld interface {
		Get() string
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
func (r *HelloWorld) Get() string {
	return r.text
}
