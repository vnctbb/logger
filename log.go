package main

import (
	"log"
	"runtime"
)

func Info(message string) {

	funcName,file,line,ok := runtime.Caller(1) // ok is a boolean

	if ok {
		_ = runtime.FuncForPC(funcName).Name() // function to convert funcName in string
	}

	log.Println("-", file, "- line", line, "- [INFO]", message)


}

func main () {

	Info("Bonjour")
}