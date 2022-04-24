package router

import (
	"github.com/labstack/echo/v4"
)

// Init initialize
func Init(e *echo.Echo) {
	InitV1(e)
}
