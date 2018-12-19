package main

var waterTower = Behaviour{
	setup: func() {
		logger.trace("setting up water_tower behaviour")
		machine.add(State{name: "water_tower.verify", transitions: []string{"water_tower.turn", "water_tower.reverse"}})
		machine.add(State{name: "water_tower.turn", transitions: []string{"water_tower.skirt"}})
		machine.add(State{name: "water_tower.skirt", transitions: []string{"water_tower.recapture"}})
		machine.add(State{name: "water_tower.recapture", transitions: []string{"follow_line"}})

		machine.add(State{name: "water_tower.reverse", transitions: []string{"follow_line"}})
	},
}