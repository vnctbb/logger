package log

import (
	"log"
	"path"
	"runtime"
)

func Info(message string) {

	funcName, file, line, ok := runtime.Caller(1) // ok is a boolean

	if ok {
		_ = runtime.FuncForPC(funcName).Name() // function to convert funcName in string
	}

	fileName := path.Base(file)

	log.Println(fileName, "line", line, "- [INFO] -", message)

}

func main() {

	Info("Bonjour")
}
