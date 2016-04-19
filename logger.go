package logger

import (
	"errors"
	"fmt"
	"log/syslog"
	"os"
	"sync"
)

var (
	mu     sync.RWMutex
	level  Level
	writer *syslog.Writer
)

type Level uint8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

func SetSyslog(w *syslog.Writer) {
	mu.Lock()
	defer mu.Unlock()

	writer = w
}

func SetLogLevel(v Level) {
	mu.Lock()
	defer mu.Unlock()

	level = v
}

func write(l Level, msg string) error {
	mu.RLock()
	defer mu.RUnlock()

	if level > l {
		return nil
	}
	if writer == nil {
		fmt.Println(msg)
		return nil
	}

	switch l {
	case LevelDebug:
		return writer.Debug(msg)
	case LevelInfo:
		return writer.Info(msg)
	case LevelWarning:
		return writer.Warning(msg)
	case LevelError:
		return writer.Err(msg)
	case LevelFatal:
		writer.Crit(msg)
		os.Exit(1)
	default:
		return errors.New("unexpected log level")
	}

	// Makes compiler happy
	return nil
}

func Debug(m string) error {
	return write(LevelDebug, m)
}

func Info(m string) error {
	return write(LevelInfo, m)
}

func Warning(m string) error {
	return write(LevelWarning, m)
}

func Error(m string) error {
	return write(LevelError, m)
}

func Fatal(m string) error {
	return write(LevelFatal, m)
}
