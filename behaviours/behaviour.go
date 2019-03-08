package behaviours

import "github.com/liam-b/robocup-2019/logger"

var behaviours = []Behaviour{
	followLine,
	// GreenTurn,
	// WaterTower,
	// pause,
}

type Behaviour struct {
	setup func()
}

func Start() {
	logger.Trace("running setup for behaviours")

	for _, behaviour := range behaviours {
		behaviour.setup()
	}
}
