package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rymiyamoto/memer-server/handler"
)

type (
	// IV1 interface
	IV1 interface {
		withoutToken(g *echo.Group)
	}

	// V1 struct
	V1 struct {
		helloWorldHandler handler.IHelloWorld
	}
)

// InitV1 initialize
func InitV1(e *echo.Echo) {
	var r IV1 = &V1{
		helloWorldHandler: handler.NewHelloWorld(),
	}

	g := e.Group("1.0")
	g.Use()
	{
		r.withoutToken(g)
	}
}

func (r V1) withoutToken(g *echo.Group) {
	g.GET("/hello_world", r.helloWorldHandler.Index)
}
