package state_machine

import (
	"github.com/liam-b/robocup-2019/logger"
)

type State struct {
	Name        string
	Transitions []string
	Enter       func()
	Update      func()
	Exit        func()
}

type stateEvent string
const (
	enterEvent    stateEvent = ">"
	updateEvent   stateEvent = "-"
	exitEvent     stateEvent = "<"
	outsideEvent  stateEvent = "."
	internalEvent stateEvent = "!"
)

var (
	Current string
	States  map[string]State = make(map[string]State)
	Event   string
)

func Init() {}

func Transition(destination string) {
	if _, exists := States[destination]; !exists {
		setEvent(internalEvent)
		logger.Warn("attempted transition to unknown state")
		setEvent(outsideEvent)
	} else if contains(States[Current].Transitions, destination) {
		setEvent(exitEvent)
		logTransition("exited")
		if (States[Current].Exit != nil) {
			States[Current].Exit()
		}

		Current = destination

		setEvent(enterEvent)
		logTransition("entered")
		if (States[Current].Enter != nil) {
			States[Current].Enter()
		}
		setEvent(outsideEvent)
	}
}

func Update() {
	setEvent(updateEvent)
	logTransition("updated")
	if (States[Current].Update != nil) {
		States[Current].Update()
	}
	setEvent(outsideEvent)
}

func Add(state State) {
	States[state.Name] = state
	if Current == "" {
		Current = state.Name
		setEvent(enterEvent)
		logTransition("entered")
		if (States[Current].Enter != nil) {
			States[Current].Enter()
		}
		setEvent(outsideEvent)
	}
}

func logTransition(text string) {
	logger.State(text + " " + Current)
}

func setEvent(event stateEvent) {
	Event = (string)(event)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
