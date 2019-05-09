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
	bot.ColorSensorLeft = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 0}.New()
	bot.ColorSensorMiddle = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 4}.New()
	bot.ColorSensorRight = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 3}.New()
	bot.UltrasonicSensor = i2c.UltrasonicSensor{}.New()

	bot.DriveMotorLeft = lego.Motor{Port: lego.PORT_MA}.New()
	bot.DriveMotorRight = lego.Motor{Port: lego.PORT_MD}.New()
	bot.ClawMotor = lego.Motor{Port: lego.PORT_MB}.New()
	bot.ClawElevatorMotor = lego.Motor{Port: lego.PORT_MC}.New()

	bot.Setup()
	helper.Setup()
	behaviour.Setup()

	// state_machine.Transition("follow_line")

	time.Sleep(time.Second)

	// file, _ = os.Create("./graph.csv")

	// helper.RunDrive(20)

	// helper.CloseClaw()

	// doSumCanStuff()
}

// func loop(frequency float64, cycle int64) {
func loop() {
	bot.Update()
	// logger.Debug(bot.UltrasonicSensor.Distance())
	// logger.Debug(bot.ColorSensorLeft.Intensity(), bot.ColorSensorMiddle.Intensity(), bot.ColorSensorRight.Intensity())

	// helper.PID()
	left, right := helper.PID()
	bot.DriveMotorLeft.Run(left)
	bot.DriveMotorRight.Run(right)

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

func doSumCanStuff() {
	helper.CloseClaw()
	time.Sleep(time.Second)
	helper.RaiseClaw()
	time.Sleep(time.Second * 2)
	helper.RunToPositionDrive(300, 300)
	time.Sleep(time.Second * 2)
	helper.ReleaseClaw()
	time.Sleep(time.Second)
	helper.OpenClaw()
	time.Sleep(time.Second * 1)
	helper.RunToPositionDrive(0, 300)
	time.Sleep(time.Second * 1)
	helper.LowerClaw()
	time.Sleep(time.Second * 3)

	helper.RaiseClaw()
	time.Sleep(time.Second * 2)
	helper.RunToPositionDrive(300, 300)
	time.Sleep(time.Second * 2)
	helper.CloseClaw()
	time.Sleep(time.Second)
	helper.RunToPositionDrive(0, 300)
	time.Sleep(time.Second * 2)
	helper.LowerClaw()
	time.Sleep(time.Second * 2)
	helper.OpenClaw()
}
