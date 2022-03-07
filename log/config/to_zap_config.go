package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (c Config) ToZapConfig() (zap.Config, error) {

	encoding := "json"
	if c.Encoding == "console" {
		encoding = "encoding"
	}
	outPaths := c.OutputPaths
	errorOutputPaths := c.ErrorOutputPaths
	if len(outPaths) == 0 {
		outPaths = append(outPaths, "stderr")
	}
	if len(errorOutputPaths) == 0 {
		errorOutputPaths = append(errorOutputPaths, "stderr")
	}

	zc := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(c.Level)),
		Development:       false,
		DisableCaller:     c.DisableCaller,
		DisableStacktrace: c.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: encoding,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "logger_msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder},
		OutputPaths:      outPaths,
		ErrorOutputPaths: errorOutputPaths,
		InitialFields:    nil,
	}
	return zc, nil
}
