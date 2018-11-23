package model

type Logger interface {
	Errorf(format string, v ...interface{})
	Panic(v ...interface{})
}
