package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLog_NewLogger(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		config := LoggerConfig{
			UseDebug: true,
			Encoding: "json",
			Output:   "stdout",
		}

		got, err := NewLogger(config)
		assert.IsType(t, got, &zap.Logger{})
		assert.Nil(t, err)
	})
	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		config := LoggerConfig{
			UseDebug: true,
			Encoding: "hoge",
			Output:   "fuga",
		}

		got, err := NewLogger(config)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
}
