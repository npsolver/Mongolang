package dfa

import (
	"errors"

	"github.com/npsolver/Mongolang/global"
)

type Traveler struct {
	dfa       *DFA
	chars     []Char
	n         int
	nextIndex int
}

func (tv *Traveler) NextSymbol() (*global.Symbol, error) {
	if tv.nextIndex == tv.n {
		return nil, nil
	}
	currState := tv.dfa.startingState
	foundAcceptingSymbol := false
	var lastAcceptingState *State
	var currValue Char
	var lastAcceptingChar Char
	indexAfterLastAcceptingSymbol := -1
	for tv.nextIndex < tv.n {
		currState = currState.nextState(tv.chars[tv.nextIndex])
		currValue = currValue.append(tv.chars[tv.nextIndex])
		if currState == nil {
			break
		}
		if currState.isAccepting {
			foundAcceptingSymbol = true
			lastAcceptingState = currState
			lastAcceptingChar = currValue
			indexAfterLastAcceptingSymbol = tv.nextIndex + 1
		}
		tv.nextIndex += 1
	}
	if foundAcceptingSymbol {
		tv.nextIndex = indexAfterLastAcceptingSymbol
		return global.NewTerminatingSymbol(lastAcceptingState.name, string(lastAcceptingChar)), nil
	}
	return nil, errors.New("could not produce symbol")
}
