package behaviour

import "github.com/liam-b/robocup-2019/logger"

var behaviours = []Behaviour{
	followLine,
	greenTurn,
	// waterTower,
	// pause,
}

type Behaviour struct {
	Setup   func()
	Cleanup func()
}

func Setup() {
	logger.Trace("running setup for behaviours")

	for _, behaviour := range behaviours {
		if behaviour.Setup != nil {
			behaviour.Setup()
		}
	}
}

func Cleanup() {
	logger.Trace("running cleanup for behaviours")

	for _, behaviour := range behaviours {
		if behaviour.Cleanup != nil {
			behaviour.Cleanup()
		}
	}
}
