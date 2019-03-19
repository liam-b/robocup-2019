package main

import (
	"github.com/liam-b/robocup-2019/behaviour"
	"github.com/liam-b/robocup-2019/helper"
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/logger"
	"github.com/liam-b/robocup-2019/state_machine"
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/io/i2c"

	"time"
	"runtime"
)

func main() {
	logger.Init(5, &state_machine.Current, &state_machine.Event)
	state_machine.Init()
	bot.Init(start, exit, loop, update)
}

func start() {
	logger.Info("started")

	logger.Debug("max goroutines:", runtime.GOMAXPROCS(0))

	logger.Debug("initialising io devices")
	bot.Multiplexer = i2c.Multiplexer{}.New()
	bot.ColorSensorMiddle = i2c.ColorSensor{Multiplexer: &bot.Multiplexer, Channel: 0}.New()
	bot.UltrasonicSensor = i2c.UltrasonicSensor{}.New()

	bot.LeftDriveMotor = lego.Motor{Port: lego.PORT_MA}.New()
	bot.RightDriveMotor = lego.Motor{Port: lego.PORT_MD}.New()
	bot.ClawMotor = lego.Motor{Port: lego.PORT_MB}.New()
	bot.ClawElevatorMotor = lego.Motor{Port: lego.PORT_MC}.New()

	bot.Setup()
	helper.Setup()
	behaviour.Setup()

	state_machine.Transition("follow_line")

	time.Sleep(time.Second)
	
	doSumCanStuff()
}

func loop(frequency float64, cycle int64) {
	logger.Debug(bot.UltrasonicSensor.Distance())
	// logger.Debug(bot.ColorSensorMiddle.Intensity())

	time.Sleep(time.Millisecond * 20)
}

func update(frequency float64, cycle int64) {
	bot.Update()
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