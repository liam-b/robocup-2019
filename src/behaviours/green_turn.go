package main

var greenTurn = Behaviour{
	setup: func() {
		logger.trace("setting up green_turn behaviour")
		machine.add(State{name: "green_turn.verify", transitions: []string{"green_turn.turn_left", "water_tower.turn_right"}})

		machine.add(State{name: "green_turn.turn_left", transitions: []string{"green_turn.cooldown"}})
		machine.add(State{name: "green_turn.turn_right", transitions: []string{"green_turn.cooldown"}})

		machine.add(State{name: "green_turn.cooldown", transitions: []string{"follow_line"}})
	},
}