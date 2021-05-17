package main

import "fmt"

var NoConsole bool = false

func ReadLine() string {
	var s string
	if NoConsole {
		s = ""
	} else {
		fmt.Scanln(&s)
	}
	return s
}

func Prompt(msg string) string {
	print(msg)
	return ReadLine()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
