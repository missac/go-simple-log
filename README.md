usage:


import "muslog"
muslog.InitLog(muslog.DEBUG, "", muslog.LevelTrace) //the log will write at stdout with DEBUG mode
muslog.Trace("log should be print at stdout")

or to write the log into a file:

logFile, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
muslog.SetOutput(logFile)
muslog.Trace("log should be at test.log")

or 

muslog.InitLog(muslog.RELEASE, "test.log", muslog.LevelTrace) // log will write to file test.log with RELASE mode
muslog.Trace("log should be print at stdout")
