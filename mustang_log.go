package muslog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelPanic
	LevelFatal
)

var (
	DEBUG   = "debug"
	RELEASE = "release"
)

var lev = LevelTrace
var once sync.Once

var myLog *log.Logger

func InitLog(mode string, logFile string, level LogLevel) error {
	once.Do(func() {
		lev = level
		if mode == DEBUG {
			myLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
			if err != nil {
				log.Fatal(err)
				return
			}
			myLog = log.New(file, "", log.Ldate|log.Ltime)
		}
	})
	return nil
}

func SetOutput(w io.Writer) {
	myLog.SetOutput(w)
}

func SetLevel(level LogLevel) {
	lev = level
}

func Trace(v ...interface{}) {
	if lev > LevelTrace {
		return
	}
	myLog.SetPrefix("TRACE ")
	myLog.Output(2, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	if lev > LevelInfo {
		return
	}
	myLog.SetPrefix("INFO ")
	myLog.Output(2, fmt.Sprintln(v...))
}

func Warning(v ...interface{}) {
	if lev > LevelWarning {
		return
	}
	myLog.SetPrefix("WARNING ")
	myLog.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	if lev > LevelError {
		return
	}
	myLog.SetPrefix("ERROR ")
	myLog.Output(2, fmt.Sprintln(v...))
}

func Panic(v ...interface{}) {
	if lev > LevelPanic {
		return
	}
	myLog.SetPrefix("PANIC ")
	s := fmt.Sprintln(v...)
	myLog.Output(2, s)
	panic(s)
}

func Fatal(v ...interface{}) {
	if lev > LevelFatal {
		return
	}
	myLog.SetPrefix("FATAL ")
	myLog.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
