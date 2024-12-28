package dfa

type DFA[T any, S comparable] struct {
	symbols       []*Symbol[S]
	states        map[StateName]*State[T, S]
	startingState *State[T, S]
}

func NewEmptyDFA[T any, S comparable]() *DFA[T, S] {
	m := make(map[StateName]*State[T, S])
	dfa := &DFA[T, S]{states: m}
	return dfa
}
