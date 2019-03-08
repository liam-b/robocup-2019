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

	"runtime"
	"strconv"
)

func main() {
	logger.Init(5, &state_machine.Current, &state_machine.Event)
	state_machine.Init()
	bot.Init(start, exit, loop, update)
}

func start() {
	logger.Info("started")

	logger.Debug("GOMAXPROCS: " + strconv.Itoa(runtime.GOMAXPROCS(0)))

	behaviours.Start()
}

func loop(frequency float64, cycle int64) {
}

func update(frequency float64, cycle int64) {
	bot.UpdateSensorCaches()
}

func exit() {
	logger.Info("exiting")
}
