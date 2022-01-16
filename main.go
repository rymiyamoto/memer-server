package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rymiyamoto/memer-server/conf"
	"github.com/rymiyamoto/memer-server/middleware"
	"github.com/rymiyamoto/memer-server/router"
	"github.com/rymiyamoto/memer-server/util/log"
)

const (
	// Timeout timeout
	Timeout = 10 * time.Second
	// Buffer buffer
	Buffer = 100
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

	// Routing
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
