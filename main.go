package main

import (
	"github.com/liam-b/robocup-2019/behaviour"
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/io/i2c"
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/logger"
	"os"
	"runtime"
	"fmt"
)

var file *os.File

func main() {
	bot.Init(Start, Exit, Loop)
}

func Start() {
	logger.Print("started")
	logger.Print("max goroutines:", runtime.GOMAXPROCS(0))

	logger.Print("initialising io")
	bot.Multiplexer = i2c.Multiplexer{}.New()
	bot.ColorSensorLeft = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 4}.New()
	bot.ColorSensorMiddle = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 3}.New()
	bot.ColorSensorRight = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 0}.New()
	bot.UltrasonicSensor = i2c.UltrasonicSensor{}.New()

	bot.DriveMotorLeft = lego.Motor{Port: lego.PORT_MA}.New()
	bot.DriveMotorRight = lego.Motor{Port: lego.PORT_MD}.New()
	bot.ClawMotor = lego.Motor{Port: lego.PORT_MB}.New()
	bot.ClawElevatorMotor = lego.Motor{Port: lego.PORT_MC}.New()

	bot.Setup()
	bot.ClawMotor.ResetPosition()
	bot.ClawElevatorMotor.ResetPosition()

	// bot.ClawMotor.RunToAbsolutePositionAndBrake(-170, 300)

	logger.Print("waiting for start...")
	var input string
	fmt.Scanln(&input)

	logger.Print("starting behaviour")
	go behaviour.FollowLine()
}

func Loop() {
	bot.Update()
	// logger.Print(bot.UltrasonicSensor.Distance())
	// logger.Print(bot.ColorSensorLeft.Intensity(), bot.ColorSensorRight.Intensity())
	// logger.Print(bot.ColorSensorMiddle.Intensity())
}

func Exit() {
	logger.Print("exiting")

	helper.OpenClaw()
	helper.LowerClaw()
	bot.DriveMotorLeft.Coast()
	bot.DriveMotorRight.Coast()

	bot.Cleanup()
}
