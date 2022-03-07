package log

import (
	"base/log/config"
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

type Logger struct {
}

func Info(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.Info(fmt.Sprintf(format, args...))
}
func Error(ctx context.Context, err error, format string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.Error(fmt.Sprintf(format, args...), zap.Error(err))

}

func Debug(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.Debug(fmt.Sprintf(format, args...))

}
func Warn(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.Warn(fmt.Sprintf(format, args...))

}
func DPanic(ctx context.Context, err error, format string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.DPanic(fmt.Sprintf(format, args...), zap.Error(err))
}
func Panic(ctx context.Context, err error, format string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.Panic(fmt.Sprintf(format, args...), zap.Error(err))
}
func initLogger() {
	once.Do(func() {
		var err error
		logger, err = getLogger()
		if err != nil {
			panic(err)
		}
	})
}
func getLogger() (*zap.Logger, error) {
	zapConfig, _ := config.GetConfig().ToZapConfig()
	l, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return l, nil
}
