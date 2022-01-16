package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rymiyamoto/memer-server/conf"
	"github.com/rymiyamoto/memer-server/middleware"
	"github.com/rymiyamoto/memer-server/util/log"
)

func main() {
	logger, err := log.NewLogger(log.LoggerConfig{
		UseDebug: !conf.IsProd(),
		Encoding: "json",
		Output:   "stdout",
	})
	if err != nil {
		panic(fmt.Errorf("failed to log.NewLogger. err: %w", err))
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			panic(fmt.Errorf("failed to logger.Sync. err: %w", err))
		}
	}()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.ZapLogger(logger))
	e.Use(echoMiddleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}

// hello Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
