// Package logger 封装 mqtt 客户端内部日志接口
package logger

import (
	"github.com/cryring/Heimdall/internal/log"
)

// DebugLogger mqtt debug logger
type DebugLogger struct{}

// Println implement mqtt logger interface
func (DebugLogger) Println(v ...interface{}) {
	log.Debug(v...)
}

// Printf implement mqtt logger interface
func (DebugLogger) Printf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// ErrorLogger mqtt error logger
type ErrorLogger struct{}

// Println implement mqtt logger interface
func (ErrorLogger) Println(v ...interface{}) {
	log.Error(v...)
}

// Printf implement mqtt logger interface
func (ErrorLogger) Printf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}
