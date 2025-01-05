package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func InitLogger(fname string) {
	writeSyncer := getLogWriter(fname)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(fname string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fname,
		MaxSize:    50,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Fatalf(fmt string, v ...interface{}) {
	logger.Fatalf(fmt, v...)
}

func Errorf(fmt string, v ...interface{}) {
	logger.Errorf(fmt, v...)
}

func Infof(fmt string, v ...interface{}) {
	logger.Infof(fmt, v...)
}

func Warnf(fmt string, v ...interface{}) {
	logger.Warnf(fmt, v...)
}

func Debugf(fmt string, v ...interface{}) {
	logger.Debugf(fmt, v...)
}
