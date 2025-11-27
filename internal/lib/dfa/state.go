package dfa

type StateName string

type State[T any, S comparable] struct {
	name        StateName
	value       T
	bridges     map[Symbol[S]]*State[T, S]
	isAccepting bool
}

func (curr *State[T, S]) NextState(sym Symbol[S]) *State[T, S] {
	elem, ok := curr.bridges[sym]
	if !ok {
		return nil
	}
	return elem
}

func (d *DFA[T, S]) getState(name StateName) *State[T, S] {
	return d.states[name]
}

func (curr *State[T, S]) addBridge(sm S, next *State[T, S]) {
	curr.bridges[Symbol[S]{sm}] = next
}
