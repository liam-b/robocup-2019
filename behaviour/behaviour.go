package behaviour

import "github.com/liam-b/robocup-2019/logger"

var behaviours = []Behaviour{
}

type Behaviour struct {
	Setup   func()
	Cleanup func()
}

func Tester() {
	logger.Trace("running setup for behaviours")
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
