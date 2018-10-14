package main

import "fmt"
import "strconv"
import "strings"
import "time"

const TEXT_PURPLE = "\x1b[35m"
const TEXT_BOLD = "\x1b[1m"
const TEXT_BLUE = "\x1b[34m"
const TEXT_CYAN = "\x1b[36m"
const TEXT_GREEN = "\x1b[32m"
const TEXT_YELLOW = "\x1b[33m"
const TEXT_RED = "\x1b[31m"
const TEXT_BLACK = "\x1b[30m"
const TEXT_END = "\x1b[0m"
const TEXT_WHITE = ""

type Logger struct {
  machine *StateMachine
  startTime time.Time
  counter int
}

func (logger Logger) new() Logger {
  fmt.Println("  ____   ___  ____   ___   ____ _   _ ____    ____   ___  _  ___  \r\n |  _ \\ / _ \\| __ ) / _ \\ / ___| | | |  _ \\  |___ \\ / _ \\/ |( _ ) \r\n | |_) | | | |  _ \\| | | | |   | | | | |_) |   __) | | | | |/ _ \\ \r\n |  _ <| |_| | |_) | |_| | |___| |_| |  __/   / __/| |_| | | (_) |\r\n |_| \\_\\\\___/|____/ \\___/ \\____|\\___/|_|     |_____|\\___/|_|\\___/ \r\n          ")
  logger.startTime = time.Now()
  logger.counter = 0
  return logger
}

func (logger *Logger) log(text string) {
  logger._print("log", TEXT_GREEN, text)
}

func (logger *Logger) info(text string) {
  logger._print("info", TEXT_BLUE, text)
}

func (logger *Logger) warn(text string) {
  logger._print("warn", TEXT_YELLOW, text)
}

func (logger *Logger) error(text string) {
  logger._print("error", TEXT_RED, text)
}

func (logger *Logger) _print(level string, color string, text string) {
  fmt.Println(TEXT_BLACK + logger._timeDifference() + " [" + pad(strconv.Itoa(logger.counter), 5) + "]" + TEXT_END + " " + TEXT_BOLD + color + strings.ToUpper(level) + TEXT_END + " " + TEXT_PURPLE + logger.machine.state + TEXT_END + " " + text)
  logger.counter += 1
}

func (logger *Logger) _timeDifference() string {
  return pad(strconv.Itoa(int(time.Since(logger.startTime).Minutes())), 2) + ":" + pad(strconv.Itoa(int(time.Since(logger.startTime).Seconds()) % 60), 2)
}

func pad(str string, plength int) string {
  for i := len(str); i < plength; i++ {
    str = "0" + str
  }
  return str
}