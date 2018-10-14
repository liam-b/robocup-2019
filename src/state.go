package main

import "strings"

type StateMachineTemplate struct {
  States []string
  Links map[string][]string
}

type StateMachine struct {
  state string
  states []string
  links map[string][]string
  nextLinkSources []string
  nextLinkDestination string
}

func (stateMachine StateMachine) new() StateMachine {
  stateMachine.links = make(map[string][]string)
  return stateMachine
}

func (stateMachine *StateMachine) add(state string) {
  stateMachine.states = append(stateMachine.states, state)
  if stateMachine.state == "" {
    stateMachine.state = state
  }
}

func (stateMachine *StateMachine) include(template StateMachineTemplate) {
  stateMachine.states = append(stateMachine.states, template.States...)
  for destination, sources := range template.Links {
    stateMachine.links[destination] = append(stateMachine.links[destination], sources...)
  }
}

func (stateMachine *StateMachine) link(links ...string) *StateMachine {
  stateMachine.nextLinkDestination = links[0]
  stateMachine.nextLinkSources = links
  return stateMachine
}

func (stateMachine *StateMachine) to(destination string) {
  if contains(stateMachine.states, destination) && len(stateMachine.nextLinkSources) != 0 {
    stateMachine.links[destination] = append(stateMachine.links[destination], stateMachine.nextLinkSources...)
  }
  stateMachine.nextLinkDestination = ""
  stateMachine.nextLinkSources = stateMachine.nextLinkSources[:0]
}

func (stateMachine *StateMachine) from(sources ...string) {
  if contains(stateMachine.states, stateMachine.nextLinkDestination) && stateMachine.nextLinkDestination != "" {
    stateMachine.links[stateMachine.nextLinkDestination] = append(stateMachine.links[stateMachine.nextLinkDestination], sources...)
  }
  stateMachine.nextLinkDestination = ""
  stateMachine.nextLinkSources = stateMachine.nextLinkSources[:0]
}

func (stateMachine *StateMachine) set(state string) {
  if contains(stateMachine.states, state) {
    stateMachine.state = state
  }
}

func (stateMachine StateMachine) is(state string) bool {
  if strings.HasSuffix(state, ".*") && strings.HasPrefix(stateMachine.state, strings.TrimSuffix(state, ".*")) {
    return true
  }
  return stateMachine.state == state
}

func (stateMachine *StateMachine) transition(destination string) bool {
  if contains(stateMachine.states, destination) && stateMachine.can(destination) {
    stateMachine.state = destination
    return true
  }
  return false
}

func (stateMachine StateMachine) can(destination string) bool {
  for _, source := range stateMachine.links[destination] {
    if strings.HasSuffix(source, ".*") && strings.HasPrefix(stateMachine.state, strings.TrimSuffix(source, ".*")) {
      return true
    }
  }

  return contains(stateMachine.links[destination], stateMachine.state)
}

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}
