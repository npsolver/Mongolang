package dfa

type Char string

func (s Char) append(t Char) Char {
	return s + t
}
