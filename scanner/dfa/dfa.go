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
	allChars      []Char
	allStates     map[StateName]*State
	startingState *State
}

func (d *DFA) NewDFATraveler(s string) *Traveler {
	chars := []Char{}
	for _, c := range s {
		chars = append(chars, Char(c))
	}
	return &Traveler{
		dfa:       d,
		chars:     chars,
		n:         len(chars),
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
	dfa.scanChars(fileScanner)
	dfa.scanStates(fileScanner)
	dfa.scanTransitions(fileScanner)

	return dfa
}

func (d *DFA) GetTerminatingStates() []string {
	states := []string{}
	for _, v := range d.allStates {
		if v.isAccepting {
			states = append(states, v.name)
		}
	}
	return states
}
