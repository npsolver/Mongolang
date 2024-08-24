package dfa

import (
	"fmt"
	"strings"
)

type Token struct {
	name string
	val  string
}

func (tk *Token) Print() {
	fmt.Println(tk.name, tk.val)
}

func NewToken(name, val string) *Token {
	return &Token{strings.ToUpper(name), val}
}
