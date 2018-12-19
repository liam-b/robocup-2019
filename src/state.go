package main

type StateEvent string; const (
  EnterEvent StateEvent = ">"
  UpdateEvent StateEvent = "-"
  ExitEvent StateEvent = "<"

  InternalEvent StateEvent = ":"
  ExternalEvent StateEvent = "."
)

type State struct {
  name string
  transitions []string
  enter func()
  update func()
  exit func()
}

type StateMachine struct {
  logger *Logger
  state string
  states map[string]State
  event StateEvent
}

func (machine StateMachine) new() StateMachine {
  machine.states = make(map[string]State)
  return machine
}

func (machine *StateMachine) transition(destination string) {
  if _, exists := machine.states[destination]; !exists {
    machine.event = InternalEvent
    logger.warn("attempted transition to unknown state")
    machine.event = ExternalEvent
  } else {
    if contains(machine.states[machine.state].transitions, destination) {
      machine._callExit()
      machine.state = destination
      machine._callEnter()
    } else {
      machine.event = InternalEvent
      logger.state("attempted illegal transition")
      machine.event = ExternalEvent
    }
  }
}

func (machine *StateMachine) update() {
  machine._callUpdate()
}

func (machine *StateMachine) add(state State) {
  logger.state("adding state " + state.name)
  machine.states[state.name] = state
  if machine.state == "" {
    logger.state("setting intital state to " + state.name)
    machine.state = state.name
    machine._callEnter()
  }
}

func (machine *StateMachine) _logTransition(text string) {
  logger.state(text + " " + machine.state)
}

func (machine *StateMachine) _callEnter() {
  machine.event = EnterEvent
  machine._logTransition("entered")
  if machine.states[machine.state].enter != nil {
    machine.states[machine.state].enter()
  }
  machine.event = ExternalEvent
}

func (machine *StateMachine) _callUpdate() {
  machine.event = UpdateEvent
  machine._logTransition("updated")
  if machine.states[machine.state].update != nil {
    machine.states[machine.state].update()
  }
  machine.event = ExternalEvent
}

func (machine *StateMachine) _callExit() {
  machine.event = ExitEvent
  machine._logTransition("exited")
  if machine.states[machine.state].exit != nil {
    machine.states[machine.state].exit()
  }
  machine.event = ExternalEvent
}