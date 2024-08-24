package dfa

import (
	"bufio"
	"log"
	"os"
)

const (
	STATES      = ".STATES"
	TRANSITIONS = ".TRANSITIONS"
	END         = ".END"
	exclamation = "!"
)

type DFA struct {
	allSymbols    []Symbol
	allStates     map[StateName]*State
	startingState *State
}

func (d *DFA) NewDFATraveler(s string) *Traveler {
	symbols := []Symbol{}
	for _, c := range s {
		symbols = append(symbols, Symbol(c))
	}
	return &Traveler{
		dfa:       d,
		symbols:   symbols,
		n:         len(symbols),
		nextIndex: 0,
	}
}

func NewDFA(filePath string) *DFA {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	dfa := &DFA{}
	dfa.scanSymbols(fileScanner)
	dfa.scanStates(fileScanner)
	dfa.scanTransitions(fileScanner)

	return dfa
}
