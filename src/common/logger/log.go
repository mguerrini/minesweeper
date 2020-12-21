package logger

var log Logger

type Logger interface {
	Info(msg string)
	Error(msg string)
}

func init () {
	log = NewConsoleLog()
}

func Info(msg string ){
	log.Info(msg)
}

func Error(msg string ){
	log.Error(msg)
}