package behaviours

import (
	"github.com/liam-b/robocup-2019/logger"
	"github.com/liam-b/robocup-2019/state_machine"
)

var folow_line = Behaviour{
	setup: func() {
		logger.Trace("setting up follow_line behaviour")
		state_machine.Add(state_machine.State{Name: "follow_line", Transitions: []string{"pause", "green_turn.verify", "water_tower.verify"}})
	},
}
