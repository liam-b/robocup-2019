package bot

import (
	"github.com/liam-b/robocup-2019/io/lego"
	"github.com/liam-b/robocup-2019/io/i2c"
	"github.com/liam-b/robocup-2019/logger"

	"strconv"
	"fmt"
	"os"
	"os/signal"
)

const (
	MAIN_CYCLE_FREQUENCY = 30
	IO_CYCLE_FREQUENCY = 10
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

	Multiplexer i2c.Multiplexer
	ColorSensorLeft i2c.ColorSensor
	ColorSensorMiddle i2c.ColorSensor
	ColorSensorRight i2c.ColorSensor
	ColorMultiplexer i2c.Multiplexer
	GyroSensor i2c.GyroSensor
	UltrasonicSensor i2c.UltrasonicSensor

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
	
	mainThread = Thread{Target: MAIN_CYCLE_FREQUENCY, Cycle: mainCycleWrapper}.New()
	ioThread = Thread{Target: IO_CYCLE_FREQUENCY, Cycle: ioCycleWrapper}.New()

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
	ioThread.Start()
	mainThread.Run()
}

func Setup() {
	logger.Trace("setting up io devices")

	Multiplexer.Setup()
	// ColorSensorLeft.Setup()
	ColorSensorMiddle.Setup()
	// ColorSensorRight.Setup()
	// GyroSensor.Setup()
	UltrasonicSensor.Setup()

	LeftDriveMotor.Setup()
	RightDriveMotor.Setup()
	ClawMotor.Setup()
	ClawElevatorMotor.Setup()
}

func Update() {
	// ColorSensorLeft.Update()
	ColorSensorMiddle.Update()
	// ColorSensorRight.Update()
	// GyroSensor.Update()
	UltrasonicSensor.Update()

	// LeftDriveMotor.Update()
	// RightDriveMotor.Update()
	// ClawMotor.Update()
	// ClawElevatorMotor.Update()
}

func Cleanup() {
	logger.Trace("cleaning up io devices")

	// Multiplexer.Cleanup()
	// ColorSensorLeft.Cleanup()
	// ColorSensorMiddle.Cleanup()
	// ColorSensorRight.Cleanup()
	// GyroSensor.Cleanup()
	// UltrasonicSensor.Cleanup()

	// LeftDriveMotor.Cleanup()
	// RightDriveMotor.Cleanup()
	ClawMotor.Cleanup()
	ClawElevatorMotor.Cleanup()
}

func Stop() {
	ioThread.Stop()
	mainThread.Stop()
	Exit()
	os.Exit(0)
}

func mainCycleWrapper(frequency float64, cycles int64) {
	droppedThreadCycles(mainThread, "main")
	MainCycle(frequency, cycles)
}

func ioCycleWrapper(frequency float64, cycles int64) {
	droppedThreadCycles(ioThread, "io")
	IOCycle(frequency, cycles)
}

func droppedThreadCycles(thread Thread, name string) {
	dropped := 100 - int(thread.frequency / thread.Target * 100)
	if dropped > int(CYCLE_DROP_THRESHOLD * 100) {
		logger.Warn(name + " thread dropping " + strconv.Itoa(dropped) + "% of cycles")
	}
}