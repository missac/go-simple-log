usage:
DEBUG:
import "muslog"
muslog.InitLog(muslog.DEBUG, "", muslog.LevelTrace)
muslog.Trace("log should be print at stdout")

or to write the log into a file:
logFile, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
muslog.SetOutput(logFile)
muslog.Trace("log should be at test.log")

or 

muslog.InitLog(muslog.RELEASE, "test.log", muslog.LevelTrace)
muslog.Trace("log should be print at stdout")
