package main

/*

// Sensors //
- i2c color sensor x4
- i2c distance sensor x1
- i2c compass sensor x1

// Motors //
- ev3 large motor x2
- ev3 medium motor x1

*/

import (
	"github.com/liam-b/robocup-2019/behaviours"
	"github.com/liam-b/robocup-2019/bot"
	"github.com/liam-b/robocup-2019/logger"
	"github.com/liam-b/robocup-2019/state_machine"
	"github.com/liam-b/robocup-2019/io"

	"runtime"
	"strconv"
)

var multiplexer io.Multiplexer
var colorSensorLeft io.ColorSensor
var colorSensorRight io.ColorSensor
var colorSensorMiddle io.ColorSensor
var compassSensor io.CompassSensor

func main() {
	logger.Init(5, &state_machine.Current, &state_machine.Event)
	state_machine.Init()
	bot.Init(start, exit, loop, thread)
}

func start() {
	logger.Info("hi")

	logger.Debug("GOMAXPROCS: " + strconv.Itoa(runtime.GOMAXPROCS(0)))

	multiplexer = io.Multiplexer{}.New()
	// multiplexer.Setup()

	colorSensorLeft = io.ColorSensor{Multiplexer: &multiplexer, Channel: 0}.New()
	// colorSensorLeft.Setup()

	colorSensorRight = io.ColorSensor{Multiplexer: &multiplexer, Channel: 1}.New()
	// colorSensorRight.Setup()

	colorSensorMiddle = io.ColorSensor{Multiplexer: &multiplexer, Channel: 3}.New()
	// colorSensorMiddle.Setup()

	compassSensor = io.CompassSensor{}.New()
	compassSensor.Setup()

	behaviours.Start()
}

func loop(frequency float64, cycle int64) {
	// logger.Debug("bit of a loop")

	// fmt.Println(sensor1.Intensity(), sensor2.Intensity(), sensor3.Intensity())

	// val1, val2 := colorSensor.Intensity()

	// fmt.Println(strconv.FormatInt(int64(val1), 2), strconv.FormatInt(int64(val2), 2))
	// bot.Stop()

	logger.Trace(compassSensor.Rotation())
}

func thread(frequency float64, cycle int64) {
	// fmt.Println("hi")
}

func exit() {
	logger.Info("exiting")
}
