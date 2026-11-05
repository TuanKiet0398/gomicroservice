package main

import "fmt"

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogWarning
	LogError
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
