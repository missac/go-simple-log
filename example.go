package main

import (
	"muslog"
)

func main() {
	muslog.InitLog(muslog.DEBUG, "", muslog.LevelTrace)
	muslog.Info("test is for test")

	muslog.SetLevel(muslog.LevelError)

	muslog.Fatal("fatal error, and progarm exit")
}
