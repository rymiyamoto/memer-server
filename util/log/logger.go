package log

import "go.uber.org/zap"

type (
	// Logger struct
	Logger struct {
		Zap *zap.Logger
	}
)
