package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {

	writable := io.Writer(os.Stdout)

	logger := log.New(writable, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writable, "DEBBUG: ", logger.Flags()),
		info:    log.New(writable, "INFO: ", logger.Flags()),
		warning: log.New(writable, "WARNING:  ", logger.Flags()),
		err:     log.New(writable, "ERROR:  ", logger.Flags()),
		writer:  writable,
	}
}

// Create non-formated logs
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Create formatade logs
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnigf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
