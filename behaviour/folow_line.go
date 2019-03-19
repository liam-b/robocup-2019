package behaviour

import (
	"github.com/liam-b/robocup-2019/state_machine"
)

var followLine = Behaviour{
	Setup: func() {
		state_machine.Add(state_machine.State{
			Name: "follow_line",
			Transitions: []string{"pause", "green_turn.verify", "water_tower.verify"},
			Enter: func() {
				
			},
			Update: func() {
				
			},
			Exit: func() {
				
			},
		})
	},
}
