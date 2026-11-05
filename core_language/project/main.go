package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)


var levelNames = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
}

func (level LogLevel) String() string {
	if level < LevelTrace || level > LevelError {
		return "UNKNOWN"
	}
	return levelNames[level]
}

func printLogLevel(level LogLevel) {
	fmt.Println("Log Level:", level.String())
}


func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
}