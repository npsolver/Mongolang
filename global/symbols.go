package global

import (
	"fmt"
)

const (
	Arrow     = "->"
	Separator = "@"
)

type Symbol struct {
	name          string
	val           string
	isTerminating bool
}

func NewTerminatingSymbol(name, val string) *Symbol {
	return &Symbol{name, val, true}
}

func NewNonTerminatingSymbol(name, val string) *Symbol {
	return &Symbol{name, val, true}
}

func (tk *Symbol) GetName() string {
	return tk.name
}

func (tk *Symbol) GetValue() string {
	return tk.val
}

func (tk *Symbol) Print() {
	fmt.Println(tk.name, tk.val)
}
