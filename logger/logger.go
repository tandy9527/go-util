package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init(level string, logFilePath string) error {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	if logFilePath == "" {
		logFilePath = "./logs/log.log"
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    100,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoderConfig := encoderConfig
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	fileWriter := zapcore.AddSync(lumberJackLogger)
	consoleWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, zapLevel),
		zapcore.NewCore(consoleEncoder, consoleWriter, zapLevel),
	)

	var err error
	once.Do(func() {
		logger = zap.New(core,
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)
	})
	return err
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Info 级别日志
func Info(msg string, args ...interface{}) {
	logger.Info(fmt.Sprintf(msg, args...))
}

// Debug 级别日志
func Debug(msg string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(msg, args...))
}

// Warn 级别日志
func Warn(msg string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(msg, args...))
}

// Error 级别日志
func Error(msg string, args ...interface{}) {
	logger.Error(fmt.Sprintf(msg, args...))
}
func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}
