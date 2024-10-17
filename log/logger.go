package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func createEncoder(symbol string) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:      "time",
		MessageKey:   "msg",
		EncodeTime:   zapcore.TimeEncoderOfLayout(symbol + " 02/01/2006 - 15:04:05 ➡"),
		EncodeLevel:  func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {},
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {},
	}
}

func createLogger(encoderConfig zapcore.EncoderConfig) *zap.Logger {
	configZap := zap.NewProductionConfig()
	configZap.EncoderConfig = encoderConfig
	configZap.Encoding = "console"
	logger, _ := configZap.Build()
	return logger
}

func Info(message string) {
	logger = createLogger(createEncoder("🔘"))
	defer logger.Sync()
	logger.Info(message)
}

func Error(message string, err error) {
	logger = createLogger(createEncoder("🔴"))
	defer logger.Sync()
	logger.Error(message, zap.Error(err))
}
