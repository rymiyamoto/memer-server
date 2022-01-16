package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rymiyamoto/memer-server/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}.Build(
		zap.AddCaller(),
		zap.Fields(
			zap.String("permanent", "value"),
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to zap.Build. err: %w", err))
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
