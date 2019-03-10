package bot

import (
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/io/i2c"

	"fmt"
	"os"
	"os/signal"
)

const (
	MAIN_CYCLE_FREQUENCY = 30
	IO_CYCLE_FREQUENCY = 1
)

var (
	looping bool = true
	mainThread Thread
	ioThread Thread

	Start func()
	Loop func()
	Exit func()

	MainCycle func(float64, int64)
	IOCycle func(float64, int64)

	ColorSensorLeft i2c.ColorSensor
	ColorSensorMiddle i2c.ColorSensor
	ColorSensorRight i2c.ColorSensor
	ColorMultiplexer i2c.Multiplexer
	GyroSensor i2c.GyroSensor

	LeftDriveMotor lego.Motor
	RightDriveMotor lego.Motor
	ClawMotor lego.Motor
	ClawElevatorMotor lego.Motor
)

func Init(_Start func(), _Exit func(), _MainCycle func(float64, int64), _IOCycle func(float64, int64)) {
	Start = _Start
	Exit = _Exit

	MainCycle = _MainCycle
	IOCycle = _IOCycle
	
	mainThread = Thread{Target: MAIN_CYCLE_FREQUENCY, Cycle: MainCycle}.New()
	ioThread = Thread{Target: IO_CYCLE_FREQUENCY, Cycle: IOCycle}.New()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
			for _ = range stop {
				fmt.Print("\n")
				Stop()
			}
	}()

	Start()
	mainThread.Run()
	ioThread.Start()
}

func Setup() {
	// ColorSensorLeft.Setup()
	// ColorSensorMiddle.Setup()
	// ColorSensorRight.Setup()
	// GyroSensor.Setup()

	LeftDriveMotor.Setup()
	RightDriveMotor.Setup()
	ClawMotor.Setup()
	ClawElevatorMotor.Setup()
}

func UpdateCaches() {
	// ColorSensorLeft.Update()
	// ColorSensorMiddle.Update()
	// ColorSensorRight.Update()
	// GyroSensor.Update()

	// LeftDriveMotor.Update()
	// RightDriveMotor.Update()
	// ClawMotor.Update()
	// ClawElevatorMotor.Update()
}

func Cleanup() {
	// ColorSensorLeft.Cleanup()
	// ColorSensorMiddle.Cleanup()
	// ColorSensorRight.Cleanup()
	// GyroSensor.Cleanup()

	// LeftDriveMotor.Cleanup()
	// RightDriveMotor.Cleanup()
	ClawMotor.Cleanup()
	ClawElevatorMotor.Cleanup()
}

func Stop() {
	mainThread.Stop()
	mainThread.Start()
	Exit()
	os.Exit(0)
}