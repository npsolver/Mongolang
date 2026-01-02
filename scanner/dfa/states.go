package dfa

import (
	"bufio"
	"strings"
)

type StateName string

type State struct {
	name        string
	bridges     map[Char]*State
	isAccepting bool
}

func (d *DFA) scanStates(scanner *bufio.Scanner) {
	d.allStates = make(map[StateName]*State)
	scanner.Scan()
	scanner.Scan()
	startingStateName := scanner.Text()
	d.startingState = processState(startingStateName)
	d.allStates[StateName(startingStateName)] = d.startingState
	scanner.Scan()
	for scanner.Text() != TRANSITIONS {
		name := scanner.Text()
		state := processState(name)
		d.allStates[StateName(state.name)] = state
		scanner.Scan()
	}
}

func processState(name string) *State {
	n := len(name)
	m := make(map[Char]*State)
	currState := State{bridges: m}
	if name[n-1:n] == exclamation {
		currState.name = strings.Split(strings.Trim(name, " "), " ")[0] // Todo: remove other whitespaces like tabs
		currState.isAccepting = true
	} else {
		currState.name = name
		currState.isAccepting = false
	}
	return &currState
}

func (d *DFA) getState(name StateName) *State {
	return d.allStates[name]
}

func (curr *State) addBridge(sm Char, next *State) {
	curr.bridges[sm] = next
}

func (curr *State) nextState(sm Char) *State {
	elem, ok := curr.bridges[sm]
	if !ok {
		return nil
	}
	return elem
}
