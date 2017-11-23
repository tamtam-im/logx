package logx

import (
	"fmt"
)

type Log struct {
	prefix string
	tags   []string

	appender Appender
	callDepth   int
}

// NewLog returns new log. Output Writer must be thread safe.
func NewLog(appender Appender, prefix string, tags ...string) (res *Log) {
	res = &Log{
		tags:        tags,
		prefix:      prefix,
		appender: appender,
		callDepth:   2,
	}
	return
}

// GetLog returns new independent log instance with given prefix.
func (l *Log) GetLog(prefix string, tags ...string) (res *Log) {
	res = NewLog(l.appender, prefix, tags...)
	return
}

// Prefix returns log prefix.
func (l *Log) Prefix() (res string) {
	res = string(l.prefix)
	return
}

// Tags returns log tags.
func (l *Log) Tags() (res []string) {
	res = l.tags
	return
}

// Print is synonym to Info used for compatibility.
func (l *Log) Print(v ...interface{}) {
	l.appendLine(lInfo, fmt.Sprint(v...))
}

// Printf is synonym to Infof used for compatibility.
func (l *Log) Printf(format string, v ...interface{}) {
	l.appendLine(lInfo, fmt.Sprintf(format, v...))
}

// Info logs value with INFO severity level.
func (l *Log) Info(v ...interface{}) {
	l.appendLine(lInfo, fmt.Sprint(v...))
}

// Infof logs formatted value with INFO severity level.
func (l *Log) Infof(format string, v ...interface{}) {
	l.appendLine(lInfo, fmt.Sprintf(format, v...))
}

// Warning logs value with WARNING severity level.
func (l *Log) Warning(v ...interface{}) {
	l.appendLine(lWarning, fmt.Sprint(v...))
}

// Warningf logs formatted value with WARNING severity level.
func (l *Log) Warningf(format string, v ...interface{}) {
	l.appendLine(lWarning, fmt.Sprintf(format, v...))
}

// Error logs value with ERROR severity level.
func (l *Log) Error(v ...interface{}) {
	l.appendLine(lError, fmt.Sprint(v...))
}

// Errorf logs formatted value with ERROR severity level.
func (l *Log) Errorf(format string, v ...interface{}) {
	l.appendLine(lError, fmt.Sprintf(format, v...))
}

// Critical logs value with CRITICAL severity level.
func (l *Log) Critical(v ...interface{}) {
	l.appendLine(lCritical, fmt.Sprint(v...))
}

// Criticalf logs formatted value with CRITICAL severity level.
func (l *Log) Criticalf(format string, v ...interface{}) {
	l.appendLine(lCritical, fmt.Sprintf(format, v...))
}

func (l *Log) appendLine(level, line string) {
	l.appender.Append(level, l.prefix, line, l.tags...)
}
