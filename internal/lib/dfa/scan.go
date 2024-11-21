package dfa

func (d *DFA[T, S]) AddSymbol(s S) {
	d.symbols = append(d.symbols, &Symbol[S]{s})
}

func (d *DFA[T, S]) AddState(name StateName, value T, isAccepting, isStartingState bool) {
	temp := &State[T, S]{
		name:        name,
		value:       value,
		bridges:     make(map[Symbol[S]]*State[T, S]),
		isAccepting: isAccepting,
	}
	d.states[name] = temp
	if isStartingState {
		d.startingState = temp
	}
}

func (d *DFA[T, S]) AddTransition(p StateName, b S, n StateName) {
	prev := d.getState(p)
	next := d.getState(n)
	prev.addBridge(b, next)
}
