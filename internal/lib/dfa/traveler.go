package dfa

type Traveler[T any, S comparable] struct {
	dfa       *DFA[T, S]
	str       []*Symbol[S]
	n         int
	nextIndex int
}

func (d *DFA[T, S]) NewDFATraveler(s []*Symbol[S]) *Traveler[T, S] {
	return &Traveler[T, S]{
		dfa:       d,
		str:       s,
		n:         len(s),
		nextIndex: 0,
	}
}
