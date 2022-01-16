package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rymiyamoto/memer-server/response"
	"github.com/rymiyamoto/memer-server/usecase"
)

type (
	// IHelloWorld interface
	IHelloWorld interface {
		Index(c echo.Context) error
	}

	// HelloWorld struct
	HelloWorld struct {
		helloWorldUsecase usecase.IHelloWorld
	}

	// JSONHelloWorldIndex struct
	JSONHelloWorldIndex struct {
		HelloWorld *response.HelloWorld `json:"hello_world"`
	}
)

// NewHelloWorld initialize
func NewHelloWorld() IHelloWorld {
	return &HelloWorld{
		helloWorldUsecase: usecase.NewHelloWorld(),
	}
}

// Index index
func (h *HelloWorld) Index(c echo.Context) error {
	m := h.helloWorldUsecase.Get()

	return c.JSON(http.StatusOK, &JSONHelloWorldIndex{
		HelloWorld: response.NewHelloWorld(m),
	})
}
