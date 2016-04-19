package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	mu     sync.RWMutex
	level  Level
	writer *log.Logger
)

type Level uint8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

func SetLogger(w *log.Logger) {
	mu.Lock()
	defer mu.Unlock()

	writer = w
}

func SetLogLevel(v Level) {
	mu.Lock()
	defer mu.Unlock()

	level = v
}

func write(l Level, msg string) {
	mu.RLock()
	defer mu.RUnlock()

	if level > l {
		return
	}
	if writer == nil {
		fmt.Println(msg)
		return
	}

	writer.Println(msg)
}

func Debug(m string) {
	write(LevelDebug, m)
}

func Info(m string) {
	write(LevelInfo, m)
}

func Warning(m string) {
	write(LevelWarning, m)
}

func Error(m string) {
	write(LevelError, m)
}

func Fatal(m string) {
	write(LevelFatal, m)
	os.Exit(1)
}
