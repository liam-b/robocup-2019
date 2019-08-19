package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"
	"strconv"
)

const (
	End = "\x1b[0m"
	Bold = "\x1b[1m"
	Underline = "\x1b[4m"
	Black = "\x1b[30m"
	Red = "\x1b[31m"
	Green = "\x1b[32m"
	Yellow = "\x1b[33m"
	Blue = "\x1b[34m"
	Purple = "\x1b[35m"
	Cyan = "\x1b[36m"
	White = "\x1b[37m"
)

var (
	counter = 0
	startTime = time.Now()
)

func Print(text ...interface{}) {
	fmt.Println(Black + timeDifference() + " [" + pad(strconv.Itoa(counter), 5) + "]" + End + " " + Cyan + traceCallingFunc() + End + " " + format(text))
	counter++
}

func traceCallingFunc() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frames.Next()
	frames.Next()
	frame, _ := frames.Next()
	name := strings.Split(frame.Function, ".")[1]
	return name
}

func format(text ...interface{}) string {
	value := []rune(fmt.Sprint(text))
	return string(value[2:len(value) - 2])
}

func timeDifference() string {
	return pad(strconv.Itoa(int(time.Since(startTime).Minutes())), 2) + ":" + pad(strconv.Itoa(int(time.Since(startTime).Seconds()) % 60), 2)
}

func pad(str string, plength int) string {
	for i := len(str); i < plength; i++ {
		str = "0" + str
	}
	return str
}