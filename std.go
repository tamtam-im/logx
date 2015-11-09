package logx
import "io"

var std *Log = New()

func init() {
	std.calldepth = 3
}

func SetLevel(level string) {
	std.SetLevel(level)
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetPrefix(prefix string)  {
	std.SetPrefix(prefix)
}

func Trace(v ...interface{}) {
	std.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
	std.Tracef(format, v...)
}

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

func Info(v ...interface{}) {
	std.Info(v...)
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

func Warning(v ...interface{}) {
	std.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

func Panic(v ...interface{}) {
	std.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

func OnWarning(err interface{}, v ...interface{}) error {
	return std.OnWarning(err, v...)
}

func OnWarningf(err interface{}, format string, v ...interface{}) error {
	return std.OnWarningf(err, format, v...)
}

func OnError(err interface{}, v ...interface{}) error {
	return std.OnError(err, v...)
}

func OnErrorf(err interface{}, format string, v ...interface{}) error {
	return std.OnErrorf(err, format, v...)
}

func OnFatal(err interface{}, v ...interface{}) {
	std.OnFatal(err, v...)
}

func OnFatalf(err interface{}, format string, v ...interface{}) {
	std.OnFatalf(err, format, v...)
}

func OnPanic(err interface{}, v ...interface{}) {
	std.OnPanic(err, v...)
}

func OnPanicf(err interface{}, format string, v ...interface{}) {
	std.OnPanicf(err, format, v...)
}

