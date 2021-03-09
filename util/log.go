package util

import (
	"fmt"
	"log"
	"runtime/debug"
)

func init() {
	/*dir := GetPath()
	logFile, err := os.Open(dir + "/log/judger.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)*/
}

func LogFatalln(v ...interface{}) {
	v = append(v, fmt.Sprintf("\n%s", debug.Stack()))
	log.Fatalln(v)
}

func LogFatalf(format string, v ...interface{}) {
	format += "\n%s\n"
	v = append(v, debug.Stack())
	log.Fatalf(format, v)
}

func LogPanicln(v ...interface{}) {
	v = append(v, fmt.Sprintf("\n%s", debug.Stack()))
	log.Panicln(v)
}

func LogPanicf(format string, v ...interface{}) {
	format += "\n%s\n"
	v = append(v, debug.Stack())
	log.Panicf(format, v)
}
func LogPrintln(v ...interface{}) {
	v = append(v, fmt.Sprintf("\n%s", debug.Stack()))
	log.Println(v)
}

func LogPrintf(format string, v ...interface{}) {
	format += "\n%s\n"
	v = append(v, debug.Stack())
	log.Println(format, v)
}
