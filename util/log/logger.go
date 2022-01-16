package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	// LoggerConfig struct
	LoggerConfig struct {
		UseDebug bool
		IsDev    bool
		Encoding string
		Output   string
	}
)

// NewLogger initialize
func NewLogger(config LoggerConfig) (*zap.Logger, error) {
	level := zap.InfoLevel
	if config.UseDebug {
		level = zap.DebugLevel
	}

	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: config.IsDev,
		Encoding:    config.Encoding,
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
		OutputPaths:      []string{config.Output},
		ErrorOutputPaths: []string{config.Output},
	}.Build(
		zap.AddCaller(),
		zap.Fields(
			zap.String("permanent", "value"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to zap.Build. err: %w", err)
	}

	return logger, nil
}
