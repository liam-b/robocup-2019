package bot

import (
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/io/i2c"
	"github.com/liam-b/robocup-2019/logger"

	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	CYCLE_FREQUENCY = 60
)

var (
	looping bool = true
	CycleThread Thread

	Start func()
	Loop func()
	Exit func()

	Cycle func()

	Multiplexer i2c.Multiplexer
	ColorSensorLeft i2c.ColorSensor
	ColorSensorMiddle i2c.ColorSensor
	ColorSensorRight i2c.ColorSensor
	ColorMultiplexer i2c.Multiplexer
	GyroSensor i2c.GyroSensor
	UltrasonicSensor i2c.UltrasonicSensor

	DriveMotorLeft lego.Motor
	DriveMotorRight lego.Motor
	ClawMotor lego.Motor
	ClawElevatorMotor lego.Motor
)

func Init(_Start func(), _Exit func(), _Cycle func()) {
	Start = _Start
	Exit = _Exit
	Cycle = _Cycle
	CycleThread = Thread{Target: CYCLE_FREQUENCY, Cycle: Cycle}.New()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		for _ = range stop {
			fmt.Print("\n")
			logger.Debug("caught interrupt")
			Stop()
		}
	}()

	Start()
	CycleThread.Run()
}

func Stop() {
	CycleThread.Stop()
	time.Sleep(time.Millisecond * 500)
	Exit()
	time.Sleep(time.Millisecond * 500)
	CycleThread.Destroy()
	os.Exit(0)
}

func Setup() {
	logger.Trace("setting up io devices")

	Multiplexer.Setup()
	ColorSensorLeft.Setup()
	ColorSensorMiddle.Setup()
	ColorSensorRight.Setup()
	// GyroSensor.Setup()
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
	// GyroSensor.Update()
	UltrasonicSensor.Update()

	DriveMotorLeft.Update()
	DriveMotorRight.Update()
	ClawMotor.Update()
	ClawElevatorMotor.Update()
}

func Cleanup() {
	logger.Trace("cleaning up io devices")

	Multiplexer.Cleanup()
	ColorSensorLeft.Cleanup()
	ColorSensorMiddle.Cleanup()
	ColorSensorRight.Cleanup()
	// GyroSensor.Cleanup()
	UltrasonicSensor.Cleanup()

	DriveMotorLeft.Cleanup()
	DriveMotorRight.Cleanup()
	ClawMotor.Cleanup()
	ClawElevatorMotor.Cleanup()
}