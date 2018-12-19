package main

import "fmt"
import "strconv"
import "strings"
import "time"

const TEXT_HIDE = "\x1b[8m"
const TEXT_SHOW = "\x1b[28m"
const TEXT_PURPLE = "\x1b[35m"
const TEXT_BOLD = "\x1b[1m"
const TEXT_UNDERLINE = "\x1b[4m"
const TEXT_ITALIC = "\x1b[3m"
const TEXT_BLUE = "\x1b[34m"
const TEXT_CYAN = "\x1b[36m"
const TEXT_GREEN = "\x1b[32m"
const TEXT_YELLOW = "\x1b[33m"
const TEXT_RED = "\x1b[31m"
const TEXT_BLACK = "\x1b[30m"
const TEXT_END = "\x1b[0m"
const TEXT_WHITE = ""

type LogLevel string; const (
  StateLevel LogLevel = "state"
  TraceLevel LogLevel = "trace"
  DebugLevel LogLevel = "debug"
  InfoLevel LogLevel = " info"
  WarningLevel LogLevel = " warn"
  ErrorLevel LogLevel = "error"
)

type Logger struct {
  machine *StateMachine
  startTime time.Time
  counter int
  level LogLevel
  levelInt int
}

func (logger Logger) new() Logger {
  fmt.Println("  ____   ___  ____   ___   ____ _   _ ____    ____   ___  _  ___  \n |  _ \\ / _ \\| __ ) / _ \\ / ___| | | |  _ \\  |___ \\ / _ \\/ |/ _ \\ \n | |_) | | | |  _ \\| | | | |   | | | | |_) |   __) | | | | | (_) |\n |  _ <| |_| | |_) | |_| | |___| |_| |  __/   / __/| |_| | |\\__, |\n |_| \\_\\\\___/|____/ \\___/ \\____|\\___/|_|     |_____|\\___/|_|  /_/" + "\n")
  logger._setLevelInt()
  logger.startTime = time.Now()
  logger.counter = 0
  return logger
}

func (logger *Logger) state(text string) {
  if logger.levelInt >= 5 {
    logger._print(StateLevel, TEXT_BOLD + TEXT_CYAN, text)
  }
}

func (logger *Logger) trace(text string) {
  if logger.levelInt >= 4 {
    logger._print(TraceLevel, TEXT_BOLD + TEXT_WHITE, text)
  }
}

func (logger *Logger) debug(text string) {
  if logger.levelInt >= 3 {
    logger._print(DebugLevel, TEXT_BOLD + TEXT_GREEN, text)
  }
}

func (logger *Logger) info(text string) {
  if logger.levelInt >= 2 {
    logger._print(InfoLevel, TEXT_BOLD + TEXT_BLUE, text)
  }
}

func (logger *Logger) warn(text string) {
  if logger.levelInt >= 1 {
    logger._print(WarningLevel, TEXT_BOLD + TEXT_YELLOW, text)
  }
}

func (logger *Logger) error(text string) {
  if logger.levelInt >= 0 {
    logger._print(ErrorLevel, TEXT_BOLD + TEXT_RED, text)
  }
}

// func (logger *Logger) red(text string) string {
//   return TEXT_UNDERLINE + text + TEXT_END
// }

func (logger *Logger) _print(level LogLevel, color string, text string) {
  fmt.Println(TEXT_BLACK + logger._timeDifference() + " [" + pad(strconv.Itoa(logger.counter), 5) + "]" + TEXT_END + " " + color + strings.ToUpper((string)(level)) + TEXT_END + " " + TEXT_PURPLE + logger._state() + TEXT_END + " " + text)
  logger.counter += 1
}

func (logger *Logger) _timeDifference() string {
  return pad(strconv.Itoa(int(time.Since(logger.startTime).Minutes())), 2) + ":" + pad(strconv.Itoa(int(time.Since(logger.startTime).Seconds()) % 60), 2)
}

func (logger Logger) _state() string {
  if machine.state != "" {
    return TEXT_CYAN + "[" + (string)(machine.event) + "]" + TEXT_PURPLE + machine.state
  } else {
    return TEXT_CYAN + "[.]" + TEXT_PURPLE + "{nil}"
  }
}

func (logger *Logger) _setLevelInt() {
  switch logger.level {
    case "":
    case ErrorLevel:
      logger.levelInt = 0
    case WarningLevel:
      logger.levelInt = 1
    case InfoLevel:
      logger.levelInt = 2
    case DebugLevel:
      logger.levelInt = 3
    case TraceLevel:
      logger.levelInt = 4
    default:
    case StateLevel:
      logger.levelInt = 5
  }
}

func pad(str string, plength int) string {
  for i := len(str); i < plength; i++ {
    str = "0" + str
  }
  return str
}