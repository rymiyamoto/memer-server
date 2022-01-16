package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rymiyamoto/memer-server/conf"
	"github.com/rymiyamoto/memer-server/docs"
	"github.com/rymiyamoto/memer-server/middleware"
	"github.com/rymiyamoto/memer-server/router"
	"github.com/rymiyamoto/memer-server/util/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	// Timeout timeout
	Timeout = 10 * time.Second
	// Buffer buffer
	Buffer = 100
)

func main() {
	isProd := conf.IsProd()
	logger, err := log.NewLogger(log.LoggerConfig{
		UseDebug: !isProd,
		Encoding: "json",
		Output:   "stdout",
	})
	if err != nil { // nolint:wsl
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
	e.Use(echoMiddleware.GzipWithConfig(echoMiddleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.ZapLogger(logger))
	e.Use(echoMiddleware.Recover())

	// Routing
	if !isProd {
		docs.SwaggerInfo.Host = "localhost:1324"

		e.GET("/swagger/*", echoSwagger.EchoWrapHandler(echoSwagger.URL("/swagger/doc.json")))
	}

	router.Init(e)

	// Start server
	go func() {
		if err := e.Start(":1324"); err != nil {
			e.Logger.Info("shutting down the server ")
		}
	}()

	quit := make(chan os.Signal, Buffer)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
