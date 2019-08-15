package main

import (
	"github.com/liam-b/robocup-2019/behaviour"
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/io/i2c"
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/logger"
	"github.com/liam-b/robocup-2019/state_machine"

	"os"
	"runtime"
	"time"
)

var file *os.File

func main() {
	logger.Init(5, &state_machine.Current, &state_machine.Event)
	state_machine.Init()
	bot.Init(start, exit, loop)
}

func start() {
	logger.Info("started")
	logger.Debug("max goroutines:", runtime.GOMAXPROCS(0))

	logger.Debug("initialising io devices")
	bot.Multiplexer = i2c.Multiplexer{}.New()
	bot.ColorSensorLeft = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 4}.New()
	bot.ColorSensorMiddle = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 3}.New()
	bot.ColorSensorRight = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 0}.New()
	bot.GyroSensor = i2c.GyroSensor{}.New()
	bot.UltrasonicSensor = i2c.UltrasonicSensor{}.New()

	bot.DriveMotorLeft = lego.Motor{Port: lego.PORT_MA}.New()
	bot.DriveMotorRight = lego.Motor{Port: lego.PORT_MD}.New()
	bot.ClawMotor = lego.Motor{Port: lego.PORT_MB}.New()
	bot.ClawElevatorMotor = lego.Motor{Port: lego.PORT_MC}.New()

	bot.Setup()
	helper.Setup()
	behaviour.Setup()

	bot.ClawMotor.RunToAbsolutePositionAndBrake(-170, 300)

	state_machine.Transition("follow_line.follow") 
	time.Sleep(time.Second)
}

func loop() {
	bot.Update()
	state_machine.Update()

	// logger.Debug(bot.DriveMotorLeft.Port, bot.DriveMotorLeft.Speed())

	logger.Debug(bot.UltrasonicSensor.Distance())

	// logger.Debug(helper.LeftColor(), helper.MiddleColor(), helper.RightColor())
	
	// logger.Debug(bot.UltrasonicSensor.Distance())
	// logger.Debug(bot.ColorSensorLeft.Intensity(), bot.ColorSensorMiddle.Intensity(), bot.ColorSensorRight.Intensity())

	// left, right := helper.PID()
	// bot.DriveMotorLeft.Run(left)
	// bot.DriveMotorRight.Run(right)

	// logger.Debug(helper.LineError())
	// logger.Debug(bot.ColorSensorLeft.Intensity())

	// time.Sleep(time.Millisecond * 20)
}

func exit() {
	logger.Info("exiting")

	behaviour.Cleanup()
	helper.Cleanup()
	bot.Cleanup()
}
