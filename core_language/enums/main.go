package main

import "fmt"

const (
	Sunday = 0
	Monday = 1
	Tuesday = 2
	Wednesday = 3
	Thursday = 4
	Friday = 5
	Saturday = 6
)

type LogLevel int

const (
	LogError=0 LogLevel = iota
	LogInfo=1
	LogWarning=2
	LogError=3
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