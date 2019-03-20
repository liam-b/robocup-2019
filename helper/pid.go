package helper

import (
  "github.com/liam-b/robocup-2019/bot"
)

const (
  PROPORTIONAL = 15.0
  INTEGRAL = 1.0
  DERIVATIVE = 5.0

  // INTEGRAL_RANGE = 

  BASE_SPEED = 150
)

var (
  lastError = 0.0
  integral = 0.0
)

func PID() (int, int) {
  currentError := LineError()
  integral += currentError * (1.0 / float64(bot.MAIN_CYCLE_FREQUENCY))
  derivative := (currentError - lastError) * float64(bot.MAIN_CYCLE_FREQUENCY) / 4

  speed := (PROPORTIONAL * currentError) + (INTEGRAL * integral) + (DERIVATIVE * derivative)

  lastError = currentError;

  left := min(max(BASE_SPEED + int(speed), -1000), 1000)
  right := min(max(BASE_SPEED - int(speed), -1000), 1000)

  return left, right
}

func LineError() float64 {
  return float64(bot.ColorSensorLeft.Intensity() - bot.ColorSensorRight.Intensity())
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