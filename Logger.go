package logging

import (
	"fmt"
	"os"
)

type Logger struct {
	LogLevel int
	DebugFile *os.File
	ErrorFile *os.File
	InfoFile *os.File
	WarnFile *os.File
}

func (l *Logger) Debug(m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_DEBUG {
		fmt.Fprintln(l.DebugFile, m...)
	}
}

func (l *Logger) Debugf(f string, m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_DEBUG {
		fmt.Fprintln(l.DebugFile, fmt.Sprintf(f, m...))
	}
}

func (l *Logger) Error(m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_ERROR {
		fmt.Fprintln(l.ErrorFile, m...)
	}
}


func (l *Logger) Errorf(f string, m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_ERROR {
		fmt.Fprintln(l.ErrorFile, fmt.Sprintf(f, m...))
	}
}
func (l *Logger) Info(m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_INFO {
		fmt.Fprintln(l.InfoFile, m...)
	}
}

func (l *Logger) Infof(f string, m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_INFO {
		fmt.Fprintln(l.InfoFile, fmt.Sprintf(f, m...))
	}
}
func (l *Logger) Warn(m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_WARN {
		fmt.Fprintln(l.WarnFile, m...)
	}
}

func (l *Logger) Warnf(f string, m...interface{}) {
	if l.LogLevel >= LOG_LEVEL_WARN {
		fmt.Fprintln(l.WarnFile, fmt.Sprintf(f, m...))
	}
}

func NewLogger() Logger {
	return Logger{
		DebugFile: os.Stderr,
		ErrorFile: os.Stderr,
		InfoFile: os.Stdout,
		WarnFile: os.Stderr,
		LogLevel: LOG_LEVEL_INFO,
	}
}
