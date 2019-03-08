package logger

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"runtime"
	"regexp"
)

const (
	TextEnd       = "\x1b[0m"
	TextBold      = "\x1b[1m"
	TextUnderline = "\x1b[4m"
	TextBlack     = "\x1b[30m"
	TextRed       = "\x1b[31m"
	TextGreen     = "\x1b[32m"
	TextYellow    = "\x1b[33m"
	TextBlue      = "\x1b[34m"
	TextPurple    = "\x1b[35m"
	TextCyan      = "\x1b[36m"
	TextWhite     = "\x1b[37m"
)

var (
	callerRegex, _ = regexp.Compile("[A-z\\.0-9]+$")

	startTime time.Time
	counter   int
	level     int
	status    *string
	event     *string
)

func Init(_level int, _status *string, _event *string) {
	fmt.Println("  ____   ___  ____   ___   ____ _   _ ____    ____   ___  _  ___  \n |  _ \\ / _ \\| __ ) / _ \\ / ___| | | |  _ \\  |___ \\ / _ \\/ |/ _ \\ \n | |_) | | | |  _ \\| | | | |   | | | | |_) |   __) | | | | | (_) |\n |  _ <| |_| | |_) | |_| | |___| |_| |  __/   / __/| |_| | |\\__, |\n |_| \\_\\\\___/|____/ \\___/ \\____|\\___/|_|     |_____|\\___/|_|  /_/" + "\n")
	level = _level
	status = _status
	event = _event
	startTime = time.Now()
	counter = 0
}

func State(text ...interface{}) {
	if level >= 6 {
		print("state", TextBlack, format(text))
	}
}

func Trace(text ...interface{}) {
	if level >= 5 {
		print("trace", TextWhite, format(text))
	}
}

func Debug(text ...interface{}) {
	if level >= 4 {
		print("debug", TextGreen, format(text))
	}
}

func Info(text ...interface{}) {
	if level >= 3 {
		print("info ", TextBlue, format(text))
	}
}

func Warn(text ...interface{}) {
	if level >= 2 {
		print("warn ", TextYellow, format(text))
	}
}

func Error(text ...interface{}) {
	if level >= 1 {
		print("error", TextRed, format(text))
	}
}

func Fatal(text ...interface{}) {
	if level >= 0 {
		print("fatal", TextRed, format(text))
	}
}

func format(text ...interface{}) string {
	value := []rune(fmt.Sprint(text))
	return string(value[2:len(value) - 2])
}

func print(level string, color string, text string) {
	fmt.Println(TextBlack + timeDifference() + " [" + pad(strconv.Itoa(counter), 5) + "]" + TextEnd + " " + TextBold + color + strings.ToUpper((string)(level)) + TextEnd + " " + TextPurple + getContext() + TextEnd + " " + text + " " + TextBlack + "(" + callerFunction() + ")" + TextEnd)
	counter++
}

func timeDifference() string {
	return pad(strconv.Itoa(int(time.Since(startTime).Minutes())), 2) + ":" + pad(strconv.Itoa(int(time.Since(startTime).Seconds()) % 60), 2)
}

func getContext() string {
	if *status != "" {
		return TextCyan + "[" + *event + "]" + TextPurple + *status
	} else {
		return TextCyan + "[.]" + TextPurple + "{nil}"
	}
}

func pad(str string, plength int) string {
	for i := len(str); i < plength; i++ {
		str = "0" + str
	}
	return str
}

func callerFunction() string {
	pointers := make([]uintptr, 15)
	callers := runtime.Callers(4, pointers)
	frames := runtime.CallersFrames(pointers[:callers])
	frame, _ := frames.Next()
	return callerRegex.FindString(frame.Function)
}