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

func NewNonTerminatingSymbol(name string) *Symbol {
	return &Symbol{name, "", true}
}

func (sym *Symbol) GetName() string {
	return sym.name
}

func (sym *Symbol) GetValue() string {
	return sym.val
}

func (sym *Symbol) IsTerminating() bool {
	return sym.isTerminating
}

func (sym *Symbol) Print() {
	fmt.Println(sym.name, sym.val)
}
