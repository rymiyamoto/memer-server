package log_test

import (
	"testing"

	"github.com/rymiyamoto/memer-server/util/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLog_NewLogger(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		config := log.LoggerConfig{
			UseDebug: true,
			Encoding: "json",
			Output:   "stdout",
		}

		got, err := log.NewLogger(config)
		assert.IsType(t, got, &zap.Logger{})
		assert.Nil(t, err)
	})
	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		config := log.LoggerConfig{
			UseDebug: true,
			Encoding: "hoge",
			Output:   "fuga",
		}

		got, err := log.NewLogger(config)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}
