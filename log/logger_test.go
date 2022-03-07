package log

import (
	"context"
	"errors"
	"testing"
)

func TestInfo(t *testing.T) {

	Info(context.TODO(), "%s  say hello", "amy")
}

func TestError(t *testing.T) {

	Error(context.TODO(), errors.New("测试err"), "%s errtest", "Tom")
}

func TestDPanic(t *testing.T) {
	DPanic(context.TODO(), errors.New("测试err"), "%s errtest", "Tom")
}
