package bot

import (
	"github.com/liam-b/robocup-2019/io/i2c"
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/logger"

	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	CYCLE_FREQUENCY = 50
)

var (
	looping bool = true
	nextCycleFrame bool = false
	CycleThread Thread

	Start func()
	Exit  func()
	Cycle func()

	Multiplexer       i2c.Multiplexer
	ColorSensorLeft   i2c.ColorSensor
	ColorSensorMiddle i2c.ColorSensor
	ColorSensorRight  i2c.ColorSensor
	// ColorSensorCan    i2c.ColorSensor
	ColorMultiplexer  i2c.Multiplexer
	UltrasonicSensor  i2c.UltrasonicSensor

	DriveMotorLeft    lego.Motor
	DriveMotorRight   lego.Motor
	ClawMotor         lego.Motor
	ClawElevatorMotor lego.Motor
)

func Init(_Start func(), _Exit func(), _Cycle func()) {
	Start = _Start
	Exit = _Exit
	Cycle = _Cycle
	CycleThread = Thread{Target: CYCLE_FREQUENCY, Cycle: cycleWrapper}.New()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		<-stop
		fmt.Print("\n")
		logger.Print("caught interrupt")
		Stop()
	}()

	Start()
	CycleThread.Run()
}

func Stop() {
	CycleThread.Destroy()
	Exit()
}

func Setup() {
	Multiplexer.Setup()
	ColorSensorLeft.Setup()
	ColorSensorMiddle.Setup()
	ColorSensorRight.Setup()
	// ColorSensorCan.Setup()
	UltrasonicSensor.Setup()

	DriveMotorLeft.Setup()
	DriveMotorRight.Setup()
	ClawMotor.Setup()
	ClawElevatorMotor.Setup()
}

func Update() {
	ColorSensorLeft.Update()
	ColorSensorMiddle.Update()
	ColorSensorRight.Update()
	// ColorSensorCan.Update()
	UltrasonicSensor.Update()

	DriveMotorLeft.Update()
	DriveMotorRight.Update()
	ClawMotor.Update()
	ClawElevatorMotor.Update()
}

func Cleanup() {
	Multiplexer.Cleanup()
	ColorSensorLeft.Cleanup()
	ColorSensorMiddle.Cleanup()
	ColorSensorRight.Cleanup()
	UltrasonicSensor.Cleanup()

	DriveMotorLeft.Cleanup()
	DriveMotorRight.Cleanup()
	ClawMotor.Cleanup()
	ClawElevatorMotor.Cleanup()
}

func Time(milliseconds int) int {
	return int(CYCLE_FREQUENCY * (float64(milliseconds) / 1000.0))
}

func CycleDelay() {
	for !nextCycleFrame { time.Sleep(time.Millisecond) }
	nextCycleFrame = false
	// time.Sleep(time.Duration(1.0 / float64(CYCLE_FREQUENCY) * 1000) * time.Millisecond)
}

func cycleWrapper() {
	Cycle()
	nextCycleFrame = true
}