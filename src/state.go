package main

type StateEvent string; const (
  EnterEvent StateEvent = ">"
  UpdateEvent StateEvent = "-"
  ExitEvent StateEvent = "<"

  OutsideEvent StateEvent = "."
  InternalEvent StateEvent = "!"
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
    machine.event = OutsideEvent
  } else if contains(machine.states[machine.state].transitions, destination) {
    machine.event = ExitEvent
    machine._log("exited")
    machine.states[machine.state].exit()

    machine.state = destination

    machine.event = EnterEvent
    machine._log("entered")
    machine.states[machine.state].enter()
    machine.event = OutsideEvent
  }
}

func (machine *StateMachine) update() {
  machine.event = UpdateEvent
  machine._log("updated")
  machine.states[machine.state].update()
  machine.event = OutsideEvent
}

func (machine *StateMachine) add(state State) {
  machine.states[state.name] = state
  if machine.state == "" {
    machine.state = state.name
    machine.event = EnterEvent
    machine._log("entered")
    machine.states[machine.state].enter()
    machine.event = OutsideEvent
  }
}

func (machine *StateMachine) _log(text string) {
  logger.state(text + " " + machine.state)
}
