package main

var behaviours = []Behaviour{
  followLine,
  greenTurn,
  waterTower,
  pause,
}

type Behaviour struct {
  setup func()
}

func setupBehaviourStates()  {
  logger.debug("running setup for behaviours")

  for _, behaviour := range behaviours {
    behaviour.setup()
  }
}