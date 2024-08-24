package dfa

import (
	"errors"
)

type Traveler struct {
	dfa       *DFA
	symbols   []Symbol
	n         int
	nextIndex int
}

func (tv *Traveler) NextToken() (*Token, error) {
	if tv.nextIndex == tv.n {
		return nil, nil
	}
	currState := tv.dfa.startingState
	foundAcceptingToken := false
	var lastAcceptingState *State
	var currValue Symbol
	var lastAcceptingSymbol Symbol
	indexAfterLastAcceptingToken := -1
	for tv.nextIndex < tv.n {
		currState = currState.nextState(tv.symbols[tv.nextIndex])
		currValue = currValue.append(tv.symbols[tv.nextIndex])
		if currState == nil {
			break
		}
		if currState.isAccepting {
			foundAcceptingToken = true
			lastAcceptingState = currState
			lastAcceptingSymbol = currValue
			indexAfterLastAcceptingToken = tv.nextIndex + 1
		}
		tv.nextIndex += 1
	}
	if foundAcceptingToken {
		tv.nextIndex = indexAfterLastAcceptingToken
		return NewToken(lastAcceptingState.name, string(lastAcceptingSymbol)), nil
	}
	return nil, errors.New("could not produce token")
}
