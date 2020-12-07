package log

import (
	"fmt"

	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/log"
)

func Trace(msg string) {
	host.ProxyLog(log.Trace, utils.StringToBytePtr(msg), len(msg))
}

func Tracef(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Trace, utils.StringToBytePtr(msg), len(msg))
}

func Debug(msg string) {
	host.ProxyLog(log.Debug, utils.StringToBytePtr(msg), len(msg))
}

func Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Debug, utils.StringToBytePtr(msg), len(msg))
}

func Info(msg string) {
	host.ProxyLog(log.Info, utils.StringToBytePtr(msg), len(msg))
}

func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Info, utils.StringToBytePtr(msg), len(msg))
}

func Warn(msg string) {
	host.ProxyLog(log.Warn, utils.StringToBytePtr(msg), len(msg))
}

func Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Warn, utils.StringToBytePtr(msg), len(msg))
}

func Error(msg string) {
	host.ProxyLog(log.Error, utils.StringToBytePtr(msg), len(msg))
}

func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Error, utils.StringToBytePtr(msg), len(msg))
}

func Critical(msg string) {
	host.ProxyLog(log.Critical, utils.StringToBytePtr(msg), len(msg))
}

func Criticalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Critical, utils.StringToBytePtr(msg), len(msg))
}

func Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	host.ProxyLog(log.Error, utils.StringToBytePtr(msg), len(msg))
	panic("aborted from fatal operation")
}

func Fatal(msg string) {
	host.ProxyLog(log.Error, utils.StringToBytePtr(msg), len(msg))
	panic("aborted from fatal operation")
}
