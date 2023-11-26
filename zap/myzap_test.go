package zap

import (
	"go.uber.org/zap"
	"testing"
)

func TestInit(t *testing.T) {

	zap.L().Info("request received", zap.String(, "ok"), zap.Int("test cnt", 1))
	zap.L().Debug("test debug", zap.String("test String", "ok"), zap.Int("test cnt", 2))
	zap.L().Error("test error", zap.String("test String", "ok"), zap.Int("test cnt", 3))
}
