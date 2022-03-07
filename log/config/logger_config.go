package config

type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

var config Config

type Config struct {
	// Level 日志等级
	Level Level `json:"level" yaml:"level"`
	// Encoding 日志形式
	Encoding string `json:"encoding" yaml:"encoding"`
	// DisableCaller 调用文件位置
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
	// DisableStacktrace and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
	// OutputPaths 日志输出路径
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
	// ErrorOutputPaths 错误输出路径
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`
}

func GetConfig() Config {
	return config
}
