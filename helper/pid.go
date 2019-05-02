package helper

import (
  // "github.com/liam-b/robocup-2019/bot"
  // "github.com/liam-b/robocup-2019/logger"
)

const (
  PROPORTIONAL = 400.0
  INTEGRAL = 1.5
  DERIVATIVE = 3000.0

  BASE_SPEED = 300
)

var (
  lastError = 0.0
  integral = 0.0
)

func PID() (int, int) {
  currentError := LineError()
  integral += currentError
  derivative := currentError - lastError
  

  // logger.Debug("err:", int(currentError), "P:", int(PROPORTIONAL * proportional), "I:", int(INTEGRAL * integral), "D:", int(DERIVATIVE * derivative))

  speed := (PROPORTIONAL * currentError) + (INTEGRAL * integral) + (DERIVATIVE * derivative)
  lastError = currentError

  left := min(max(-1000, BASE_SPEED + int(speed)), 1000)
  right := min(max(-1000, BASE_SPEED - int(speed)), 1000)

  // logger.Debug(delta, bot.IOThread.LastCycleTime, bot.MainThread.LastCycleTime)
  // logger.Debug("c", currentError)
  // logger.Debug("d", derivative)
  // logger.Debug("D", DERIVATIVE * derivative)
  // logger.Debug("s", speed)

  return left, right
}

func ResetPID() {
  lastError = 0.0
  integral = 0.0
}

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}

func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}

func cap(v int, r int) int {
  if v > r {
    return r
  }
  if v < -r {
    return -r
  }
  return v
}

func maxf(a float64, b float64) float64 {
  if a > b {
    return a
  }
  return b
}

func minf(a float64, b float64) float64 {
  if a < b {
    return a
  }
  return b
}

func capf(v float64, r float64) float64 {
  if v > r {
    return r
  }
  if v < -r {
    return -r
  }
  return v
}