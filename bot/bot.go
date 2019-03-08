package bot

import "github.com/liam-b/robocup-2019/io/i2c"

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
	Multiplexer i2c.Multiplexer
	GyroSensor i2c.GyroSensor
)

func Init(_Start func(), _Exit func(), _MainCycle func(float64, int64), _IOCycle func(float64, int64)) {
	Start = _Start
	Exit = _Exit

	MainCycle = _MainCycle
	IOCycle = _IOCycle
	
	mainThread = Thread{Target: MAIN_CYCLE_FREQUENCY, Cycle: MainCycle}.New()
	ioThread = Thread{Target: IO_CYCLE_FREQUENCY, Cycle: IOCycle}.New()

	Start()
	mainThread.Run()
	ioThread.Start()
}

func Stop() {
	mainThread.Stop()
	mainThread.Start()
	Exit()
}

func UpdateSensorCaches() {
	ColorSensorLeft.Update()
	ColorSensorMiddle.Update()
	ColorSensorRight.Update()
	GyroSensor.Update()
}
